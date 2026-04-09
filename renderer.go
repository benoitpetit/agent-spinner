package agentspinner

import (
	"fmt"
	"io"
	"os"
)

// TerminalRenderer outputs to a terminal using ANSI escape codes.
// For a fully customizable renderer, use NewRawRenderer.
type TerminalRenderer struct {
	// Output is the destination writer (os.Stderr by default).
	Output io.Writer
}

// NewTerminalRenderer creates a renderer that writes to os.Stderr.
// Note: This renderer adds 2-space padding and automatic newlines.
// Use NewRawRenderer for full control over formatting.
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
// Formats as: "\r\033[K  <frame> <message>"
func (r *TerminalRenderer) RenderFrame(frame, message string) {
	fmt.Fprintf(r.Output, "\r\033[K  %s %s", frame, message)
}

// RenderFinal outputs the final state with a symbol and message.
// Formats as: "\r\033[K  <symbol> <message>\n"
func (r *TerminalRenderer) RenderFinal(symbol, message string) {
	fmt.Fprintf(r.Output, "\r\033[K  %s %s\n", symbol, message)
}

// RawRenderer is a fully customizable renderer with no imposed formatting.
// It gives you complete control over output formatting.
type RawRenderer struct {
	// Output is the destination writer (os.Stdout by default).
	Output io.Writer
	// FormatFrame is called for each frame. Use %s for frame and message.
	// Default: "\r\033[K%s %s"
	FormatFrame string
	// FormatFinal is called for the final output. Use %s for symbol and message.
	// Default: "\r\033[K%s %s\n"
	FormatFinal string
	// EnableCursor controls whether to send cursor hide/show ANSI codes.
	EnableCursor bool
}

// NewRawRenderer creates a fully customizable renderer.
// Defaults: stdout output, ANSI cursor control enabled.
func NewRawRenderer() *RawRenderer {
	return &RawRenderer{
		Output:       os.Stdout,
		FormatFrame:  "\r\033[K%s %s",
		FormatFinal:  "\r\033[K%s %s\n",
		EnableCursor: true,
	}
}

// NewRawRendererWithOutput creates a raw renderer with a custom output writer.
func NewRawRendererWithOutput(w io.Writer) *RawRenderer {
	return &RawRenderer{
		Output:       w,
		FormatFrame:  "\r\033[K%s %s",
		FormatFinal:  "\r\033[K%s %s\n",
		EnableCursor: true,
	}
}

// HideCursor hides the terminal cursor (if EnableCursor is true).
func (r *RawRenderer) HideCursor() {
	if r.EnableCursor {
		fmt.Fprint(r.Output, "\033[?25l")
	}
}

// ShowCursor shows the terminal cursor (if EnableCursor is true).
func (r *RawRenderer) ShowCursor() {
	if r.EnableCursor {
		fmt.Fprint(r.Output, "\033[?25h")
	}
}

// RenderFrame outputs a frame using the configured FormatFrame.
// FormatFrame receives: frame, message
func (r *RawRenderer) RenderFrame(frame, message string) {
	fmt.Fprintf(r.Output, r.FormatFrame, frame, message)
}

// RenderFinal outputs the final state using the configured FormatFinal.
// FormatFinal receives: symbol, message
func (r *RawRenderer) RenderFinal(symbol, message string) {
	fmt.Fprintf(r.Output, r.FormatFinal, symbol, message)
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
