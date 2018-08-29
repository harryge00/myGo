package main
import "fmt"
import "time"

func FibonacciProducer(ch chan int, count int) {
    time.Sleep(2 * time.Second)
    n2, n1 := 0, 1
    for count >= 0 {
        ch <- n2
        count--
        n2, n1 = n1, n2+n1
    }
    close(ch)
}

func caller(ch chan int, count int) {
    go FibonacciProducer(ch, count)
    fmt.Println("Now caller quits")
}

func main() {
    ch := make(chan int)
    go caller(ch, 10)
    fmt.Println("Now caller quits")
    idx := 0
    for num := range ch {
        fmt.Printf("F(%d): \t%d\n", idx, num)
        idx++
    }
}