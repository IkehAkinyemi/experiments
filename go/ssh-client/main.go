package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
    // Set up the SSH client configuration
    sshConfig := &ssh.ClientConfig{
        User: "username",
        Auth: []ssh.AuthMethod{
            // You can use a password, key, or agent-based authentication
            ssh.Password("password"),
            // ssh.PublicKeys(getPrivateKey()),
        },
        Timeout: 30 * time.Second,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    // Connect to the remote server
    conn, err := ssh.Dial("tcp", "remote-server.com:22", sshConfig)
    if err != nil {
        fmt.Printf("Failed to dial: %s", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Open a new SSH session
    session, err := conn.NewSession()
    if err != nil {
        fmt.Printf("Failed to create session: %s", err)
        os.Exit(1)
    }
    defer session.Close()

    // Run a command on the remote server
    out, err := session.Output("ls -la")
    if err != nil {
        fmt.Printf("Failed to run command: %s", err)
        os.Exit(1)
    }
    fmt.Println(string(out))
}

// Helper function to load a private key from a file
func getPrivateKey() ssh.Signer {
    key, err := ioutil.ReadFile("/path/to/private/key")
    if err != nil {
        fmt.Printf("Failed to read private key: %s", err)
        os.Exit(1)
    }
    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        fmt.Printf("Failed to parse private key: %s", err)
        os.Exit(1)
    }
    return signer
}
