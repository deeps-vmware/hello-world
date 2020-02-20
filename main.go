package main

import (
	"fmt"
	"os"
	"net"
	"net/http"
	"io/ioutil"
)

var count = 0

func index(w http.ResponseWriter, req *http.Request) {
	count = count + 1
	fmt.Fprintf(w, "Hello World! from %s:%s (%d)\n", os.Getenv("NODE_ID"), os.Getenv("PORT"), count)
	getUpstream(w)
}

func getUpstream(w http.ResponseWriter) {
	upstream := os.Getenv("UPSTREAM")
	if upstream != "" {
		resp, err := http.Get(upstream)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, "Upstream: %s", string(body))
	}
}


func main() {
	http.HandleFunc("/", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		os.Setenv("PORT", port)
	}

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID, _ = os.Hostname()
		os.Setenv("NODE_ID", nodeID)
	}
	fmt.Printf("Hello World! from %s:%s\n", nodeID + " " + localAddr.IP.String(), os.Getenv("PORT"))
	if os.Getenv("UPSTREAM") != "" {
		fmt.Printf("Upstream: %s\n", os.Getenv("UPSTREAM"))
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
