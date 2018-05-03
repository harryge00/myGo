package main

import (
	"fmt"
)

type testClient struct {
	id   int  // A unique id identifying this client.
	slow bool // True iff this client reads slowly.
}

func newTestClients(num int, slow bool) []*testClient {
	clients := make([]*testClient, num)
	for i := range clients {
		clients[i] = &testClient{id: i, slow: slow}
	}
	return clients
}

type uid string
func main() {
	ss := fmt.Sprintf("%s-%d", "asss", 19)
	fmt.Println(ss)
	manyClients := newTestClients(9, false)
	for i := range manyClients {
		cli := manyClients[i]
		fmt.Printf("id:%d + %b\n", cli.id, cli.slow)
	}
	var id uid
	id = "aaaa"
	var s string
	s = string(id)
	fmt.Println(s)
}
