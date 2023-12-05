package main

import (
	"bytes"
	"io"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	go startTelnetServer()

	time.Sleep(100 * time.Millisecond)

	go func() {
		cmd := exec.Command("go", "run", "task.go", "--host=127.0.0.1", "--port=8080", "--timeout=20s")
		cmd.Stdin = bytes.NewBufferString("Hello, Telnet Server!\n")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(5000 * time.Millisecond)
}

func startTelnetServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to the Telnet Server!\n"))
	io.Copy(conn, conn)
}
