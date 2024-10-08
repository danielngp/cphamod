package main

import (
	"fmt"
//	"io/ioutil"
	"log"
	"golang.org/x/crypto/ssh"
	"time"
)

// SSHConfig stores SSH connection information
type SSHConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func main() {
	config := SSHConfig{
		Host:     "192.168.1.1",  // Replace with your router's IP
		Port:     "22",           // SSH port
		Username: "",        // Replace with your SSH username
		Password: "",     // Replace with your SSH password
	}

	// Create SSH client configuration
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Not secure, for testing purposes only
		Timeout:         5 * time.Second,
	}

	// Connect to the router
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// Open a session
	
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()
	// Run a command on the router
	output, err := session.CombinedOutput("cat /etc/openwrt_release") // Replace with the command you want to run
	if err != nil {
		log.Fatalf("Failed to run command: %s", err)
	}

	// Print the output of the command
	fmt.Println(string(output))
}
