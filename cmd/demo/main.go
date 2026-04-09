package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"time"

	agentspinner "github.com/benoitpetit/agent-spinner"
)

func main() {
	var (
		spinnerName = flag.String("spinner", "", "Name of the spinner to display")
		showAll     = flag.Bool("all", false, "Display all available spinners")
	)
	flag.Parse()

	if !*showAll && *spinnerName == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *showAll {
		demoAll()
		return
	}

	demoSingle(*spinnerName)
}

func demoSingle(name string) {
	reg := agentspinner.NewDefaultRegistry()
	spinner := reg.Get(agentspinner.Name(name))

	inst := agentspinner.StartCustom(name, spinner)
	time.Sleep(3 * time.Second)
	inst.Stop()
}

func demoAll() {
	reg := agentspinner.NewDefaultRegistry()
	names := reg.List()

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	for _, name := range names {
		spinner := reg.Get(name)
		inst := agentspinner.StartCustom(string(name), spinner)
		time.Sleep(1500 * time.Millisecond)
		inst.Stop()
	}

	fmt.Println("\nAll spinners demonstrated!")
}
