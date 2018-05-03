package main

import (
	"fmt"
	// "io/ioutil"
	"strings"
	"path/filepath"

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
	ddd:=strings.Split("aaa", "-")
	fmt.Println(ddd, len(ddd))

	logFiles := strings.Split("/var/log/*.log /home/*.txt abcdefg", " ")
	var logDirs []string
	// To avoid mounting unnecessary volume, which is not necessary but
	// helps the pod look clean
	// '/var/log/a.log' and '/var/log/dir/b.log' may mount both
	// '/var/log' and 'var/log/dir'
	// we sort directories and mount only their common parent directory
	for _, file := range logFiles {
		if len(file) < 2 || !strings.HasPrefix(file, "/") {
			continue
		}
		logDirs = append(logDirs, filepath.Dir(file))
	}
	fmt.Println(logDirs)
	sss := strings.Join(logFiles, ",")
	fmt.Println(sss)
	var strArr []string
	fmt.Println(strings.Join(strArr, ","))
	strArr = make([]string, 0)
	fmt.Println(strings.Join(strArr, ","))
}
