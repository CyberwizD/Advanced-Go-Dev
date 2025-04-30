package faninpattern

// FanIn is a function that demonstrates the fan-in pattern in Go.
// It reads data from two CSV files concurrently and merges the results into a single channel.
// The function uses goroutines and channels to achieve concurrency and synchronization.

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func FanIn() {
	ch1, err := read("C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/fan-in_pattern/file1.csv")

	if err != nil {
		panic(fmt.Errorf("error reading csv file: %v", err))
	}

	ch2, err := read("C:/Users/WISDOM/Documents/Python Codes/GoLang/Advanced Go Dev/concurrency/fan-in_pattern/file2.csv")

	if err != nil {
		panic(fmt.Errorf("error reading csv file: %v", err))
	}

	// Create a channel to gracefully shutdown the channels
	exit := make(chan struct{})

	// Merge the two channels into one
	chM := mergedChannel(ch1, ch2)

	// goroutine to range over `ch1` and `ch2`
	go func() {
		for v := range chM {
			fmt.Println(v)
		}

		close(exit)
	}()

	<-exit

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
		defer close(ch) // Close the channel when done

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

func mergedChannel(channels ...<-chan []string) <-chan []string {
	var wg sync.WaitGroup
	out := make(chan []string)

	// Start a goroutine for each channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan []string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	// Close the output channel when all channels are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
