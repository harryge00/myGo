package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	"strings"

)

func wordPattern(pattern string, str string) bool {
    arr := strings.Split(str, " ")
    m := make(map[string]string)
    m2 := make(map[string]bool)
    for i, word := range arr {
    	fmt.Println(string(pattern[i]))
        if val, ok := m[word]; !ok {
        	if m2[string(pattern[i])] {
        		return false
        	}
        	m2[string(pattern[i])] = true
            m[word] = string(pattern[i])
        } else {
            if val != string(pattern[i]) {
                return false
            }
        }
    fmt.Println(m, m2)
    }
    fmt.Println(m, m2)
    return true
}

func main() {
    abc := "aaaabb"
    fmt.Println(abc[0:3])
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
    arr1 := []string{"aaa","bbb", "ccc"}
    str1 := strings.Join(arr1, ";")
    fmt.Println(str1)
    fmt.Println(strings.Split(str1, ";"))
    arr2 := []string{}
    arr2[0] = "aaaa"
    fmt.Println(arr2 == nil)
    str2 := strings.Join(arr2, ";")
    fmt.Println(str2)
    fmt.Println(strings.Split(str2, ";"))
}
