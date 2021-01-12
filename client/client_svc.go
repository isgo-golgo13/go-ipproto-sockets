package main

import (
	"net"

	log "github.com/Sirupsen/logrus"

	"network-poc/ipprotoloop_svc/errors"

	"golang.org/x/net/ipv4"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	proto := 1

	data := []byte("000001-Proto Data")

	/** Loop **/
	for i := 0; i < 250; i++ {

		// conn, err := net.ListenPacket(fmt.Sprintf("ip4:%d", proto), "0.0.0.0")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer conn.Close()

		client, err := NewIPClient(proto)
		errors.CheckError(err)

		defer client.Close()

		// ipConn, err := ipv4.NewRawConn(client.ClientIPConn)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		ipv4Hdr := &ipv4.Header{
			Version:  ipv4.Version,
			Len:      ipv4.HeaderLen,
			TotalLen: ipv4.HeaderLen + len(data),
			ID:       00001,
			Protocol: proto,
			Dst:      ip.To4(),
		}

		if err := client.ClientIPConn.WriteTo(ipv4Hdr, data, nil); err != nil {
			log.Println(err)
		}

	} /** End for loop */

}
