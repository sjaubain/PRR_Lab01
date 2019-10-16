package protocol

const MULTICAST_ADDR = "224.0.0.1:6666"
const MASTER_ADDR = "127.0.0.1:6000"
const FOLLOW_UP = "FOLLOW_UP"
const SYNC = "SYNC"
const DELAY_REQUEST = "DELAY_REQUEST"
const DELAY_RESPONSE = "DELAY_RESPONSE"
const K = 4
const SLAVE_ID_LENGTH = 4 // We decided to assign a 4 digit random id to slaves
const LATENCE = 5 // Simulated latency time in seconds