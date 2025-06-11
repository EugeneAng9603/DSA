package main

import "context"

func main() {
	// passing context to a goroutine
	// to allow it to be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// perform work
			}
		}
	}(ctx)
}
