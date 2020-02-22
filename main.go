package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// Counter ...
type Counter int32

func (c *Counter) increment() int32 {
	var next int32
	for {
		next = int32(*c) + 1
		if atomic.CompareAndSwapInt32((*int32)(c), int32(*c), next) {
			return next
		}
	}
}

var counter Counter

func index(w http.ResponseWriter, req *http.Request) {
	count := counter.increment()
	fmt.Fprintf(w, "Hello World! from %s:%s (%d)\n", os.Getenv("NODE_ID"), os.Getenv("PORT"), count)
	getUpstream(w)
}

func getUpstream(w http.ResponseWriter) {
	upstream := os.Getenv("UPSTREAM")
	if upstream != "" {
		timeout, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
		client := http.Client{
			Timeout: time.Duration(time.Duration(timeout) * time.Second),
		}

		resp, err := client.Get(upstream)
		if err != nil {
			fmt.Fprintf(w, "Upstream: %s\n", err)
			return
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

	if os.Getenv("NODE_ID") == "" {
		hostname, _ := os.Hostname()
		os.Setenv("NODE_ID", hostname)
	}

	if os.Getenv("TIMEOUT") == "" {
		os.Setenv("TIMEOUT", "5")
	}

	fmt.Printf("Hello World! from %s:%s\n", os.Getenv("NODE_ID")+" "+localAddr.IP.String(), os.Getenv("PORT"))
	if os.Getenv("UPSTREAM") != "" {
		fmt.Printf("Upstream: %s Timeout: %ss\n", os.Getenv("UPSTREAM"), os.Getenv("TIMEOUT"))
	}
	log.WithFields(log.Fields{
		"nodeID":   os.Getenv("NODE_ID"),
		"ip":       localAddr.IP.String(),
		"port":     os.Getenv("PORT"),
		"upstream": os.Getenv("UPSTREAM"),
		"timeout":  os.Getenv("TIMEOUT"),
	}).Info("Hello World!")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
