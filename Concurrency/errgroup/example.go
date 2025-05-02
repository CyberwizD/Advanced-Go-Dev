package errgroup

// The `errgroup` package is not part of go standard library,
// It is imported from `golang.org/x/sync/errgroup` and uses a way to
// synchronize, propagate errors, and context cancellation of goroutines.
// It works for a group of goroutines working on a common task.

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

func ErrGroup() {
	// The errgroup package is useful for managing a group of goroutines
	// and handling errors in a clean and efficient way.
	syncWaitGroup := waitGroups() // Using the `sync` standard library
	errWaitGroup := errGroup()    // Using the `errgroup` package

	<-syncWaitGroup
	<-errWaitGroup

	fmt.Println("All goroutines finished")
}

func read(file string) (<-chan []string, error) {
	ch := make(chan []string)

	// Open the CSV file
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	// Create a new CSV reader
	csvReader := csv.NewReader(f)

	// Read the CSV file in a goroutine
	go func() {
		// defer close(ch) // Close the channel when done

		for {
			// Read a record from the CSV file
			record, err := csvReader.Read()

			if errors.Is(err, io.EOF) {
				close(ch)

				return // End of file reached
			}

			ch <- record // Send the record to the channel
		}
	}()

	return ch, nil
}

func waitGroups() <-chan struct{} {
	ch := make(chan struct{}, 1)

	var wg sync.WaitGroup

	for _, file := range []string{"C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/file1.csv", "C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/file1.csv"} {
		file := file

		wg.Add(1)

		go func() {
			defer wg.Done()

			ch, err := read(file)

			if err != nil {
				fmt.Printf("Error reading %v", err)
			}

			for line := range ch {
				fmt.Println(line)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func errGroup() <-chan struct{} {
	ch := make(chan struct{}, 1)

	var g errgroup.Group

	for _, file := range []string{"C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/file1.csv", "C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/file2.csv"} {
		file := file

		g.Go(func() error {
			ch, err := read(file)

			if err != nil {
				return fmt.Errorf("error reading %v", err)
			}

			for line := range ch {
				fmt.Println(line)
			}

			return nil
		}) // `g.Go(func() error)` - When any of the goroutines returns an error (not nil), the `errgroup` will cancel all other goroutines.
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading file: %v", err)
		}

		close(ch) // Close the channel when done
	}()

	return ch
}
