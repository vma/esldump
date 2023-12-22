package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vma/esl"
	"github.com/vma/getopt"
)

type Handler struct{}

var (
	subscribedEvents = "CHANNEL_ANSWER CHANNEL_HANGUP"
	events           = getopt.ListLong("event", 'e', "events to capture, may be used multiple times (default: "+subscribedEvents+")")
	eslHost          = getopt.StringLong("host", 'H', "127.0.0.1", "freeswitch ESL host")
	eslPort          = getopt.IntLong("port", 'p', 8021, "freeswitch ESL port")
)

func main() {
	getopt.SetParameters("")
	getopt.Parse()
	for {
		log.Printf("connecting to freeswitch %s:%d", *eslHost, *eslPort)
		con, err := esl.NewConnection(fmt.Sprintf("%s:%d", *eslHost, *eslPort), &Handler{})
		if err != nil {
			log.Printf("unable to connect to freeswitch: %v, retrying...", err)
			time.Sleep(4 * time.Second)
			continue
		}
		if len(*events) > 0 {
			subscribedEvents = strings.Join(*events, " ")
		}
		log.Printf("connected, waiting for events...")
		con.HandleEvents()
		log.Printf("disconnected from esl, reconnecting...")
		time.Sleep(4 * time.Second)
	}
}

func (h *Handler) OnDisconnect(c *esl.Connection, e *esl.Event) {}

func (h *Handler) OnClose(c *esl.Connection) {}

func (h *Handler) OnConnect(con *esl.Connection) {
	if _, err := con.SendRecv("event", "plain", subscribedEvents); err != nil {
		log.Printf("ERR: unable to subscribe to events: %v", err)
	}
}

func (h *Handler) OnEvent(con *esl.Connection, ev *esl.Event) {
	fmt.Println(ev)
}
