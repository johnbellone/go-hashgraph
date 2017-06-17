package hashgraph

import (
	"crypto"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Hashgraph struct {
	events map[string]*Event
	nodes  []Node

	logger *log.Logger
}

func Create(c *Config) (*Hashgraph, error) {
	h := &Hashgraph{
		config.logger,
	}
	return h, nil
}

func (h *Hashgraph) Start() {
	rand.Seed(time.Now().Unix())

	node = h.nodes[rand.Intn(len(h.nodes))]
}

func (h *Hashgraph) Shutdown() {
}

func (h *Hashgraph) publishEvent() {
	s := crypto.Sign()
}

func (h *Hashgraph) divideRounds(hashes []string) {
	for _, hash := range hashes {
		ev, ok := h.events[hash]
		if !ok {
			// TODO: This should *not* be reachable. Since
			// we should add new events to graph prior // to
			// running DivideRounds.
		}
	}
}

func (h *Hashgraph) decideFame() {}
