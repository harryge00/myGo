package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1 := regexp.MustCompile(`([a-zA-Z0-9_.-/]{0,253}/)?[a-zA-Z0-9]([a-zA-Z0-9_.-]{0,61}[a-zA-Z0-9])?`)
	fmt.Println(re1.MatchString("k8s_ns/lab.el/kubernetes-admin.caicloud.io/partition"))
	// fmt.Println(re.MatchString("k8s_ns/lab.el/kubernetes-admin.caicloud.io/partition", 1))

	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", 3))
	fmt.Println(re.FindAllString("none", -1))

	fmt.Println(regexp.MatchString("([a-zA-Z0-9-]+)-[0-9]", "haoyuan-40"))
	fmt.Println(regexp.MatchString("([a-zA-Z0-9-]+)-[0-9]", "kube-system"))
	fmt.Println(regexp.MatchString("([a-zA-Z0-9-]+)-[0-9]", "default"))
	fmt.Println(regexp.MatchString("([a-zA-Z0-9-]+)-[0-9]", "senyang-34"))
	fmt.Println(regexp.MatchString("[pP][0-9]+", "max"))
	fmt.Println(regexp.MatchString("[pP][0-9]+", "avg"))
	fmt.Println(regexp.MatchString("p[0-9]+", "p99"))
	fmt.Println(regexp.MatchString("p[0-9]+", "p33"))
	fmt.Println(regexp.MatchString("p[0-9]+", "p999"))
}
