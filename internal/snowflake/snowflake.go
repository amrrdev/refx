package snowflake

import (
	"sync"
	"time"
)

const (
	epoch       = int64(1700000000000) // custom epoch (Nov 2023)
	machineBits = 10
	seqBits     = 12

	maxMachineID = -1 ^ (-1 << machineBits)
	maxSequence  = -1 ^ (-1 << seqBits)

	timeShift    = machineBits + seqBits
	machineShift = seqBits
)

type Generator struct {
	mu        sync.Mutex
	lastTime  int64
	sequence  int64
	machineID int64
}

func New(machineID int64) *Generator {
	if machineID < 0 || machineID > maxMachineID {
		panic("invalid machine id")
	}
	return &Generator{machineID: machineID}
}

func (g *Generator) NextID() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixMilli()

	if now == g.lastTime {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			for now <= g.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		g.sequence = 0
	}

	g.lastTime = now

	return ((now - epoch) << timeShift) |
		(g.machineID << machineShift) |
		g.sequence
}
