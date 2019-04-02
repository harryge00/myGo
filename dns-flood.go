package main

import (
	"net"
	"fmt"
	"time"
	"os"
)

func doDNSQuery(domain string) error {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		return err
	}
	for _, ip := range ips {
		fmt.Printf("%v. IN A %s\n", domain, ip.String())
	}
	return nil
}

func main() {
	domain := "test.marathon.slave.mesos.zj.chinamobile.com"
	if len(os.Args) != 0 {
		domain = os.Args[1]
	}
	qps := 3
	fmt.Println(domain)
	ticker := time.NewTicker(1000 * time.Millisecond)
    for t := range ticker.C {
        fmt.Println("Tick at", t)
        for i := 0; i < qps; i++ {
        	go doDNSQuery(domain)
        }
    }
}