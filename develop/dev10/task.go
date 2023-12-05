package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	host := flag.String("host", "", "Host")
	port := flag.String("port", "", "Port")
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout")
	flag.Parse()

	if *host == "" || *port == "" {
		fmt.Println("Usage: go-telnet --timeout=<timeout> <host> <port>")
		os.Exit(1)
	}

	address := *host + ":" + *port

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Connection error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		close(done)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		conn.Close()
	}()

	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close()
	}()

	<-done
}
