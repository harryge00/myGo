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

func isOneEditDistance(s string, t string) bool {
    if s == t {
        return false
    }
    if len(s) > len(t) {
        // always make s the shorter string
        return isOneEditDistance(t, s)
    }
    if len(t) - len(s) > 1 {
        return false
    }
    for i := 0; i < len(s); i++ {
        if s[i] != t[i] {
            if i == (len(s) - 1) && len(s) == len(t) {
                return true
            } else if s[i + 1:] == t[i+1:] {
                return true
            } else if s[i:] == t[i+1:] {
                return true
            } else {
                return false
            }
        }
    }
    return true
    
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
    if len(s) < 3 {
        return len(s)
    }
    startIndices := make([]int, len(s))
    sameCharStart, res, curLen := 0, 0, 1
    otherChar := s[0]
    for i := 1; i < len(s); i++ {
        if s[i] != otherChar && s[i] != s[startIndices[i - 1]] {
            otherChar = s[i]
            curLen = i - sameCharStart + 1
            startIndices[i] = sameCharStart
        } else {
            startIndices[i] = startIndices[i - 1]
            curLen++
        }
        if curLen > res {
            res = curLen
        }
        if s[i] != s[i - 1] {
            sameCharStart = i
        }
    }
    return res
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
    if len(s) < 3 {
        return len(s)
    }
    sameCharStart, res, curLen, left := 0, 0, 1, 0
    otherChar := s[0]
    for i := 1; i < len(s); i++ {
        if s[i] != otherChar && s[i] != s[left] {
            otherChar = s[i]
            curLen = i - sameCharStart + 1
            left = sameCharStart
        } else {
            curLen++
        }
        if curLen > res {
            res = curLen
        }
        if s[i] != s[i - 1] {
            sameCharStart = i
        }
    }
    return res
}

func main() {
    fmt.Println(lengthOfLongestSubstringTwoDistinct("abc"))
    fmt.Println(lengthOfLongestSubstringTwoDistinct("ab"))
    fmt.Println(lengthOfLongestSubstringTwoDistinct("abbbbbx"))
    fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
    fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))

    fmt.Println(isOneEditDistance("abc", "bc"))
    fmt.Println(isOneEditDistance("abc", "xbc"))
    fmt.Println(isOneEditDistance("abc", "xc"))
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
