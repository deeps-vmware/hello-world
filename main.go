package main

import (
	"fmt"
	"os"
	"net"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World! from %s:%s\n", os.Getenv("NODE_ID"), os.Getenv("PORT"))
}

func main() {
	http.HandleFunc("/", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		os.Setenv("PORT", port)
	}

	nodeIP := os.Getenv("NODE_ID")
	if nodeIP == "" {
		conn, _ := net.Dial("udp", "8.8.8.8:80")
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		nodeIP = localAddr.IP.String()
		os.Setenv("NODE_ID", nodeIP)
	}
	fmt.Printf("Hello World! from %s:%s\n", os.Getenv("NODE_ID"), os.Getenv("PORT"))
	http.ListenAndServe(":"+port, nil)
}
