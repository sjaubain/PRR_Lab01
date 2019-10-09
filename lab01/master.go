package main

import (

	"time"
	"net"
	"log"
	"strconv"
	
	"lab01/protocol"
	"lab01/clock"
)

const k = 4

var masterClock = clock.New(0)

func main() {

	conn, err := net.Dial("udp", protocol.MULTICAST_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var id = 0
	for {
		// SYNC
		payload := protocol.SYNC + strconv.Itoa(id)
		_, err := conn.Write([]byte(payload))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		log.Println("sent SYNC to multicast address")
		
		// FOLLOW_UP
		tMaster := masterClock.GetTime()
		payload = protocol.FOLLOW_UP + strconv.Itoa(tMaster) + strconv.Itoa(id)
		_, err = conn.Write([]byte(payload))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		log.Println("sent FOLLOW_UP to multicast address")
		
		time.Sleep(k * time.Second)
		id++
	}
}

// TODO : function sendMulticast()
