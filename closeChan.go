package main

import "fmt"

func worker(i int, jobs <-chan int, done chan bool) {
	for {
		j, more := <-jobs
		if more {
			fmt.Println(i, "received job", j)
		} else {
			fmt.Println(i, "received all jobs")
			done <- true
			return
		}
	}
}
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go worker(0, jobs, done)

	go worker(1, jobs, done)

	for j := 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	<-done
	fmt.Println("done")
}
