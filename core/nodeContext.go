package core

import (
	"fmt"
	"time"
)

type NodeId int

type timer struct {
	message TimerMessage
	timer   *time.Timer
	timerId int
}

type nodeContext struct {
	core   *Core
	nodeId NodeId
}

type NodeTimer struct {
	message TimerMessage
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

func (ctx *nodeContext) SetTimer(nodeTimer NodeTimer, delay time.Duration) {
	ctx.core.mu.Lock()
	defer ctx.core.mu.Unlock()
	timerId := ctx.core.nextTimerId[ctx.nodeId]
	ctx.core.nextTimerId[ctx.nodeId]++

	t := time.AfterFunc(delay, func(id int) func() {
		return func() {
			ctx.core.channels[ctx.nodeId] <- nodeTimer.message

			ctx.core.mu.Lock()
			delete(ctx.core.timers[ctx.nodeId], id)
			ctx.core.mu.Unlock()
		}
	}(timerId))

	timer := timer{
		message: nodeTimer.message,
		timer:   t,
		timerId: timerId,
	}
	ctx.core.timers[ctx.nodeId][timerId] = timer

}

func (ctx *nodeContext) Log(level string, format string, args ...any) {
	fmt.Printf("[n%d][%s] %s\n", ctx.nodeId, level, fmt.Sprintf(format, args...))
}
