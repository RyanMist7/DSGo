package core

import (
	"fmt"
	"time"
)

// NodeContext is the API Core gives to each node
type NodeContext interface {
	Send(destID NodeId, msg Message)
	SetTimer(timer NodeTimer, delay time.Duration)
	Log() *NodeLogger
}

func (ctx *nodeContext) Send(destId NodeId, msg Message) {
	go func() {
		// eventually we can add injectable failures here
		// we would likely need a lock as i imagine somethings like partitions would be a map
		ch, exists := ctx.core.channels[destId]
		if exists {
			ch <- msg
		}
	}()
}

func (ctx *nodeContext) Broadcast(nodeIDs []NodeId, msg Message) {
	for _, id := range nodeIDs {
		ctx.Send(id, msg)
	}
}

func (ctx *nodeContext) SetTimer(nodeTimer NodeTimer, delay time.Duration) {
	ctx.core.mu.Lock()
	defer ctx.core.mu.Unlock()
	timerId := ctx.core.nextTimerId[ctx.nodeId]
	ctx.core.nextTimerId[ctx.nodeId]++

	t := time.AfterFunc(delay, func() {
		ctx.core.channels[ctx.nodeId] <- nodeTimer.message
		ctx.core.mu.Lock()
		delete(ctx.core.timers[ctx.nodeId], timerId)
		ctx.core.mu.Unlock()
	})

	timer := activeTimer{
		message: nodeTimer.message,
		timer:   t,
		timerId: timerId,
	}
	ctx.core.timers[ctx.nodeId][timerId] = timer

}

func (ctx *nodeContext) Log() *NodeLogger {
	return &ctx.logger
}

func (l *NodeLogger) Info(format string, args ...any) {
	fmt.Printf("[n%d][INFO] %s\n", l.nodeId, fmt.Sprintf(format, args...))
}
func (l *NodeLogger) Warn(format string, args ...any) {
	fmt.Printf("[n%d][WARN] %s\n", l.nodeId, fmt.Sprintf(format, args...))
}
func (l *NodeLogger) Error(format string, args ...any) {
	fmt.Printf("[n%d][ERROR] %s\n", l.nodeId, fmt.Sprintf(format, args...))
}
