package main
import "fmt"

func FibonacciProducer(ch chan int, count int) {
    n2, n1 := 0, 1
    for count >= 0 {
        ch <- n2
        count--
        n2, n1 = n1, n2+n1
    }
    // close(ch)
}
func main() {
    ch := make(chan int)
    go FibonacciProducer(ch, 10)
    idx := 0
    for num := range ch {
        fmt.Printf("F(%d): \t%d\n", idx, num)
        idx++
    }
}