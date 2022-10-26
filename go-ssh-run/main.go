package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type connect struct {
	user       string
	addr       string
	privateKey string
	cmd        string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Insuffient arguments!")
		os.Exit(1)
	}
	host := strings.Split(os.Args[1], "@")
	c := connect{
		user:       host[0],
		addr:       host[1],
		privateKey: os.Getenv("HOME") + "/.ssh/id_rsa",
		cmd:        os.Args[2],
	}
	fmt.Printf("Connecting %v@%v using %v\n", c.user, c.addr, c.privateKey)
	output, err := c.remoteRun()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}

func (c connect) remoteRun() (string, error) {

	privateKeyBytes, err := os.ReadFile(c.privateKey)
	if err != nil {
		panic("Failed to load private key")
	}
	key, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}
	config := &ssh.ClientConfig{
		User:            c.user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort(c.addr, "22"), config)
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(c.cmd)
	return b.String(), err
}
