# Agent Spinner

Unicode spinner animations for Go CLI applications.

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

## Usage

### Basic

```go
spinner := agentspinner.Start("Processing...")
spinner.Stop("Complete!")
```

### With Style

```go
spinner := agentspinner.Start("Working...", agentspinner.Helix)
```

### Update Message

```go
spinner := agentspinner.Start("Step 1...")
spinner.Update("Step 2...")
spinner.Stop("Finished!")
```

### Run Helper

```go
err := agentspinner.Run("Uploading...", func() error {
    // do work
    return nil
})
```

### With Result

```go
result, err := agentspinner.RunWithResult("Computing...", func() (string, error) {
    return "42", nil
})
```

### Custom Spinner

```go
custom := agentspinner.Spinner{
    Frames:   []string{"◐", "◓", "◑", "◒"},
    Interval: 100,
}
spinner := agentspinner.StartCustom("Loading...", custom)
```

## Available Spinners

- `Braille` - Classic braille spinner (default)
- `BrailleWave`, `DNA`, `Scan`, `Rain`, `Pulse`
- `Matrix`, `Loading`, `Typing`, `Progress`
- `Helix`, `Radar`, `Ripple`, `Wave`
- And 30+ more...

## License

MIT
