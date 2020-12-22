//demo1
package main
import "fmt"

func test_string(s string){
    fmt.Printf("inner: %v, %v\n",s, &s)
    s = "b"
    fmt.Printf("inner: %v, %v\n",s, &s)
}

func main() {
    s := "a"
    fmt.Printf("outer: %v, %v\n",s, &s)
    test_string(s)
    fmt.Printf("outer: %v, %v\n",s, &s)
}
