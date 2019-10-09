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
		
		// FOLLOW_UP
		//tMaster := masterClock.GetTimeMillis()
		log.Println("sent SYNC to multicast address\n")
		
		// Convert timeMillis (int) to byte array
		//_, err = conn.Write([]byte(strconv.Itoa(tMaster)))
		
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		
		time.Sleep(k * time.Second)
		id++
	}
}

