package main

import (
    //"fmt"
    "encoding/json"

    "log"
    "io/ioutil"
    "net/http"
    "math/rand"
    "strconv"
)


type IpResp struct {
    IP  string  `json:"ip,omitempty"`
    Mask  int  `json:"mask,omitempty"`
    Occupied  int  `json:"occupied,omitempty"`
}

func getIPhandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Ip Allocating:")
    body, _ := ioutil.ReadAll(r.Body)
    log.Println(string(body))
    ip := "192.168."
    i1 := rand.Intn(256)
    i2 := rand.Intn(256)
    ip += strconv.Itoa(i1) + "."
    ip += strconv.Itoa(i2)
    mask := 24
    resp := IpResp{
        IP: ip,
        Mask: mask,
    }
    log.Println(resp)
    bytes, _ := json.Marshal(resp)
    w.Header().Set("Content-Type", "application/json")
    w.Write(bytes)
}

func releaseIPhandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Ip Releasing:")
    body, _ := ioutil.ReadAll(r.Body)
    log.Println(string(body))
    w.WriteHeader(200)
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    log.Println("Serving")
    http.HandleFunc("/api/net/ip/occupy", getIPhandler)
    http.HandleFunc("/api/net/ip/release", releaseIPhandler)
    http.ListenAndServe(":8080", nil)
}