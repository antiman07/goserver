package gorillaws

import (
	"trunk/cellnet"
	"trunk/cellnet/proc"
)

func init() {

	transmitter := new(WSMessageTransmitter)
	hooker := new(MsgHooker)

	proc.RegisterProcessor("gorillaws.ltv", func(bundle proc.ProcessorBundle, userCallback cellnet.EventCallback) {

		bundle.SetTransmitter(transmitter)
		bundle.SetHooker(hooker)
		bundle.SetCallback(proc.NewQueuedEventCallback(userCallback))

	})
}
