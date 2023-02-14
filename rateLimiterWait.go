package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func wait() {
	limiter := rate.NewLimiter(20, 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; ; i++ {
		fmt.Printf("%03d %s\n", i, time.Now().Format("2006-01-02 15:04:05.000"))
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
	}
	// 8 3 3 3 3
}

func main() {
	wait()
}
