package main

import (
	"fmt"
	"time"
)

type GeneralHandler interface {
	Stop()
	loop() error
}

type Handler struct {
	reqCh chan string
	stopCh chan struct{}
}

func (h *Handler) Stop() {
    close(h.stopCh)
 
    // 可以使用WaitGroup等待所有协程退出
}

func handle(s string) {
	fmt.Printf("Handle %s\n", s)
}
// 收到停止后，不再处理请求
func (h *Handler) loop() error {
    for {
        select {
        case req := <-h.reqCh:
            go handle(req)
		case <-h.stopCh:
			fmt.Println("Received stop signal")
            return nil
        }
	}
	fmt.Println("stop loop")
	return nil
}

func tryWriteChan(stopCh chan struct{}) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in tryWriteChan", r)
        }
    }()
	tryWriteChanChild(stopCh)
}

func tryWriteChanChild(stopCh chan struct{}) {
	// defer func() {
    //     if r := recover(); r != nil {
    //         fmt.Println("Recovered in tryWriteChanChild", r)
    //     }
    // }()
	var s struct{}
	stopCh <- s
}

func main() {

	reqCh := make(chan string)
	stopCh := make(chan struct{})
	h1 := &Handler{reqCh, stopCh}
	fmt.Println("init")

	go h1.loop()
	fmt.Println("looping")
	// time.Sleep(1 * time.Second)
	
	h1.Stop()
	go tryWriteChan(h1.stopCh)
	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}