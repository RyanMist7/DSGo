package core

import (
	"sync"
)

type Node interface {
	Init(ctx NodeContext)
	HandleMessage(msg Message)
	HandleTimer(timer NodeTimer)
}

type Core struct {
	mu          sync.Mutex
	nodes       map[NodeId]Node                // map of node id --> node
	channels    map[NodeId]chan Message        // map of node id --> channels
	timers      map[NodeId]map[int]activeTimer // map of node id --> map of timer id ->  list of active timers
	nextTimerId map[NodeId]int                 // map of node id -> next timer id
	// TODO: eventually add other configurable stuff here
	// (1) message latency
	// (2) message loss / duplication / reordering (might want to be opinionated on duplication & reordering)
	// (2) node crash / recovery rate
	// (3) node partition rate
}

func NewCore() *Core {
	return &Core{
		nodes:    make(map[NodeId]Node),
		channels: make(map[NodeId]chan Message),
		timers:   make(map[NodeId]map[int]activeTimer),
	}
}

func (c *Core) RegisterNode(id NodeId, node Node) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.nodes[id]; exists {
		return
	}

	c.nodes[id] = node
	c.channels[id] = make(chan Message, 10)
	c.timers[id] = make(map[int]activeTimer)
	c.nextTimerId[id] = 1

	ctx := &nodeContext{
		core:   c,
		nodeId: id,
	}

	node.Init(ctx)
	go c.runNode(id, node)
}

// go routine that runs for each node
// this pull messages from channel and calls the nodes handle message
func (c *Core) runNode(id NodeId, node Node) {
	ch := c.channels[id]
	for msg := range ch {
		switch m := msg.(type) {
		case TimerMessage:
			node.HandleTimer(NodeTimer{message: m})
		case Message:
			node.HandleMessage(m)
		default:
			panic("run node getting non allowed message")
		}
	}
}
