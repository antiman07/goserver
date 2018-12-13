package tests

import (
	"fmt"
	"github.com/cellnet"
	"github.com/cellnet/peer"
	_ "github.com/cellnet/peer/tcp"
	_ "github.com/cellnet/peer/udp"
	"github.com/cellnet/proc"
	_ "github.com/cellnet/proc/tcp"
	_ "github.com/cellnet/proc/udp"
	"github.com/cellnet/util"
	"testing"
	"time"
)

type echoContext struct {
	Address   string
	Protocol  string
	Processor string
	Tester    *util.SignalTester
	Acceptor  cellnet.GenericPeer
}

var (
	echoContexts = []*echoContext{
		{
			Address:   "127.0.0.1:7701",
			Protocol:  "tcp",
			Processor: "tcp.ltv",
		},
		{
			Address:   "127.0.0.1:7702",
			Protocol:  "udp",
			Processor: "udp.ltv",
		},
	}
)

func echo_StartServer(context *echoContext) {
	queue := cellnet.NewEventQueue()

	context.Acceptor = peer.NewGenericPeer(context.Protocol+".Acceptor", context.Protocol+"server", context.Address, queue)

	proc.BindProcessorHandler(context.Acceptor, context.Processor, func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			fmt.Println("server accepted")
		case *TestEchoACK:

			fmt.Printf("server recv %+v\n", msg)

			ev.Session().Send(&TestEchoACK{
				Msg:   msg.Msg,
				Value: msg.Value,
			})

		case *cellnet.SessionClosed:
			fmt.Println("session closed: ", ev.Session().ID())
		}

	})

	context.Acceptor.Start()

	queue.StartLoop()
}

func echo_StartClient(echoContext *echoContext) {
	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer(echoContext.Protocol+".Connector", echoContext.Protocol+"client", echoContext.Address, queue)

	proc.BindProcessorHandler(p, echoContext.Processor, func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
			ev.Session().Send(&TestEchoACK{
				Msg:   "hello",
				Value: 1234,
			})
		case *TestEchoACK:

			fmt.Printf("client recv %+v\n", msg)

			echoContext.Tester.Done(1)

		case *cellnet.SessionClosed:
			fmt.Println("client closed")
		}
	})

	p.Start()

	queue.StartLoop()

	echoContext.Tester.WaitAndExpect("not recv data", 1)
}

func runEcho(t *testing.T, index int) {

	ctx := echoContexts[index]

	ctx.Tester = util.NewSignalTester(t)
	ctx.Tester.SetTimeout(time.Hour)

	echo_StartServer(ctx)

	echo_StartClient(ctx)

	ctx.Acceptor.Stop()
}

func TestEchoTCP(t *testing.T) {

	runEcho(t, 0)
}

func TestEchoUDP(t *testing.T) {

	runEcho(t, 1)
}
