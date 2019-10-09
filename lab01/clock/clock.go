package clock

import (
	"time"
)

const MAX_OFFSET = 5000

type clock struct {

	// To simulate desynchronization
	offset int
}

func New(offset int) clock {
	c := clock {offset}
	return c
}

// Return current local time in milliseconds
func (c clock) GetTimeMillis() int {
	return int (time.Now().Unix()) + c.offset
}
	
func (c clock) ToString(timeMillis int) string {
	return "TODO : time to string"
}
	