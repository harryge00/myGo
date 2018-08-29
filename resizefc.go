package main

import (
	"os"
	"fmt"
	"os/exec"
)

const resizeSh = "/usr/local/bin/sync2fs.sh"

func main() {
	volumeID := os.Args[1]
	fmt.Println(volumeID)
	stdoutStderr, err  := exec.Command(resizeSh, "dellsc", volumeID).CombinedOutput()

	if err != nil {
		fmt.Println(string(stdoutStderr), err)
	}
}