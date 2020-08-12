package main

import "fmt"
import "bytes"

func countAndSay(n int) string {
    str := "1"
    var buffer bytes.Buffer
    for i := 1; i < n; i++ {
        cnt := 0
        lastChar := ' '
        for _, char := range str {
            if char == lastChar {
                cnt++
                continue
            }
            if lastChar != ' ' {
                buffer.WriteString(fmt.Sprintf("%d%c", cnt, lastChar))
            }
            lastChar = char
            cnt = 1
        }
        buffer.WriteString(fmt.Sprintf("%d%c", cnt, lastChar))
        str = buffer.String()
        buffer.Reset()
    }    
    return str
}

func main() {
    for pos, char := range "abcd" {
        fmt.Println(pos, char, 'a' == char)
    }
    fmt.Println(countAndSay(4))
    fmt.Println(countAndSay(5))
}