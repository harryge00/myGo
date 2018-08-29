package main

import (
	"os"
	"log"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ns"
		"github.com/vishvananda/netlink"

)

func setDownAndRenameLink(linkName string) error {
	macvlanIface, err := netlink.LinkByName(linkName)
	if err != nil {
		return err
	}
	err = netlink.LinkSetDown(macvlanIface)
	if err != nil {
		return err
	}
	randName, err := ip.RandomVethName()
	if err != nil {
		return err
	}
	err = ip.RenameLink(linkName, randName)
	if err != nil {
		return err
	}
	return netlink.LinkDel(macvlanIface)
}

func main() {
	netdev, sandBoxKey := os.Args[1], os.Args[2]

	log.Printf("netdev: %v, sandBoxKey: %v", netdev, sandBoxKey)

	err := ns.WithNetNSPath(sandBoxKey, func(_ ns.NetNS) error {
		// randName, err := ip.RandomVethName()
		// if err != nil {
		// 	log.Printf("failed to random name %v", err)
		// 	randName = netdev
		// } else {
		// 	err = ip.RenameLink(netdev, randName)
		// 	if err != nil {
		// 		log.Printf("RenameLink %v: %v", randName, err)
		// 		// return err
		// 	}
		// }
		// if _, err := ip.DelLinkByNameAddr(netdev); err != nil {
		// 	log.Printf("Failed to DelLinkByNameAddr: %v", err)
		// 	if err != ip.ErrLinkNotFound {
		// 		return err
		// 	}
		// }
		// log.Printf("Successfully DelLinkByNameAddr %v", netdev)
		err := setDownAndRenameLink(netdev)
		log.Println(err)
		return nil
	})

	log.Println(err)
}