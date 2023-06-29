package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	ctx, _ := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(3)

	stopCh := make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				log.Println("Close stopCh 1")
				return
			// msg from other goroutine finish
			case <-ctx.Done():
				// end
				log.Println("Close 1")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				log.Println("Close stopCh 2")
				return
			// msg from other goroutine finish
			case <-ctx.Done():
				// end
				log.Println("Close 2")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		// your operation
		// call cancel when this goroutine ends
		log.Println("Sleep 3")
		time.Sleep(3 * time.Second)
		log.Println("Going to cancel")
		//cancel()
		close(stopCh)
	}()

	wg.Wait()

}
