package main

import "fmt"

type s struct {
    a int
    p *bool
    i *int
}
func main() {
    a := s{}
    a.a=1234
    b := true
    a.p = &b
    fmt.Println(a)
}

