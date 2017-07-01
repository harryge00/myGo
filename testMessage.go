package main

import "fmt"
import "lsp"
import "encoding/json"

func main() {
	ack := lsp.NewAck(21, 0)
	ackBytes, _ := json.Marshal(ack)
	fmt.Printf("original ack: %s\n", ack)
	var msg2 *lsp.Message
	json.Unmarshal(ackBytes, msg2)
	fmt.Printf("Marshalled ack: %s\n", msg2)

	payload := []byte("abcdefg")
	data := lsp.NewData(255, 2, payload)
	dataBytes, _ := json.Marshal(data)
	fmt.Printf("data: %s\n", dataBytes)
	var msg lsp.Message
	json.Unmarshal(dataBytes, &msg)
	if msg.Type == lsp.MsgData {
		fmt.Println("Marshalled data: %s\n", &msg)
	}
	fmt.Println(msg.Type, lsp.MsgConnect, lsp.MsgAck)
	if msg.Type == 1 {
		fmt.Println("Marshalled data: %s\n", &msg)
	}
	msgmap := make(map[int]*lsp.Message)
	msgmap[0] = ack
	msgmap[1] = data
	msgmap[2] = lsp.NewConnect()
	fmt.Println("len map ", len(msgmap))
	msgmap[1] = nil
	fmt.Println("len map ", len(msgmap))
}
