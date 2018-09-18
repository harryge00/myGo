package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	links, err := netlink.LinkList()
	if err != nil {
		panic(err)
	}
	for i := range links {
		fmt.Println(links[i].Attrs())
	}
}