// Package agentspinner provides Unicode spinner animations for CLI applications.
package agentspinner

import "time"

// Name identifies a spinner animation style.
type Name string

// Spinner defines an animation sequence with timing.
type Spinner struct {
	// Frames contains the animation frames as Unicode strings.
	Frames []string

	// Interval is the delay between frames in milliseconds.
	Interval int
}

// State represents the current state of a running spinner.
type State int

const (
	// StateRunning indicates the spinner is active.
	StateRunning State = iota
	// StateStopped indicates the spinner completed successfully.
	StateStopped
	// StateFailed indicates the spinner stopped with an error.
	StateFailed
)

// Renderer defines the interface for output operations.
// This abstraction allows testing without actual terminal output.
type Renderer interface {
	// HideCursor hides the terminal cursor.
	HideCursor()
	// ShowCursor shows the terminal cursor.
	ShowCursor()
	// RenderFrame outputs a single frame with the given message.
	RenderFrame(frame, message string)
	// RenderFinal outputs the final state with a symbol and message.
	RenderFinal(symbol, message string)
}

// Clock defines the interface for time-based operations.
// This abstraction enables deterministic testing of timing behavior.
type Clock interface {
	// NewTicker creates a new ticker with the specified duration in milliseconds.
	NewTicker(interval int) Ticker
}

// Ticker defines the interface for time-based events.
type Ticker interface {
	// Chan returns the channel that receives ticks.
	Chan() <-chan time.Time
	// Stop stops the ticker.
	Stop()
}

// Registry defines the interface for spinner lookup.
// Implementations can be predefined or custom.
type Registry interface {
	// Get returns a spinner by name, falling back to a default if not found.
	Get(name Name) Spinner
	// List returns all available spinner names.
	List() []Name
}

// Option configures a Spinner instance.
type Option func(*config)

type config struct {
	renderer Renderer
	clock    Clock
	registry Registry
}

// WithRenderer sets a custom renderer.
func WithRenderer(r Renderer) Option {
	return func(c *config) { c.renderer = r }
}

// WithClock sets a custom clock for testing.
func WithClock(clk Clock) Option {
	return func(c *config) { c.clock = clk }
}

// WithRegistry sets a custom spinner registry.
func WithRegistry(reg Registry) Option {
	return func(c *config) { c.registry = reg }
}
