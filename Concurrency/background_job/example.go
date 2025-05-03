package backgroundjob

// Background Job is a process in charge of doing some work "behind the scenes"
// Initialized by another "parent" process, which simply means that a goroutine launching `N` goroutines

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func BackgroundJob() {
	fmt.Println("Process ID", os.Getpid())

	listenForWork()

	<-waitToExit()

	fmt.Println("Exiting background job")
}

func listenForWork() {
	const WorkerN int = 5

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM)

	workersC := make(chan struct{}, WorkerN)

	// Listen for messages to process
	go func() {
		for {
			<-sc
			workersC <- struct{}{} // Send to processing channel
		}
	}()

	go func() {
		var workers int

		for range workersC { // Wait for the messaages to process
			workerID := (workers % WorkerN) * 1
			workers++

			fmt.Printf("%d<-\n", workerID)

			go func() {
				doWork(workerID)
			}()
		}
	}()
}

func waitToExit() <-chan struct{} {
	runC := make(chan struct{}, 1)

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, os.Interrupt)

	go func() {
		defer close(runC)

		<-sc
	}()

	return runC
}

func doWork(id int) {
	fmt.Printf("<-%d starting\n", id)

	time.Sleep(3 * time.Second)

	fmt.Printf("<-%d completed\n", id)
}
