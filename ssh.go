package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.Password("Hya8srTWLSQRgOZp"),
		},
	}
	client, err := ssh.Dial("tcp", "180.101.191.88:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("cat /etc/kubernetes/known-tokens.csv"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	// client2, err := ssh.Dial("tcp", "180.101.191.88:22", config)
	// if err != nil {
	//     panic("Failed to dial: " + err.Error())
	// }
	// session2, err := client2.NewSession()
	// if err != nil {
	//     panic("Failed to create session: " + err.Error())
	// }

	// // Once a Session is created, you can execute a single command on
	// // the remote side using the Run method.
	// var b2 bytes.Buffer
	// session2.Stdout = &b2
	// if err := session2.Run("echo $HOME"); err != nil {
	//     panic("Failed to run: " + err.Error())
	// }
	// fmt.Println(b2.String())
	// session2.Close()
}
