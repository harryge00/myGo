package main

import "time"
import "log"

func Ticker() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		log.Printf("%v\n", now)
	}
}
func cal(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(1000 * time.Millisecond)
	}
}

// func timer(tch chan time.Time) {
// 	for {
// 		tch <- time.After(time.Duration(1500)*time.Millisecond)
// 	}
// }
func main() {
	ch := make(chan int)
	go cal(ch)
	ticker := time.NewTicker(1500 * time.Millisecond)
	// go timer(tchan)
	for {
		select {
		case num := <-ch:
			log.Printf("num %d", num)
		case <-ticker.C:
			log.Printf("time %v", time.Now())
		}
	}

}
