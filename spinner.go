// Package agentspinner provides Unicode spinner animations for CLI applications.
package agentspinner

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Instance represents a running spinner.
// This is the primary type users interact with.
type Instance struct {
	// Spinner is the animation definition (read-only after start).
	Spinner Spinner

	// Message is the current display message (thread-safe).
	Message string

	// State is the current state of the spinner.
	State State

	renderer Renderer
	clock    Clock
	ticker   Ticker
	stopChan chan struct{}
	mu       sync.RWMutex
	frameIdx int
	running  bool
}

// Start creates and starts a new spinner with the given message.
// Uses the default registry and renderer.
func Start(message string, name ...Name) *Instance {
	n := Braille
	if len(name) > 0 {
		n = name[0]
	}

	reg := NewDefaultRegistry()
	s := reg.Get(n)

	inst := &Instance{
		Spinner:  s,
		Message:  message,
		State:    StateRunning,
		renderer: NewTerminalRenderer(),
		clock:    &realClock{},
		stopChan: make(chan struct{}),
	}

	inst.start()
	return inst
}

// StartCustom starts a spinner with a custom spinner definition.
func StartCustom(message string, spinner Spinner, opts ...Option) *Instance {
	cfg := &config{
		renderer: NewTerminalRenderer(),
		clock:    &realClock{},
	}
	for _, opt := range opts {
		opt(cfg)
	}

	inst := &Instance{
		Spinner:  spinner,
		Message:  message,
		State:    StateRunning,
		renderer: cfg.renderer,
		clock:    cfg.clock,
		stopChan: make(chan struct{}),
	}

	inst.start()
	return inst
}

// start begins the animation loop.
func (inst *Instance) start() {
	inst.mu.Lock()
	defer inst.mu.Unlock()

	if inst.running {
		return
	}

	inst.running = true
	inst.ticker = inst.clock.NewTicker(inst.Spinner.Interval)
	inst.renderer.HideCursor()

	go inst.render()
}

// render is the animation loop running in its own goroutine.
func (inst *Instance) render() {
	for {
		select {
		case <-inst.stopChan:
			return
		case <-inst.ticker.Chan():
			inst.mu.RLock()
			frame := inst.Spinner.Frames[inst.frameIdx%len(inst.Spinner.Frames)]
			msg := inst.Message
			inst.mu.RUnlock()

			inst.renderer.RenderFrame(frame, msg)
			inst.frameIdx++
		}
	}
}

// Update changes the spinner message.
func (inst *Instance) Update(message string) {
	inst.mu.Lock()
	defer inst.mu.Unlock()
	inst.Message = message
}

// Stop stops the spinner with a success message.
func (inst *Instance) Stop(message ...string) {
	inst.stop("✓", message...)
	inst.State = StateStopped
}

// Fail stops the spinner with an error message.
func (inst *Instance) Fail(message ...string) {
	inst.stop("✗", message...)
	inst.State = StateFailed
}

// stop contains the common stop logic (DRY principle).
func (inst *Instance) stop(symbol string, message ...string) {
	inst.mu.Lock()
	if !inst.running {
		inst.mu.Unlock()
		return
	}
	inst.running = false

	if inst.ticker != nil {
		inst.ticker.Stop()
	}
	close(inst.stopChan)
	inst.mu.Unlock()

	inst.renderer.ShowCursor()

	msg := inst.Message
	if len(message) > 0 {
		msg = message[0]
	}
	inst.renderer.RenderFinal(symbol, msg)
}

// IsRunning returns true if the spinner is currently active.
// Deprecated: use inst.State == StateRunning instead.
func (inst *Instance) IsRunning() bool {
	inst.mu.RLock()
	defer inst.mu.RUnlock()
	return inst.running
}

// HandleSignals configures graceful shutdown on interrupt signals.
func (inst *Instance) HandleSignals() *Instance {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		inst.Stop()
		signal.Stop(sigChan)
	}()

	return inst
}

// Run executes a function with a spinner.
func Run(message string, fn func() error, name ...Name) error {
	s := Start(message, name...).HandleSignals()
	defer s.Stop()

	if err := fn(); err != nil {
		s.Fail()
		return err
	}
	return nil
}

// RunWithResult executes a function that returns a result with a spinner.
func RunWithResult[T any](message string, fn func() (T, error), name ...Name) (T, error) {
	var zero T
	s := Start(message, name...).HandleSignals()

	result, err := fn()
	if err != nil {
		s.Fail()
		return zero, err
	}

	s.Stop()
	return result, nil
}

// realClock implements Clock using the standard library.
type realClock struct{}

func (c *realClock) NewTicker(interval int) Ticker {
	return &realTicker{ticker: time.NewTicker(time.Duration(interval) * time.Millisecond)}
}

// realTicker wraps time.Ticker to implement the Ticker interface.
type realTicker struct {
	ticker *time.Ticker
}

func (t *realTicker) Chan() <-chan time.Time {
	return t.ticker.C
}

func (t *realTicker) Stop() {
	t.ticker.Stop()
}
