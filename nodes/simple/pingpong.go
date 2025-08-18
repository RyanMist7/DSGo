package simple

import (
	"DSGo/core"
)

type PingPongNode struct {
	ctx  core.NodeContext
	peer core.NodeId
}

func (n *PingPongNode) Init(ctx core.NodeContext) {
	n.ctx = ctx
}

func (n *PingPongNode) HandleMessage(msg core.Message) {
	switch m := msg.(type) {
	case string:
		n.ctx.Log().Info("Received message: %s", m)
		if n.peer == 0 {
			return
		}
		if m == "ping" {
			n.ctx.Send(n.peer, "pong")
		}
	}
}

func (n *PingPongNode) HandleTimer(timer core.NodeTimer) {
	n.ctx.Log().Info("Timer fired: %s", timer.Message.TimerName)
}

func (n *PingPongNode) SetPeer(peer core.NodeId) {
	n.peer = peer
}

func (n *PingPongNode) SendMessageToPeer(msg core.Message) {
	if n.peer != 0 {
		n.ctx.Send(n.peer, msg)
	}
}
