# Agent Spinner

Unicode spinner animations for Go CLI applications. Beautiful, smooth, and highly customizable terminal spinners with 47+ built-in styles.

## Features

- 🎨 **47+ Built-in Spinners** - From classic braille to modern animations
- ⚡ **Lightweight** - Zero external dependencies
- 🎯 **Thread-Safe** - Concurrent access safe by design
- 🔧 **Customizable** - Create your own spinners easily
- 📦 **Simple API** - Start spinning in 3 lines of code
- 🖥️ **Terminal Aware** - Proper cursor handling and cleanup

## Installation

```bash
go get github.com/benoitpetit/agent-spinner
```

## Quick Start

```go
package main

import (
    "time"
    agentspinner "github.com/benoitpetit/agent-spinner"
)

func main() {
    spinner := agentspinner.Start("Loading...")
    time.Sleep(2 * time.Second)
    spinner.Stop("Done!")
}
```

## Usage Examples

### Basic Usage

```go
// Default braille spinner
spinner := agentspinner.Start("Processing...")
spinner.Stop("Complete!")
```

### With Different Styles

```go
// Sci-fi style
spinner := agentspinner.Start("Training AI...", agentspinner.Helix)

// Matrix/cyberpunk style
spinner = agentspinner.Start("Decrypting...", agentspinner.Matrix)

// Progress bar style
spinner = agentspinner.Start("Uploading...", agentspinner.Progress)
```

### Update Message During Operation

```go
spinner := agentspinner.Start("Step 1/3...")
spinner.Update("Step 2/3...")
spinner.Update("Step 3/3...")
spinner.Stop("Finished!")
```

### Error Handling

```go
spinner := agentspinner.Start("Validating...")
// On error:
spinner.Fail("Validation failed!")
```

### Run Helper (Simplified)

```go
// Execute function with automatic spinner
err := agentspinner.Run("Uploading...", func() error {
    // Your work here
    return nil
}, agentspinner.Radar)
```

### Run with Result

```go
result, err := agentspinner.RunWithResult("Computing...", func() (string, error) {
    return "42", nil
}, agentspinner.Star)
```

### Custom Spinner

```go
custom := agentspinner.Spinner{
    Frames:   []string{"◐", "◓", "◑", "◒"},
    Interval: 100, // milliseconds
}
spinner := agentspinner.StartCustom("Loading...", custom)
```

### Signal Handling

```go
// Graceful shutdown on Ctrl+C
spinner := agentspinner.Start("Working...").HandleSignals()
spinner.Stop()
```

### Custom Output Format

```go
// Fully control the output format
renderer := agentspinner.NewRawRenderer()
renderer.Output = os.Stdout
renderer.FormatFrame = "\r→ %s %s"     // Custom prefix
renderer.FormatFinal = "\r→ %s %s\n"   // Custom final format

spinner := agentspinner.StartCustom("Processing...", 
    agentspinner.NewDefaultRegistry().Get(agentspinner.Dots),
    agentspinner.WithRenderer(renderer),
)
spinner.Stop("Done!")
```

## Available Spinners

### Classic Braille Spinners

| Name | Preview | Description |
|------|---------|-------------|
| `Braille` | `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏` | Classic rotating braille (default) |
| `BrailleWave` | `⠁⠂⠄⡀ → ⠂⠄⡀⢀ → ⠄⡀⢀⠠` | Wave pattern across braille cells |
| `DNA` | `⠋⠉⠙⠚ → ⠉⠙⠚⠒ → ...` | Double helix DNA animation |

### Scan & Movement

| Name | Preview | Description |
|------|---------|-------------|
| `Scan` | Vertical scanner bar moving left to right |
| `ScanDual` | Dual vertical scanners converging |
| `ScanLine` | Horizontal scanning line |
| `Cascade` | Diagonal cascading effect |
| `FillSweep` | Bottom-to-top fill animation |
| `DiagSwipe` | Diagonal wipe transition |

### Wave Animations

| Name | Preview | Description |
|------|---------|-------------|
| `Wave` | Horizontal sine wave |
| `WaveVertical` | Vertical sine wave |
| `WaveRows` | Row-based wave motion |
| `Ripple` | Expanding circular ripple |

### Geometric Patterns

| Name | Preview | Description |
|------|---------|-------------|
| `Helix` | Double helix spiral |
| `Orbit` | Orbiting dot pattern |
| `Circle` | Drawing circle animation |
| `Diamond` | Expanding diamond shape |
| `Cross` | Rotating cross/plus |
| `Star` | Four-pointed star rotation |
| `Zigzag` | Zigzag pattern |

### Progress & Loading

| Name | Preview | Description |
|------|---------|-------------|
| `Progress` | Left-to-right progress bar |
| `Loading` | Progressive fill animation |
| `Bars` | Animated bar chart |
| `Columns` | Filling columns sequentially |
| `Expand` | Expanding rectangle |
| `Shrink` | Shrinking rectangle |

### Modern & Fun

| Name | Preview | Description |
|------|---------|-------------|
| `Matrix` | Matrix-style falling code |
| `Pulse` | Pulsing heartbeat effect |
| `Radar` | Radar/sonar sweep |
| `Typing` | Typing cursor animation |
| `Bounce` | Bouncing ball effect |
| `PingPong` | Ping pong game simulation |
| `Snake` | Snake game-like path |

### Decorative

| Name | Preview | Description |
|------|---------|-------------|
| `Sparkle` | Random sparkle pattern |
| `Rain` | Matrix-style rain drops |
| `Tiles` | Tiling pattern animation |
| `Checkerboard` | Alternating checker pattern |
| `RandomFill` | Random cell filling |
| `Border` | Border tracing animation |

### Arrows & Indicators

| Name | Preview | Description |
|------|---------|-------------|
| `Arrow` | Moving arrow pointer |
| `Dots` | Animated dot matrix |
| `Breathe` | Breathing effect |

### Complete List (47 spinners)

**Classic:** `Braille`, `BrailleWave`, `DNA`

**Scan:** `Scan`, `ScanLine`, `ScanDual`, `Cascade`, `FillSweep`, `DiagSwipe`

**Wave:** `Wave`, `WaveVertical`, `WaveRows`, `Ripple`

**Geometric:** `Helix`, `Orbit`, `Circle`, `Diamond`, `Cross`, `Star`, `Zigzag`, `Checkerboard`

**Progress:** `Progress`, `Loading`, `Bars`, `Columns`, `Expand`, `Shrink`, `RandomFill`, `Border`

**Modern:** `Matrix`, `Pulse`, `Radar`, `Typing`, `Bounce`, `PingPong`, `Snake`, `Rain`, `Sparkle`

**Decorative:** `Tiles`, `Dots`, `Breathe`

**Arrows:** `Arrow`

## API Reference

### Functions

| Function | Description |
|----------|-------------|
| `Start(message, name?)` | Create and start a spinner |
| `StartCustom(message, spinner, opts...)` | Start with custom spinner definition |
| `Run(message, fn, name?)` | Execute function with spinner |
| `RunWithResult(message, fn, name?)` | Execute function returning a value |
| `NewTerminalRenderer()` | Default renderer (stderr, formatted) |
| `NewRawRenderer()` | Configurable renderer (agnostic) |
| `NewSilentRenderer()` | No-op renderer (quiet mode) |

### Instance Methods

| Method | Description |
|--------|-------------|
| `Update(message)` | Change the spinner message |
| `Stop(message?)` | Stop with success (✓) |
| `Fail(message?)` | Stop with error (✗) |
| `HandleSignals()` | Enable graceful shutdown on interrupt |

### Types

```go
type Spinner struct {
    Frames   []string // Animation frames
    Interval int      // Milliseconds between frames
}
```

## Advanced Usage

### Custom Registry

```go
reg := agentspinner.NewMutableRegistry()
reg.Register("custom", agentspinner.Spinner{
    Frames:   []string{"▁", "▃", "▄", "▅", "▆", "▇", "█"},
    Interval: 100,
})
```

### With Options

```go
spinner := agentspinner.StartCustom("Loading...", customSpinner,
    agentspinner.WithRenderer(myRenderer),
    agentspinner.WithClock(testClock),
)
```

### Custom Renderer (RawRenderer)

Use `RawRenderer` for full control over output formatting:

```go
// Output to stdout with custom format
renderer := agentspinner.NewRawRenderer()
renderer.Output = os.Stdout
renderer.FormatFrame = "\r\033[K[%s] %s"   // [frame] message
renderer.FormatFinal = "\r\033[K[%s] %s\n" // [symbol] message + newline

spinner := agentspinner.StartCustom("Loading...", customSpinner,
    agentspinner.WithRenderer(renderer),
)
```

**RawRenderer options:**
- `Output` - Destination writer (default: `os.Stdout`)
- `FormatFrame` - Format string for frames (default: `"\r\033[K%s %s"`)
- `FormatFinal` - Format string for final output (default: `"\r\033[K%s %s\n"`)
- `EnableCursor` - Whether to hide/show cursor (default: `true`)

**Migration from TerminalRenderer:**

| Feature | TerminalRenderer | RawRenderer |
|---------|-----------------|-------------|
| Default output | stderr | stdout |
| Auto padding | 2 spaces | none |
| Auto newline | yes | configurable |
| Custom format | no | yes |

## Examples

Run the example:

```bash
cd examples/basic
go run main.go
```

## Requirements

- Go 1.21 or later
- Unicode-supporting terminal for best results

## License

MIT License - see [LICENSE](LICENSE) file

## Contributing

Contributions welcome! Feel free to add new spinner animations or improve existing ones.
