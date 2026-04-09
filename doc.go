// Package agentspinner provides Unicode spinner animations for CLI applications.
//
// Basic usage:
//
//	spinner := agentspinner.Start("Loading...")
//	time.Sleep(2 * time.Second)
//	spinner.Stop("Done!")
//
// With different spinner styles:
//
//	spinner := agentspinner.Start("Processing...", agentspinner.Helix)
//
// Using the Run helper:
//
//	err := agentspinner.Run("Working...", func() error {
//	    // Do work
//	    return nil
//	})
//
// Custom spinner:
//
//	custom := agentspinner.Spinner{
//	    Frames:   []string{"◐", "◓", "◑", "◒"},
//	    Interval: 100,
//	}
//	spinner := agentspinner.StartCustom("Loading...", custom)
//
// # Architecture
//
// The package follows a layered architecture with clear separation of concerns:
//
//   - types.go: Core types and interfaces (Renderer, Clock, Registry)
//   - spinner.go: Runtime with dependency injection support
//   - registry.go: Thread-safe spinner registry
//   - renderer.go: Terminal and silent renderer implementations
//   - internal/braille: Grid operations (internal implementation detail)
//   - internal/animations: Animation generators (internal implementation detail)
//
// # Testing
//
// The package is designed for testability. Use WithRenderer() and WithClock()
// to inject mocks for testing:
//
//	mockRenderer := &mockRenderer{}
//	mockClock := &mockClock{}
//	spinner := agentspinner.StartCustom("Test", s,
//	    agentspinner.WithRenderer(mockRenderer),
//	    agentspinner.WithClock(mockClock),
//	)
package agentspinner
