package main

import (

	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	//"strconv"
	//"bufio"
	//"math/big"
	//"bytes"
	
	"golang.org/x/net/ipv4"
	"lab01/protocol"
)

func main() {
	masterReader()
	// Second step : delay_request (point to point)
}


func masterReader() {

	// First step : multicast
	conn, err := net.ListenPacket("udp", protocol.MULTICAST_ADDR) // listen on port
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	p := ipv4.NewPacketConn(conn) // convert to ipv4 packetConn
	addr, err := net.ResolveUDPAddr("udp", protocol.MULTICAST_ADDR)
	if err != nil {
		log.Print(err)
	}
	var interf *net.Interface
	if runtime.GOOS == "darwin" {
		interf, _ = net.InterfaceByName("e2n0")
	}

	if err = p.JoinGroup(interf, addr); err != nil { // listen on ip multicast
		log.Fatal(err)
	}
	
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(buf) 
		if err != nil {
			log.Fatal(err)
		}	
		
		if strings.Compare(protocol.SYNC, string(buf[:len(protocol.SYNC)])) == 0 {
			fmt.Printf("Recieved SYNC with id : %s\n", buf[len(protocol.SYNC):n])
		}
		
		//fmt.Println(string(buf[:n]))
	}
}

func delayRequest() {
}
	