package pubsubpattern

/*
Publisher Subscriber pattern is typically used in distributed systems.
An event/message is being sent by the publisher to the subscribers
through a channel or an event bus or message broker.

The Publisher-Subscriber (Pub/Sub) pattern is a messaging pattern that enables components
to communicate asynchronously without requiring direct knowledge of each other.

Key Concepts:
* Publishers: Components that create and send messages (events).
* Subscribers: Components that receive and process messages they have subscribed to.
* Message: The data being transmitted, such as events, notifications, or commands.
* Channel/Topic: The medium through which messages are sent. Subscribers can subscribe to specific channels/topics.
* Message Broker/Event Bus: An intermediary that manages the flow of messages from publishers to subscribers.
*/

import (
	"fmt"
	"sync"
)

/*
`PubSub` type uses generics, and contains three fields:
- A slice of channels indicating the subscribers
- A boolean variable indicating whether the channels is closed or open
- A mutex to prevent multiple goroutines to access or share data
*/

type pubSub[T any] struct {
	subscribers []chan T
	closed      bool
	mu          sync.RWMutex
}

func NewPubSub[T any]() *pubSub[T] {
	return &pubSub[T]{
		mu: sync.RWMutex{},
	}
}

/*
Methods required:
- One for publishing events
- One for allowing subscribers to subscribe
- One for closing the whole event
*/

// Publisher Method Implementaion
func (p *pubSub[T]) Publish(value T) {
	// `RLock` & `RUnlock` is being used because
	// there's no mutation or modification internally in the `pubSub` type,
	// Just reading the values available in the subscribers slice.
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return
	}

	// Ranging through all the subscribers and sending the value to the subscribers
	for _, ch := range p.subscribers {
		ch <- value
	}
}

// Subscriber Method Implementation
func (s *pubSub[T]) Subscribe() <-chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Checking if the channel is closed
	if s.closed {
		return nil
	}

	// Creating a new subscriber
	r := make(chan T)

	// Appending to the subscribers
	s.subscribers = append(s.subscribers, r)

	return r
}

// Closing Event Implementation
func (c *pubSub[T]) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return
	}

	for _, ch := range c.subscribers {
		close(ch)
	}

	c.closed = true
}

func PubSub() {
	fmt.Println("Starting Publisher...")

	ps := NewPubSub[string]()

	wg := sync.WaitGroup{}

	sub1 := ps.Subscribe()

	go func() {
		wg.Add(1)

		for {
			select {
			case val, ok := <-sub1:
				if !ok {
					fmt.Println("Sub 1, existing...")
					wg.Done()
					return
				}

				fmt.Println("Sub 1 - value", val)
			}
		}
	}()

	sub2 := ps.Subscribe()

	go func() {
		wg.Add(1)

		for {
			select {
			case val, ok := <-sub2:
				if !ok {
					fmt.Println("Sub 2, existing...")
					wg.Done()
					return
				}

				fmt.Println("Sub 2 - value", val)
			}
		}
	}()

	ps.Publish("1.")
	ps.Publish("2.")
	ps.Publish("3.")

	ps.Close()

	wg.Wait()

	fmt.Println("Completed")
}
