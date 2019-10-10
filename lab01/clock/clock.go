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

// Return current local time in second elapsed since 1st january 1970
func (c clock) GetTime() int {
	return int (time.Now().Unix()) + c.offset
}
	
func (c clock) SetOffset(deltaT int) {
	c.offset = deltaT
}

func ToString(unixTime int) string {
	return time.Unix(int64 (unixTime), 0).String()
}
	