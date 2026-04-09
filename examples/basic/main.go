package main

import (
	"fmt"
	"time"

	"github.com/benoitpetit/agent-spinner"
)

func main() {
	fmt.Println("=== Agent Spinner Examples ===")
	fmt.Println()

	// Example 1: Basic usage
	fmt.Println("1. Basic spinner (default 'braille'):")
	spinner := agentspinner.Start("Loading data...")
	time.Sleep(1500 * time.Millisecond)
	spinner.Stop("Data loaded!")

	// Example 2: Different spinner styles
	fmt.Println("\n2. Helix spinner (sci-fi vibe):")
	spinner = agentspinner.Start("Training AI model...", agentspinner.Helix)
	time.Sleep(1500 * time.Millisecond)
	spinner.Stop("Model ready!")

	// Example 3: Matrix-style
	fmt.Println("\n3. Matrix spinner (cyberpunk):")
	spinner = agentspinner.Start("Decrypting...", agentspinner.Matrix)
	time.Sleep(1500 * time.Millisecond)
	spinner.Stop("Decryption complete!")

	// Example 4: Update message during operation
	fmt.Println("\n4. Updating message with 'columns':")
	spinner = agentspinner.Start("Step 1/4: Connecting...", agentspinner.Columns)
	time.Sleep(600 * time.Millisecond)
	spinner.Update("Step 2/4: Authenticating...")
	time.Sleep(600 * time.Millisecond)
	spinner.Update("Step 3/4: Fetching data...")
	time.Sleep(600 * time.Millisecond)
	spinner.Update("Step 4/4: Processing...")
	time.Sleep(600 * time.Millisecond)
	spinner.Stop("All steps completed!")

	// Example 5: Progress-style spinner
	fmt.Println("\n5. Progress bar style:")
	spinner = agentspinner.Start("Uploading files...", agentspinner.Progress)
	time.Sleep(2000 * time.Millisecond)
	spinner.Stop("Upload complete!")

	// Example 6: Pulse for health checks
	fmt.Println("\n6. Pulse (for health checks):")
	spinner = agentspinner.Start("Checking system health...", agentspinner.Pulse)
	time.Sleep(1500 * time.Millisecond)
	spinner.Stop("System healthy!")

	// Example 7: Using Run helper
	fmt.Println("\n7. Using Run helper with 'radar':")
	err := agentspinner.Run("Scanning for threats...", func() error {
		time.Sleep(1500 * time.Millisecond)
		return nil
	}, agentspinner.Radar)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Example 8: Custom spinner
	fmt.Println("\n8. Custom spinner:")
	custom := agentspinner.Spinner{
		Frames:   []string{"◐", "◓", "◑", "◒"},
		Interval: 100,
	}
	spinner = agentspinner.StartCustom("Custom animation...", custom)
	time.Sleep(2000 * time.Millisecond)
	spinner.Stop()

	// Example 9: Error handling
	fmt.Println("\n9. Error handling with 'cross':")
	spinner = agentspinner.Start("Validating configuration...", agentspinner.Cross)
	time.Sleep(1000 * time.Millisecond)
	spinner.Fail("Validation failed!")

	// Example 10: RunWithResult
	fmt.Println("\n10. RunWithResult with 'star':")
	result, err := agentspinner.RunWithResult("Computing result...", func() (string, error) {
		time.Sleep(1500 * time.Millisecond)
		return "42", nil
	}, agentspinner.Star)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Println("\n=== All examples completed! ===")
}
