/* 
 * File:   protocol.go
 * Author: Teklehaimanot Robel - Jobin Simon 
 *
 * All constants needed in slave.go and master.go
 */

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

/**
*
* Multicast packet master -> slaves :
*
*	[ SYNC | ID ] and [ FOLLOW_UP | MASTER_TIME | ID ]
*
* Point to point packet slave -> master : 
*
*   [ DELAY_REQUEST | SLAVE_ID ]
*
* Point to point packet master -> slave : 
*
*   [ DELAY_RESPONSE | MASTER_TIME | SLAVE_ID ]
*
*/