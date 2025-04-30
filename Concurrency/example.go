package concurrency

import (
	"fmt"
	"time"
)

type Users struct {
	Name string
	Age  int
}

func Basic_Concurrency() {
	// Make a channel with type `Users` and size `10`
	ch := make(chan Users, 10)

	// Slice of users with name and age
	users := []Users{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Dave", Age: 28},
		{Name: "Eve", Age: 22},
	}

	// Run an anonymous function in a goroutine
	go func() {
		fmt.Println(time.Now())

		for index, user := range users {
			fmt.Println("Sending user", index)
			fmt.Println(user.Name, "is", user.Age, "years old")

			// Send each user to the channel
			ch <- user

			fmt.Println("Sent.✅")
		}
	}()

	// Recieving from the channel
	v := <-ch
	i := <-ch

	fmt.Println("Validated Users :", v.Name, "and", i.Name)

	fmt.Println("Extitng concurrent program...")
}

func SelectCase_Concurrency() {
	// Make a channel with type `string` and size `10`
	ch := make(chan string, 5)

	// Slice of strings
	posts := []string{
		"Why Go is the best programming language",
		"How to use Go for backend development",
		"Concurrency in Go: A deep dive",
	}

	// Run an anonymous function in a goroutine
	go func() {
		for _, post := range posts {
			fmt.Println("\nSending post -", post)
			ch <- post
			time.Sleep(1 * time.Second) // Simulate some delay in sending
			fmt.Println("Post Sent.✅")
		}
	}()

	for i := 0; i < len(posts); i++ {
		select {
		case post1 := <-ch:
			fmt.Println("Received post:", post1)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout! No posts received.")
			return
		}
	}

	fmt.Println("Extitng concurrent program...")
}
