package utils

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func CheckPorts(host string, ports []string) {
	var wg = sync.WaitGroup{}

	for _, port := range ports {
		wg.Add(1)
		go func() {
			timeout := time.Second

			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
			if err != nil {
				fmt.Printf("The port %s is closed for the host %s\n", port, host)
			}

			if conn != nil {
				fmt.Printf("The port %s is open for the host %s\n", port, host)
				defer conn.Close()
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}
