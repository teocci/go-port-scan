package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "106.244.179.242"
	for port := 5300; port <= 5350; port++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", port, err)
			continue
		}
		conn.Close()
		log.Printf("%d OPEN\n", port)
	}
}
