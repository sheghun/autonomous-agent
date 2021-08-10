package behaviour

import (
	"context"
	"math/rand"
	"time"

	"github.com/sheghun/autonomous-agent/pkg/agent"
)

func GenerateRandomMessages(outbox chan agent.Msg, errCh chan error, ctx context.Context) {
	words := []string{
		"hello",
		"sun",
		"world",
		"space",
		"moon",
		"crypto",
		"sky",
		"ocean",
		"universe",
		"human",
	}

	str := ""

	// After every two seconds
	for _ = range time.Tick(2 * time.Second) {
		str = ""
		str += words[rand.Intn(len(words))] + " "
		str += words[rand.Intn(len(words))]

		// Send the message to the channel.
		outbox <-  agent.Msg{
			Type: "hello",
			Message: str,
		}
	}
}
