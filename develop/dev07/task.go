package main

import (
	"fmt"
	"sync"
	"time"
)

func channelProcessing(channel <-chan interface{}, single chan<- interface{}, once *sync.Once) {
	defer once.Do(func() { close(single) })

	for {
		_, open := <-channel
		if !open {
			single <- nil
			break
		}
	}
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	once := &sync.Once{}

	for _, done := range channels {
		go channelProcessing(done, single, once)
	}

	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
