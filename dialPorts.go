package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, url string, port int, ch chan int) {
	defer wg.Done()
	var req = url + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", req, time.Second*5)
	if err != nil {
		return
	}
	defer conn.Close()
	ch <- port
}

func main() {
	startTime := time.Now()
	var url = os.Args[1]
	var ports = os.Args[2]
	reader := strings.NewReader(ports)

	var firstPort, lastPort int
	var dash rune
	fmt.Fscanf(reader, "%d%c%d", &firstPort, &dash, &lastPort)

	var wg sync.WaitGroup
	ch := make(chan int)
	for i := firstPort; i <= lastPort; i++ {
		wg.Add(1)
		go worker(&wg, url, i, ch)
	}

	go func() {
		for res := range ch {
			fmt.Printf("%d ", res)
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Printf("time cost = %s\n", time.Since(startTime))
}

// usage: go run dialPorts.go google.com 1-1024
