package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1 := regexp.MustCompile(`([a-zA-Z0-9_.-/]{0,253}/)`)
	fmt.Println(re1.FindString("k8s_ns/abc"))
	x := regexp.MustCompile(`(x{0,4})`)
	fmt.Println(x.FindString("xxxx"))
	re := regexp.MustCompile(`([a-zA-Z0-9_.-/]{0,253}/)[a-zA-Z0-9]([a-zA-Z0-9_.-]{0,61}[a-zA-Z0-9])`)
	fmt.Println(re.FindString("k8s_ns/lab.el/kubernetes-admin.caicloud.io/partition"))
}
