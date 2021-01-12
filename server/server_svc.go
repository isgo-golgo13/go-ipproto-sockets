package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	ipAddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	proto := "icmp"
	for {
		conn, lerr := net.ListenIP(fmt.Sprintf("ip4:%s", proto), ipAddr)
		if lerr != nil {
			log.Fatal(lerr)
		}

		data := make([]byte, 1024)
		n, clientIP, err := conn.ReadFromIP(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("IP Server received: %X from client: %s\n", data[:n], clientIP.String())

	} /** End for loop */
}
