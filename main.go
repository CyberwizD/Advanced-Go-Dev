package main

import (
	"github.com/CyberwizD/Advanced-Go-Dev/concurrency"
	faninpattern "github.com/CyberwizD/Advanced-Go-Dev/concurrency/fan-in_pattern"
	fanoutpattern "github.com/CyberwizD/Advanced-Go-Dev/concurrency/fan-out_pattern"
	"github.com/CyberwizD/Advanced-Go-Dev/concurrency/errgroup"
)

func main() {
	// Basic Concurrency
	concurrency.Basic_Concurrency()

	// Using the `Select` Statement in Concurrency
	concurrency.SelectCase_Concurrency()

	// Fan In Concurrency Pattern
	faninpattern.FanIn()

	// Fan Out Concurrency Pattern
	fanoutpattern.FanOut()

	// Using the `errgroup` Package for Concurrency
	errgroup.ErrGroup()
}
