package main

import (
	"fmt"
	"strings"
	"regexp"
	xj "github.com/basgys/goxml2json"
)

func main() {
	// buf, err := ioutil.ReadFile("/home/tattoo/Codes/DisSys/mit2015/6.824/src/main/part_kjv12.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// value := string(buf)
	// result := strings.Fields(value)
	// for _, w := range result {
	// 	fmt.Print(w)
	// }
	str := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cliOutput>
  <opRet>0</opRet>
  <opErrno>0</opErrno>
  <opErrstr/>
  <volStatus>
    <volumes>
      <volume>
        <volName>vol_98eb9387f874b92b729711cd02ab3e39</volName>
        <nodeCount>3</nodeCount>
        <node>
          <hostname>192.168.16.125</hostname>
          <path>/var/lib/heketi/mounts/vg_c1f5dba29a78e2eba0966c396340460c/brick_d3ede4de6a26a66573e6d3a6e73e18ca/brick</path>
          <peerid>10399bc3-0fcb-45c7-8ae5-f1200f0a96b7</peerid>
          <status>1</status>
          <port>49631</port>
          <ports>
            <tcp>49631</tcp>
            <rdma>N/A</rdma>
          </ports>
          <pid>10172</pid>
          <sizeTotal>2134900736</sizeTotal>
          <sizeFree>2100740096</sizeFree>
          <device>/dev/mapper/vg_c1f5dba29a78e2eba0966c396340460c-brick_d3ede4de6a26a66573e6d3a6e73e18ca</device>
          <blockSize>4096</blockSize>
          <mntOptions>rw,seclabel,noatime,nouuid,attr2,inode64,logbsize=256k,sunit=512,swidth=512,noquota</mntOptions>
          <fsName>xfs</fsName>
        </node>
        <node>
          <hostname>192.168.16.143</hostname>
          <path>/var/lib/heketi/mounts/vg_845a67802b8a9d8205fd819bb7a40fd5/brick_ba21f7939df0d63a78c221ade781eee5/brick</path>
          <peerid>c5d2d79d-ea0e-40b3-8e44-9c0d580454c4</peerid>
          <status>1</status>
          <port>49659</port>
          <ports>
            <tcp>49659</tcp>
            <rdma>N/A</rdma>
          </ports>
          <pid>10162</pid>
          <sizeTotal>2134900736</sizeTotal>
          <sizeFree>2100740096</sizeFree>
          <device>/dev/mapper/vg_845a67802b8a9d8205fd819bb7a40fd5-brick_ba21f7939df0d63a78c221ade781eee5</device>
          <blockSize>4096</blockSize>
          <mntOptions>rw,seclabel,noatime,nouuid,attr2,inode64,logbsize=256k,sunit=512,swidth=512,noquota</mntOptions>
          <fsName>xfs</fsName>
        </node>
        <node>
          <hostname>192.168.16.175</hostname>
          <path>/var/lib/heketi/mounts/vg_0b66032b6cdf29e53a6862406c79cb71/brick_87cac89af6e2466306250c1465aa7d61/brick</path>
          <peerid>c34aee4a-191e-4e2c-b6cb-5ee22df692bd</peerid>
          <status>1</status>
          <port>49622</port>
          <ports>
            <tcp>49622</tcp>
            <rdma>N/A</rdma>
          </ports>
          <pid>2425</pid>
          <sizeTotal>2134900736</sizeTotal>
          <sizeFree>2100740096</sizeFree>
          <device>/dev/mapper/vg_0b66032b6cdf29e53a6862406c79cb71-brick_87cac89af6e2466306250c1465aa7d61</device>
          <blockSize>4096</blockSize>
          <mntOptions>rw,seclabel,noatime,nouuid,attr2,inode64,logbsize=256k,sunit=512,swidth=512,noquota</mntOptions>
          <fsName>xfs</fsName>
        </node>
      </volume>
    </volumes>
  </volStatus>
</cliOutput>`
	brickStatuses := strings.Split(str, "-")
	for i, brickStatus := range brickStatuses {
		if brickStatus == "" {
			continue
		}
		fmt.Println(i)
		re, _ := regexp.Compile(`([A-Z]\w+[\w\s])\s+:\s(\w+)`)
		strs := re.FindAllString(brickStatus, -1)
		fmt.Println(len(strs), strs)
		// for i := 0; i < len(strs) - 1; i++ {
		// 	if i % 2 == 1 {
		// 		continue
		// 	}
		// 	fmt.Println(strs[i], "::", strs[i+1])
		// }
		// fields := strings.Split(brickStatus, ": ")
		// }
	}
}
