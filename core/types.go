package core

import (
	"time"
)

type NodeId int

type activeTimer struct {
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

type Message interface{}
type TimerMessage struct{ TimerName string }
