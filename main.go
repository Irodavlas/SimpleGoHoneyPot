package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	errBadPassword = errors.New("permission denied")
	port           = ":4123"
)

func main() {

	config := &ssh.ServerConfig{
		MaxAuthTries: 0,
		// PasswordCallback is used to authenticate the attempt, if succesfull returns the permissions
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			log.Printf("[INFO] attacker: %s, password: %s ", conn.RemoteAddr().String(), string(password))
			time.Sleep(500 * time.Millisecond)
			return nil, errBadPassword
		},
		AuthLogCallback: func(conn ssh.ConnMetadata, method string, err error) {
			log.Printf("[INFO] loggin attempt from: %s", conn.User())
		},
		ServerVersion: "SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.5",
	}

	// this is needed to make the server trustworthy for the client
	private_key, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer, _ := ssh.NewSignerFromSigner(private_key)
	config.AddHostKey(signer)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("[FATAL] failed to listen: ", err)
	}
	println("[INFO] Listening on port", port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("[FATAL] error accepting connection: ", err)
		}
		go handleConnection(conn, config)

	}
}
func handleConnection(conn net.Conn, config *ssh.ServerConfig) {
	defer conn.Close()
	log.Println("[INFO] remote address:", conn.RemoteAddr())
	// Creates a new ssh server conn if there is a signer for the keys
	ssh.NewServerConn(conn, config)
}
