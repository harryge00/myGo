package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"net"
)

func test() {
	// get loopback interface
	link, err := netlink.LinkByName("enp0s31f6")
	if err != nil {
		fmt.Println(err)
		return
	}

	// bring the interface up
	if err := netlink.LinkSetUp(link); err != nil {
		fmt.Println(err)
		return
	}

	// add a gateway route
	dst := &net.IPNet{
		IP:   net.IPv4(192, 167, 0, 0),
		Mask: net.CIDRMask(24, 32),
	}

	// ip := net.IPv4(127, 1, 1, 1)
	route := netlink.Route{LinkIndex: link.Attrs().Index, Dst: dst, Gw: net.IPv4(10, 35, 59, 1)}
	if err := netlink.RouteAdd(&route); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	test()
	// ipv4 := net.ParseIP("173.3.15.4")
	// mask := net.CIDRMask(16, 32)
	// netv4 := &net.IPNet{IP: ipv4, Mask: mask}

	// fmt.Println(netv4)
	// maskedIP := ipv4.Mask(mask)
	// maskedIP[3]++
	// fmt.Println(maskedIP)

	// // ipaddr := &netlink.Addr{IPNet: netv4}
	// // err = netlink.AddrAdd(iface, ipaddr)
	// // if err != nil {
	// // 	glog.Infof("failed to add ipv4 address %v", err)
	// // 	return err
	// // }
	// iface, err := netlink.LinkByName("enp0s31f6")
	// fmt.Println(err)
	// err = netlink.LinkSetUp(iface)
	// fmt.Println(err)

	// // fmt.Println(err)
	// // 		err = netlink.AddrAdd(iface, ipaddr)

	// err = netlink.RouteAdd(&netlink.Route{
	// 	LinkIndex: iface.Attrs().Index,
	// 	Scope:     netlink.SCOPE_UNIVERSE,
	// 	Dst:       netv4,
	// 	// Gw: maskedIP
	// })
	// fmt.Println(err)
}
