package main

import "fmt"

var Version = "No Version Provided"
var buildstamp = ".."
var githash = "xxx"

func main() {
    fmt.Printf("App Version: %s\n", Version)
    fmt.Printf("Git Commit Hash: %s\n", buildstamp)
    fmt.Printf("UTC Build Time: %s\n", githash)
}