package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(ip string, port int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return
	}
	conn.Close()
	results <- port
}

func main() {
	ip := "scanme.nmap.org"
	results := make(chan int)
	var wg sync.WaitGroup

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(ip, port, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		fmt.Printf("Порт %d открыт!\n", port)
	}
}
