package core

import (
	"time"
)

type Message interface{}

// NodeContext is the API Core gives to each node
type NodeContext interface {
	Send(destID string, msg Message)
	SetTimer(timerID string, delay time.Time)
	Log(level string, format string, args ...any)
}

type Node interface {
	Init(ctx NodeContext)
	HandleMessage(msg Message)
	HandleTimer(timerID string)
}
