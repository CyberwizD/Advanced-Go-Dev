package main

import (
	"github.com/CyberwizD/Advanced-Go-Dev/concurrency"
	faninfanoutpattern "github.com/CyberwizD/Advanced-Go-Dev/concurrency/fan-in_fan-out-pattern"
)

func main() {
	// Basic Concurrency
	concurrency.Basic_Concurrency()

	// Using the `Select` Statement in Concurrency
	concurrency.SelectCase_Concurrency()

	// Fan In and Fan Out Concurrency Pattern
	faninfanoutpattern.FanIn()
}
