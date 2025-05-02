package pipeline

// Pipeline is a chain of processing elements arranged so that the output of each element is the input of the next.
// It is a common pattern in concurrent programming, especially in Go,
// where goroutines and channels are used to create a pipeline of processing stages.

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Pipeline() {
	fmt.Println("Starting pipeline process...")

	records, err := read("./assets/file3.csv")

	if err != nil {
		log.Fatalf("Error reading csv %v", err)
	}

	for val := range sanitize(titleize(records)) { // `titleize()` is taking the first character of each one of the values in each column of each line and capitalizing it
		fmt.Printf("%v\n", val)
	}

	fmt.Println("Pipeline completed")
}

func read(file string) (<-chan []string, error) {
	ch := make(chan []string)

	// Open the CSV file
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	// Read the CSV file in a goroutine
	go func() {
		// Create a new CSV reader
		csvReader := csv.NewReader(f)
		csvReader.FieldsPerRecord = 3 // Allow variable number of fields per record

		for {
			// Read a record from the CSV file
			record, err := csvReader.Read()

			if errors.Is(err, io.EOF) {
				close(ch) // Close the channel when done
				return
			}

			ch <- record // Send the record to the channel
		}
	}()

	return ch, nil
}

// sanitize takes a channel of strings and returns a channel of sanitized strings.
// It removes leading and trailing whitespace from each string in the slice.
func sanitize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for str := range strC {
			if len(str[0]) > 5 {
				continue // Skip if the slice has more than 5 elements in the first column
			}

			// Copy the slice to avoid modifying the original slice
			copy(str, str)

			ch <- str
		}

		close(ch)
	}()

	return ch
}

// titleize capitalizes the first letter of each string in the slice
// and returns a channel of the modified strings.
func titleize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		// Iterate over the strings in the channel
		for str := range strC {
			for i, s := range str {
				// Capitalize the first letter of each string in the slice
				str[i] = strings.Title(s)
			}
			ch <- str
		}
		close(ch)
	}()

	return ch
}
