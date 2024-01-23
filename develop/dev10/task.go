package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	host, port, timeout := parseCommandLineArgs()

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	go readServerData(conn)
	go sendUserData(conn)

	handleInterruptSignal(conn)
}

func parseCommandLineArgs() (string, int, time.Duration) {
	host := flag.String("host", "", "Target host (IP or hostname)")
	port := flag.Int("port", 0, "Target port")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if *host == "" || *port == 0 {
		fmt.Println("Usage: go-telnet <host> <port> --timeout=<timeout>")
		os.Exit(1)
	}

	return *host, *port, *timeout
}

func readServerData(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println("Server connection closed.")
	os.Exit(0)
}

func sendUserData(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Fprintf(conn, "%s\n", input)
	}
}

func handleInterruptSignal(conn net.Conn) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	fmt.Println("\nCtrl+C pressed. Closing connection.")
	conn.Close()
}
