package agentspinner

import (
	"bytes"
	"testing"
	"time"

	"github.com/benoitpetit/agent-spinner/internal/animations"
	"github.com/benoitpetit/agent-spinner/internal/braille"
)

func TestRegistry(t *testing.T) {
	reg := NewDefaultRegistry()

	// Test that all spinners are defined
	for _, name := range reg.List() {
		s := reg.Get(name)
		if len(s.Frames) == 0 {
			t.Errorf("Spinner %s has no frames", name)
		}
		if s.Interval <= 0 {
			t.Errorf("Spinner %s has invalid interval: %d", name, s.Interval)
		}
	}
}

func TestRegistryGet(t *testing.T) {
	reg := NewDefaultRegistry()

	// Test getting existing spinner
	s := reg.Get(Braille)
	if len(s.Frames) == 0 {
		t.Error("Expected braille spinner to have frames")
	}

	// Test getting non-existent spinner (should return default)
	s = reg.Get("nonexistent")
	if len(s.Frames) == 0 {
		t.Error("Expected default spinner to have frames")
	}
}

func TestRegistryRegister(t *testing.T) {
	reg := NewMutableRegistry()
	custom := Spinner{
		Frames:   []string{"◐", "◓", "◑", "◒"},
		Interval: 100,
	}

	reg.Register("custom", custom)
	s := reg.Get("custom")
	if len(s.Frames) != 4 {
		t.Errorf("Expected 4 frames, got %d", len(s.Frames))
	}
}

func TestGridToBraille(t *testing.T) {
	// Test empty grid
	result := braille.GridToBraille([][]bool{})
	if result != "" {
		t.Errorf("Expected empty string for empty grid, got %q", result)
	}

	// Test nil grid
	result = braille.GridToBraille(nil)
	if result != "" {
		t.Errorf("Expected empty string for nil grid, got %q", result)
	}

	// Test single dot
	grid := braille.MakeGrid(4, 2)
	grid[0][0] = true
	result = braille.GridToBraille(grid)
	if result != "⠁" {
		t.Errorf("Expected '⠁' for dot at (0,0), got %q", result)
	}

	// Test multiple dots
	grid = braille.MakeGrid(4, 4)
	grid[0][0] = true
	grid[1][1] = true
	result = braille.GridToBraille(grid)
	if result == "" {
		t.Error("Expected non-empty string for diagonal pattern")
	}
}

func TestMakeGrid(t *testing.T) {
	// Test valid grid
	grid := braille.MakeGrid(4, 8)
	if len(grid) != 4 {
		t.Errorf("Expected 4 rows, got %d", len(grid))
	}
	if len(grid[0]) != 8 {
		t.Errorf("Expected 8 columns, got %d", len(grid[0]))
	}

	// Test invalid dimensions
	grid = braille.MakeGrid(0, 4)
	if grid != nil {
		t.Error("Expected nil grid for 0 rows")
	}

	grid = braille.MakeGrid(4, 0)
	if grid != nil {
		t.Error("Expected nil grid for 0 columns")
	}
}

func TestCloneGrid(t *testing.T) {
	original := braille.MakeGrid(4, 4)
	original[0][0] = true
	original[1][1] = true

	clone := braille.CloneGrid(original)
	if len(clone) != len(original) {
		t.Error("Clone should have same dimensions")
	}

	// Modify clone should not affect original
	clone[0][0] = false
	if original[0][0] != true {
		t.Error("Modifying clone affected original")
	}
}

// mockRenderer implements Renderer for testing
type mockRenderer struct {
	frames []string
	finals []string
}

func (m *mockRenderer) HideCursor() {}
func (m *mockRenderer) ShowCursor() {}
func (m *mockRenderer) RenderFrame(frame, message string) {
	m.frames = append(m.frames, frame)
}
func (m *mockRenderer) RenderFinal(symbol, message string) {
	m.finals = append(m.finals, symbol+":"+message)
}

// mockClock implements Clock for testing
type mockClock struct {
	ticker *mockTicker
}

func (m *mockClock) NewTicker(interval int) Ticker {
	return m.ticker
}

type mockTicker struct {
	c      chan time.Time
	closed bool
}

func (m *mockTicker) Chan() <-chan time.Time { return m.c }
func (m *mockTicker) Stop()                  { m.closed = true }

func TestInstance(t *testing.T) {
	mock := &mockRenderer{}
	ticker := &mockTicker{c: make(chan time.Time)}
	clock := &mockClock{ticker: ticker}

	spinner := Spinner{
		Frames:   []string{"◐", "◓", "◑", "◒"},
		Interval: 100,
	}

	inst := StartCustom("Test message", spinner, WithRenderer(mock), WithClock(clock))

	if inst.State != StateRunning {
		t.Error("Expected spinner to be running")
	}

	// Simulate a tick
	go func() {
		ticker.c <- time.Now()
	}()
	time.Sleep(50 * time.Millisecond)

	inst.Stop()

	if inst.State != StateStopped {
		t.Error("Expected spinner to be stopped")
	}

	if len(mock.finals) != 1 {
		t.Errorf("Expected 1 final render, got %d", len(mock.finals))
	}
}

func TestInstanceUpdate(t *testing.T) {
	mock := &mockRenderer{}
	ticker := &mockTicker{c: make(chan time.Time)}
	clock := &mockClock{ticker: ticker}

	spinner := Spinner{
		Frames:   []string{"◐"},
		Interval: 100,
	}

	inst := StartCustom("Initial", spinner, WithRenderer(mock), WithClock(clock))
	inst.Update("Updated")

	if inst.Message != "Updated" {
		t.Errorf("Expected message 'Updated', got %q", inst.Message)
	}

	inst.Stop()
}

func TestInstanceFail(t *testing.T) {
	mock := &mockRenderer{}
	ticker := &mockTicker{c: make(chan time.Time)}
	clock := &mockClock{ticker: ticker}

	spinner := Spinner{
		Frames:   []string{"◐"},
		Interval: 100,
	}

	inst := StartCustom("Will fail", spinner, WithRenderer(mock), WithClock(clock))
	inst.Fail("Failed!")

	if inst.State != StateFailed {
		t.Error("Expected spinner to be in failed state")
	}

	if len(mock.finals) != 1 || mock.finals[0] != "✗:Failed!" {
		t.Errorf("Expected failure render, got %v", mock.finals)
	}
}

func TestRun(t *testing.T) {
	err := Run("Test run", func() error {
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRunWithError(t *testing.T) {
	err := Run("Test run", func() error {
		time.Sleep(50 * time.Millisecond)
		return bytes.ErrTooLarge
	})
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRunWithResult(t *testing.T) {
	result, err := RunWithResult("Test result", func() (string, error) {
		time.Sleep(50 * time.Millisecond)
		return "success", nil
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "success" {
		t.Errorf("Expected 'success', got %q", result)
	}
}

func TestSpecificSpinners(t *testing.T) {
	reg := NewDefaultRegistry()
	newSpinners := []Name{
		Dots, Bounce, Ripple, Bars, Radar,
		Typing, Loading, PingPong, Star, Arrow,
		Wave, Progress, Circle, Cross, Zigzag,
		Diamond, Tiles, Expand, Shrink, ScanDual, WaveVertical,
		RandomFill, Border, Matrix,
	}
	for _, name := range newSpinners {
		s := reg.Get(name)
		if len(s.Frames) == 0 {
			t.Errorf("Spinner %s has no frames", name)
		}
		for i, frame := range s.Frames {
			if frame == "" {
				t.Errorf("Spinner %s frame %d is empty", name, i)
			}
		}
	}
}

func BenchmarkGridToBraille(b *testing.B) {
	grid := braille.MakeGrid(4, 8)
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			grid[i][j] = (i+j)%2 == 0
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = braille.GridToBraille(grid)
	}
}

func BenchmarkSpinnerLookup(b *testing.B) {
	reg := NewDefaultRegistry()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = reg.Get(Helix)
	}
}

func BenchmarkAnimationGeneration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = animations.Generate("matrix")
	}
}
