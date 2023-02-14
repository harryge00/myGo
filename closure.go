package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func sum(a, b int) int {
	return a + b
}

func print(a, b int, f func(int, int) int) {

}

func main() {
	fmt.Println("Used cpus: ", runtime.NumCPU())
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	http.ListenAndServe(port, nil)
}
