package ratelimit

/*
Rate limiting in Go controls the rate at which requests are processed,
preventing abuse and ensuring fair resource allocation
*/

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

/*
The golang.org/x/time/rate package provides a built-in Limiter based on the token bucket algorithm.
A token bucket starts with a fixed number of tokens, and each request consumes one token.
Tokens are replenished at a steady rate. If the bucket is empty, requests are either delayed or rejected.
*/

func RateLimiter() {
	// Create a rate limiter allowing 1 request per second, with a burst capacity of 5.
	limiter := rate.NewLimiter(rate.Limit(1), 5)

	for i := 0; i < 10; i++ {
		// Wait for a token to be available.
		ctx := context.Background()
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("Error waiting for limiter:", err)
			return
		}

		// Process the request.
		fmt.Println("Processing request", i+1, "at", time.Now().Format(time.RFC3339))
	}
}
