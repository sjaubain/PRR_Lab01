package main

import (

	"log"
	"net"
	"runtime"
	"strings"
	"strconv"
	"math/rand"
	//"bufio"
	//"math/big"
	//"bytes"
	
	"golang.org/x/net/ipv4"
	"lab01/protocol"
	"lab01/clock"
)

func main() {
	masterReader()
	// Second step : delay_request (point to point)
}

var slaveClock = clock.New(rand.Intn(clock.MAX_OFFSET))

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
	
	var currentId string
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFrom(buf) 
		if err != nil {
			log.Fatal(err)
		}	
		
		// Look for SYNC
		if strings.Compare(protocol.SYNC, string(buf[:len(protocol.SYNC)])) == 0 {
			currentId = string(buf[len(protocol.SYNC):n])
			log.Printf("Received SYNC with id : %s\n", currentId)
		}
		//tSlave := slaveClock.GetTime()
		
		n, _, err = conn.ReadFrom(buf) 
		if err != nil {
			log.Fatal(err)
		}	
		
		// Look for FOLLOW_UP and check if id is the same as in SYNC message
		if strings.Compare(protocol.FOLLOW_UP, string(buf[:len(protocol.FOLLOW_UP)])) == 0 {
		
			idNbDigits, _ := strconv.Atoi(currentId)
			idNbDigits = idNbDigits / 10 + 1
			if strings.Compare(currentId, string(buf[n - idNbDigits:n])) == 0 {
				unixTime, _ := strconv.Atoi(string(buf[len(protocol.FOLLOW_UP) : n - idNbDigits]))
				log.Printf("Received FOLLOW_UP correctly with Unix time : %s", clock.ToString(unixTime))
			}
		}
		
		//fmt.Println(string(buf[:n]))
	}
}

// TODO : function readMulticast()

func delayRequest() {
}
	