package agentspinner

import (
	"fmt"
	"io"
	"os"
)

// TerminalRenderer outputs to a terminal using ANSI escape codes.
type TerminalRenderer struct {
	// Output is the destination writer (os.Stderr by default).
	Output io.Writer
}

// NewTerminalRenderer creates a renderer that writes to os.Stderr.
func NewTerminalRenderer() *TerminalRenderer {
	return &TerminalRenderer{Output: os.Stderr}
}

// NewTerminalRendererWithOutput creates a renderer with a custom output.
func NewTerminalRendererWithOutput(w io.Writer) *TerminalRenderer {
	return &TerminalRenderer{Output: w}
}

// HideCursor hides the terminal cursor.
func (r *TerminalRenderer) HideCursor() {
	fmt.Fprint(r.Output, "\033[?25l")
}

// ShowCursor shows the terminal cursor.
func (r *TerminalRenderer) ShowCursor() {
	fmt.Fprint(r.Output, "\033[?25h")
}

// RenderFrame outputs a single frame with the given message.
func (r *TerminalRenderer) RenderFrame(frame, message string) {
	fmt.Fprintf(r.Output, "\r\033[K  %s %s", frame, message)
}

// RenderFinal outputs the final state with a symbol and message.
func (r *TerminalRenderer) RenderFinal(symbol, message string) {
	fmt.Fprintf(r.Output, "\r\033[K  %s %s\n", symbol, message)
}

// SilentRenderer discards all output (useful for testing or quiet mode).
type SilentRenderer struct{}

// NewSilentRenderer creates a no-op renderer.
func NewSilentRenderer() *SilentRenderer {
	return &SilentRenderer{}
}

// HideCursor is a no-op.
func (r *SilentRenderer) HideCursor() {}

// ShowCursor is a no-op.
func (r *SilentRenderer) ShowCursor() {}

// RenderFrame is a no-op.
func (r *SilentRenderer) RenderFrame(_, _ string) {}

// RenderFinal is a no-op.
func (r *SilentRenderer) RenderFinal(_, _ string) {}
