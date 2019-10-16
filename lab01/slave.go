package main

import (

	"log"
	"net"
	"runtime"
	"strings"
	"strconv"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/ipv4"
	"PRR_Lab01/lab01/protocol"
	"PRR_Lab01/lab01/clock"
)

var c = make(chan int)
var slaveClock = clock.New(rand.Intn(clock.MAX_OFFSET))
var mutex sync.Mutex

func main() {
	
	rand.Seed(time.Now().UnixNano()) // Initialize rand
	masterReader()

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
		interf, _ = net.InterfaceByName("en0")
	}

	if err = p.JoinGroup(interf, addr); err != nil { // listen on ip multicast
		log.Fatal(err)
	}
	
	go delayRequest()
	
	var currentId string
	var readyForDelayRequest bool = true
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
		
		n, _, err = conn.ReadFrom(buf) 
		if err != nil {
			log.Fatal(err)
		}	
		
		// Register master time after receiving SYNC
		idNbDigits := len(currentId)
		tMaster, _ := strconv.Atoi(string(buf[len(protocol.FOLLOW_UP) : n - idNbDigits]))
		
		// Look for FOLLOW_UP and check if id is the same as in SYNC message
		if strings.Compare(protocol.FOLLOW_UP, string(buf[:len(protocol.FOLLOW_UP)])) == 0 {
			if strings.Compare(currentId, string(buf[n - idNbDigits:n])) == 0 {
				log.Printf("Received FOLLOW_UP with correct id\n")
			}
		}
		
		// latence (simulate time requested to get message from master)
		time.Sleep(protocol.LATENCE * time.Second)

		// First time correction (t_master - t_slave)
		tSlave := slaveClock.GetTime()
		log.Printf("Slave time before first correction : %s\n", clock.ToString(tSlave));
		
		gap := tMaster - tSlave

		mutex.Lock()
		slaveClock.SetOffset(gap)
		mutex.Unlock()
		log.Printf("Slave time after first correction : %s\n", clock.ToString(slaveClock.GetTime()));
				
		if readyForDelayRequest == true {
			// Ready to start delayRequest routine after first turn
			c <- 1
			readyForDelayRequest = false
		}
	}
}

// Second step (point to point UDP)
func delayRequest() {

	// Has to wait until master sent SYNC, FOLLOW_UP at least once
	<- c
	conn, err := net.Dial("udp", protocol.MASTER_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
			
	buf := make([]byte, 1024)
	
	// Generate a random slave id (Here, we just have 4 digit)
	var slaveId = id();
	for { 

		tSlave := slaveClock.GetTime() // Local delayRequest time
		
		// Simulate latency : time requested for send request + get response (2 * latency time)
		time.Sleep(2 * protocol.LATENCE * time.Second)
		
		payload := protocol.DELAY_REQUEST + strconv.Itoa(slaveId)
		_, _ = conn.Write([]byte(payload))
		log.Printf("sent DELAY_REQUEST\n")
		
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		
		// Look for delay_response and check if id is correct 
		if strings.Compare(protocol.DELAY_RESPONSE, string(buf[:len(protocol.DELAY_RESPONSE)])) == 0 {
			if strings.Compare(strconv.Itoa(slaveId), string(buf[n - protocol.SLAVE_ID_LENGTH : n])) == 0 {
				log.Printf("Received DELAY_RESPONSE with correct id\n")
			}
		}
		
		// Second time correction
		tMaster, _ := strconv.Atoi(string(buf[len(protocol.DELAY_RESPONSE) : n - protocol.SLAVE_ID_LENGTH]))
		delay := (tMaster - tSlave) / 2
		
		mutex.Lock()
		slaveClock.SetOffset(delay)
		mutex.Unlock()
		
		// Here should be random wait, but just fixed short time for visualisation
		time.Sleep(4 * protocol.K * time.Second)
	}
}

// Generate a 4 digits random id
func id() int {
	id :=
	rand.Intn(10) * 1000 +
	rand.Intn(10) * 100 +
	rand.Intn(10) * 10 + 
	rand.Intn(10)
	
	return id
}	