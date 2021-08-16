package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	host := "106.244.179.242"
	port := 8554
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("%d CLOSED (%s)\n", port, err)
	}
	conn.Close()
	log.Printf("%d OPEN\n", port)
}
