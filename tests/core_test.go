package tests

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"DSGo/core"
	"DSGo/nodes/simple"
)

func TestDSGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ping Pong Test")
}

var _ = Describe("Ping-Pong Node Test", func() {
	It("should exchange messages between two nodes within 500ms", func() {
		c := core.NewCore()

		ping := &simple.PingPongNode{}
		pong := &simple.PingPongNode{}

		c.RegisterNode(1, ping)
		c.RegisterNode(2, pong)

		ping.SetPeer(2)
		pong.SetPeer(1)

		ping.SendMessageToPeer("ping")

		Eventually(func() bool {
			// we dont expect anything, just the relevant print messages
			return true
		}, 500*time.Millisecond, 10*time.Millisecond).Should(BeTrue())
	})
})
