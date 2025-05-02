package fanoutpattern

// FanOut is a function that demonstrates the fan-out pattern in Go.
// It reads data from a CSV file and writes the results to multiple channels concurrently.
// The function uses goroutines and channels to achieve concurrency and synchronization.

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func FanOut() {
	ch1, err := read("./assets/file1.csv")

	if err != nil {
		panic(fmt.Errorf("error reading csv file: %v", err))
	}

	fmt.Println("Splitting channels...")
	// Split the channel into multi channels
	chan1 := writeToChannel("1", ch1)
	chan2 := writeToChannel("2", ch1)

	for {
		if chan1 == nil && chan2 == nil {
			break
		}

		select {
		case _, ok := <-chan1:
			if !ok {
				chan1 = nil
			}
		case _, ok := <-chan2:
			if !ok {
				chan2 = nil
			}
		}
	}

	fmt.Println("All done!")
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
		defer close(ch)

		for {
			// Read a record from the CSV file
			record, err := csvReader.Read()

			if err == io.EOF {
				break // End of file reached
			} else if err != nil {
				fmt.Println("Error reading CSV:", err)
				return
			}

			ch <- record // Send the record to the channel
		}
	}()

	return ch, nil
}

func writeToChannel(worker string, cha <-chan []string) chan struct{} {
	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup
	chA := make(chan struct{})

	go func(chR <-chan []string) {
		defer close(chA) // Close the channel when done

		for v := range chR {
			wg.Add(1) // Increment the wait group counter
			go func(val []string) {
				defer wg.Done() // Decrement the counter when done
				fmt.Printf("Worker %s processing: %v\n", worker, val)
			}(v)
		}

		wg.Wait() // Wait for all goroutines to finish
	}(cha)

	return chA
}
