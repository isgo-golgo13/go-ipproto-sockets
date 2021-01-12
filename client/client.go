package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	clog "github.com/Sirupsen/logrus"
	"golang.org/x/net/ipv4"
)

type ClientSvc interface {
	Connect() error
	// 	Send(bytes []byte) (int, error)
	// 	Sendn(bytes []byte, n int) (int, error)
	Close() error
	Loop()
}

type IPClient struct {
	ClientCxt        context.Context
	ClientPacketConn net.PacketConn
	ClientIPConn     *ipv4.RawConn

	Logger *clog.Logger
}

/** Variants in the Client IP Header as defined package */
type IPHeaderPkg struct {
	Data        []byte
	ID          int
	ServerProto int
	ServerIP    net.IP
}

func NewIPHeaderPkg(data []byte, id, serverProto int, server string) *IPHeaderPkg {
	var serverIP net.IP
	if server != "" {
		serverIP = net.ParseIP(server)
	}
	pkg := &IPHeaderPkg{
		Data:        data,
		ID:          id,
		ServerProto: serverProto,
		ServerIP:    serverIP,
	}
	return pkg
}

func init() {

}

func NewIPClient(proto int) (*IPClient, error) {
	client := &IPClient{
		Logger: &clog.Logger{Out: os.Stderr, Formatter: &clog.TextFormatter{ForceColors: true, FullTimestamp: true}, Hooks: make(clog.LevelHooks), Level: clog.InfoLevel | clog.ErrorLevel},
	}
	var err error
	client.ClientPacketConn, err = net.ListenPacket(fmt.Sprintf("ip4:%d", proto), "0.0.0.0")
	if err != nil {
		client.Logger.Errorf("%v", err)
		return nil, err
	}

	client.ClientIPConn, err = ipv4.NewRawConn(client.ClientPacketConn)
	if err != nil {
		client.Logger.Errorf("%v", err)
		return nil, err
	}

	//client.Logger.Infof("IP Client at: %s", config.CliConfig.Client+":"+fmt.Sprintf("%d", config.CliConfig.ClientSrcPort))
	return client, nil
}

func NewIPClientHeader(headerPkg *IPHeaderPkg) *ipv4.Header {

	return &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TotalLen: ipv4.HeaderLen + len(headerPkg.Data),
		ID:       headerPkg.ID,
		Protocol: headerPkg.ServerProto,
		Dst:      headerPkg.ServerIP.To4()
	}
}

func (c *IPClient) Close() {
	c.ClientIPConn.Close()
}

func (c *IPClient) Loop(wg *sync.WaitGroup, done <-chan interface{}) error {
	go func() error {
		defer wg.Done()

		for {
			select {
			case <-done:
				return nil
			default:

				// nw, err := c.Send(c.Data)
				// if err != nil {
				// 	c.Logger.Errorf("Error sending to server: %v", err)
				// 	return err
				// }
				// c.Logger.Infof("Client sent %d bytes Server %s: ", nw, fmt.Sprintf("%s:%d", "", 0))
			}
		}
	}()
	return nil
}
