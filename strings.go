package main

import (
	"fmt"
	// "io/ioutil"
	"strconv"
    "net"
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
    s := "4: sbsbssss222    inet 192.168.1.3/16 scope global sbsbssss222       valid_lft forever preferred_lft forever"
    lines := strings.Split(s, "\n")
    if len(lines) < 1 {
        fmt.Printf("Peiqi Macvlan Unexpected command output %s", lines)
        return
    }
    fmt.Println(lines)
    fields := strings.Fields(lines[0])
    if len(fields) < 4 {
        fmt.Printf("Peiqi Macvlan Unexpected address output %s ", lines[0])
    }
    fmt.Println(fields[3])
    ip, ipNet, err := net.ParseCIDR(fields[3])
    fmt.Println(ip, ipNet.Mask, err)
    mask, err := strconv.Atoi(string(ipNet.Mask))
    if err != nil {
        fmt.Errorf("failed to convert mask %v", ipNet.Mask)
    }
    fmt.Println(mask)
}
