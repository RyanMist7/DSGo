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
	logger NodeLogger
}

type NodeTimer struct {
	Message TimerMessage
}

type NodeLogger struct {
	nodeId NodeId
}

type Message interface{}
type TimerMessage struct{ TimerName string }
