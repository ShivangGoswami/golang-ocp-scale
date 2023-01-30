package message

import (
	"math/rand"
	"time"
)

var (
	messages = []string{
		"Time is money",
		"It's time for a beer",
		"It's time for two beers! Or three!",
		"Time it's relative",
		"Good time to ride a bike",
		"So much time",
		"It's five o'clock somewhere",
		"Time to relax",
		"It's time to get out of here",
	}
	lengthofmessages = len(messages)
)

var rdgen *rand.Rand

func init() {
	seed := rand.NewSource(time.Now().UnixNano())
	rdgen = rand.New(seed)
}

func GetMessage() string {
	return messages[rdgen.Intn(lengthofmessages)]
}
