package main

import (
	"log"
	"io"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.Header)
	log.Println(r.RemoteAddr)
	// RemoteAddr
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	log.Println("Hello, 你好bbb!")
	io.WriteString(w, "Hello, 你好bbb!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":7493", nil)
}
