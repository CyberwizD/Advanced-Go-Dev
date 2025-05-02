package pipeline

// Pipeline is a chain of processing elements arranged so that the output of each element is the input of the next.
// It is a common pattern in concurrent programming, especially in Go,
// where goroutines and channels are used to create a pipeline of processing stages.

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func Pipeline() {
	records, err := read("./assets/file1.csv")

	if err != nil {
		log.Fatalf("Error reading csv %v", err)
	}

	for val := range sanitize(titleize(records)) {
		fmt.Printf("%v\n", val)
	}
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

func sanitize()
