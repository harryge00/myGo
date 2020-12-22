package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Header)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	text := os.Getenv("RETURN_TEXT")
	if text == "" {
		text = "Hello world!"
	}
	io.WriteString(w, text)
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
