package main

import (
	"fmt"
	"sync"
	"time"
)

func joinChannels(channels ...<-chan interface{}) <-chan interface{} {
	singleCh := make(chan interface{})

	go func() {
		var isClosed bool
		var mu sync.Mutex

		for _, ch := range channels {

			go func(ch <-chan interface{}) {

				for {
					_, isChanOpened := <-ch
					if !isChanOpened {
						break
					}
				}

				mu.Lock()
				if !isClosed {
					close(singleCh)
				}
				isClosed = true
				mu.Unlock()
			}(ch)
		}

	}()

	return singleCh
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
	singleChan := joinChannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	<-singleChan

	fmt.Printf("done after %v\n", time.Since(start))

}
