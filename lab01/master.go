package main

import (

	"time"
	"net"
	"log"
	"strconv"

	"lab01/protocol"
	"lab01/clock"
)

var masterClock = clock.New(0)

func main() {

	multicast()
}

func multicast() {

	conn, err := net.Dial("udp", protocol.MULTICAST_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	go slaveReader()
	
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
		
		time.Sleep(protocol.K * time.Second)
		id++
	}
}

func slaveReader() {
	
	// Point to point communication
	conn, err := net.ListenPacket("udp", protocol.MASTER_ADDR) // listen on port
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	

	buf := make([]byte, 1024)
	for {
	
		n, slaveAddr, err := conn.ReadFrom(buf) 
		if err != nil {
			log.Fatal(err)
		}
		
		tMaster := masterClock.GetTime()
		slaveId := string(buf[n - protocol.SLAVE_ID_LENGTH : n])
		payload := protocol.DELAY_RESPONSE + strconv.Itoa(tMaster) + slaveId
		log.Printf("received DELAY_REQUEST from slave : %s\n", slaveId)
		_, err = conn.WriteTo([]byte(payload), slaveAddr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		log.Printf("sent DELAY_RESPONSE to slave : %s\n", slaveId)
	}
}

