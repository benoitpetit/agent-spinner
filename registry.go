package agentspinner

import (
	"sync"

	"github.com/benoitpetit/agent-spinner/internal/animations"
)

// DefaultRegistry provides thread-safe access to predefined spinners.
type DefaultRegistry struct {
	spinners map[Name]Spinner
	default_ Name
	mu       sync.RWMutex
}

// NewDefaultRegistry creates a registry with the built-in spinners.
func NewDefaultRegistry() *DefaultRegistry {
	r := &DefaultRegistry{
		spinners: make(map[Name]Spinner),
		default_: Braille,
	}
	r.registerBuiltins()
	return r
}

// Get returns a spinner by name, falling back to default if not found.
func (r *DefaultRegistry) Get(name Name) Spinner {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if s, ok := r.spinners[name]; ok {
		return s
	}
	return r.spinners[r.default_]
}

// List returns all available spinner names.
func (r *DefaultRegistry) List() []Name {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]Name, 0, len(r.spinners))
	for name := range r.spinners {
		names = append(names, name)
	}
	return names
}

// Register adds a custom spinner to the registry.
func (r *DefaultRegistry) Register(name Name, s Spinner) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.spinners[name] = s
}

// SetDefault changes the default spinner returned when a name is not found.
func (r *DefaultRegistry) SetDefault(name Name) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.default_ = name
}

// registerBuiltins initializes the registry with all built-in spinners.
func (r *DefaultRegistry) registerBuiltins() {
	// Classic spinners - inline definitions
	r.spinners[Braille] = Spinner{
		Frames:   []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		Interval: 80,
	}
	r.spinners[BrailleWave] = Spinner{
		Frames: []string{
			"⠁⠂⠄⡀", "⠂⠄⡀⢀", "⠄⡀⢀⠠", "⡀⢀⠠⠐",
			"⢀⠠⠐⠈", "⠠⠐⠈⠁", "⠐⠈⠁⠂", "⠈⠁⠂⠄",
		},
		Interval: 100,
	}
	r.spinners[DNA] = Spinner{
		Frames: []string{
			"⠋⠉⠙⠚", "⠉⠙⠚⠒", "⠙⠚⠒⠂", "⠚⠒⠂⠂",
			"⠒⠂⠂⠒", "⠂⠂⠒⠲", "⠂⠒⠲⠴", "⠒⠲⠴⠤",
			"⠲⠴⠤⠄", "⠴⠤⠄⠋", "⠤⠄⠋⠉", "⠄⠋⠉⠙",
		},
		Interval: 80,
	}

	// Generated spinners - loaded from internal package
	r.loadGeneratedSpinners()
}

// loadGeneratedSpinners loads animation frames from the internal package.
func (r *DefaultRegistry) loadGeneratedSpinners() {
	// Map internal names to our Name constants
	nameMap := map[string]Name{
		"scan":         Scan,
		"rain":         Rain,
		"scanline":     ScanLine,
		"pulse":        Pulse,
		"snake":        Snake,
		"sparkle":      Sparkle,
		"cascade":      Cascade,
		"columns":      Columns,
		"orbit":        Orbit,
		"breathe":      Breathe,
		"waverows":     WaveRows,
		"checkerboard": Checkerboard,
		"helix":        Helix,
		"fillsweep":    FillSweep,
		"diagswipe":    DiagSwipe,
		"dots":         Dots,
		"bounce":       Bounce,
		"ripple":       Ripple,
		"bars":         Bars,
		"radar":        Radar,
		"typing":       Typing,
		"loading":      Loading,
		"pingpong":     PingPong,
		"star":         Star,
		"arrow":        Arrow,
		"wave":         Wave,
		"progress":     Progress,
		"circle":       Circle,
		"cross":        Cross,
		"zigzag":       Zigzag,
		"diamond":      Diamond,
		"tiles":        Tiles,
		"expand":       Expand,
		"shrink":       Shrink,
		"scandual":     ScanDual,
		"wavevertical": WaveVertical,
		"randomfill":   RandomFill,
		"border":       Border,
		"matrix":       Matrix,
	}

	// Interval configuration for each spinner type
	intervals := map[Name]int{
		Scan: 70, Rain: 100, ScanLine: 120, Pulse: 180,
		Snake: 80, Sparkle: 150, Cascade: 60, Columns: 60,
		Orbit: 100, Breathe: 100, WaveRows: 90, Checkerboard: 250,
		Helix: 80, FillSweep: 100, DiagSwipe: 60, Dots: 120,
		Bounce: 100, Ripple: 100, Bars: 90, Radar: 100,
		Typing: 80, Loading: 50, PingPong: 80, Star: 90,
		Arrow: 120, Wave: 100, Progress: 80, Circle: 100,
		Cross: 80, Zigzag: 120, Diamond: 120, Tiles: 200,
		Expand: 80, Shrink: 80, ScanDual: 70, WaveVertical: 100,
		RandomFill: 60, Border: 100, Matrix: 80,
	}

	for internalName, publicName := range nameMap {
		frames := animations.Generate(internalName)
		if len(frames) > 0 {
			interval := 100 // default
			if iv, ok := intervals[publicName]; ok {
				interval = iv
			}
			r.spinners[publicName] = Spinner{
				Frames:   frames,
				Interval: interval,
			}
		}
	}
}

// MutableRegistry is a convenience wrapper for dynamic spinner registration.
type MutableRegistry struct {
	*DefaultRegistry
}

// NewMutableRegistry creates a registry that allows runtime modifications.
func NewMutableRegistry() *MutableRegistry {
	return &MutableRegistry{NewDefaultRegistry()}
}
