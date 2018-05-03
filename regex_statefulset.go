package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var statefulPodRegex = regexp.MustCompile("(.*)-([0-9]+)$")
	subMatches := statefulPodRegex.FindStringSubmatch("web123-3")
	fmt.Println(subMatches)	
	if len(subMatches) < 3 {
		return
	}
	var ordinal int
	var parent string
	parent = subMatches[1]
	if i, err := strconv.ParseInt(subMatches[2], 10, 32); err == nil {
		ordinal = int(i)
	}
	fmt.Println(parent, ordinal)


}
