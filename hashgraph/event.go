package hashgraph

import (
	"math/big"
	"time"
)

type Event struct {
	Parent    *Event
	Payload   []byte
	Signature big.Int
	Time      time.Time
}
