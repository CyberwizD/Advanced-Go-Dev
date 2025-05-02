package main

import (
	basic "github.com/CyberwizD/Advanced-Go-Dev/concurrency/basic"
	errgroup "github.com/CyberwizD/Advanced-Go-Dev/concurrency/errgroup"
	faninpattern "github.com/CyberwizD/Advanced-Go-Dev/concurrency/fan-in_pattern"
	fanoutpattern "github.com/CyberwizD/Advanced-Go-Dev/concurrency/fan-out_pattern"
	pipeline "github.com/CyberwizD/Advanced-Go-Dev/concurrency/pipeline_pattern"
)

func main() {
	// Basic Concurrency
	basic.Basic_Concurrency()

	// Using the `Select` Statement in Concurrency
	basic.SelectCase_Concurrency()

	// Fan In Concurrency Pattern
	faninpattern.FanIn()

	// Fan Out Concurrency Pattern
	fanoutpattern.FanOut()

	// Using the `errgroup` Package for Concurrency
	errgroup.ErrGroup()

	// Pipeline Concurrency Pattern
	pipeline.Pipeline()
}
