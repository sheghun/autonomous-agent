package handler

import (
	"context"
	"log"
	"strings"

	"github.com/sheghun/autonomous-agent/pkg/agent"
)

// ListenForHello that listens for hello in a message
func ListenForHello(msg agent.Msg, errCh chan error, ctx context.Context) {
	if strings.Contains(msg.Message, "hello") {
		log.Printf("handler found message: %v", msg.Message)
	}
}
