package main

import "io/ioutil"
import "net/http"
import (
	"log"
	"sync/atomic"
	"sync"
)

const (
	currentUrl = "http://10.35.48.196/api/repository/info?repo_name=library/keepalive"
)

func sendRequest(client *http.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", currentUrl, nil)
	req.SetBasicAuth("admin", "Harbor123457")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func main() {
	tr := &http.Transport{
		MaxIdleConns:       20,
		MaxIdleConnsPerHost:  20,
	}
	client := http.Client{Transport: tr}
	var count int32
	wg := sync.WaitGroup{}
	wg.Add(20000)
	for i := 0; i < 20000; i++ {
		go func(i int) {
			_, err := sendRequest(&client)
			if err != nil {
				log.Println(i, err)
			} else {
				log.Printf("Done request %v", i)
				atomic.AddInt32(&count, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("Successful: %v", count)
}