package agent

import (
	"context"
	"fmt"
	"log"
)

type Msg struct {
	Type    string
	Message string
}

type AutonomousAgentBehaviour = func(outbox chan Msg, errCh chan error, ctx context.Context)
type AutonomousAgentHandler = func(msg Msg, errCh chan error, ctx context.Context)

type AutonomousAgent struct {
	handlers   map[string]AutonomousAgentHandler
	behaviours []AutonomousAgentBehaviour
	ctx        context.Context
	errCh      chan error
	inbox      chan Msg
	outbox     chan Msg
}

func NewAutonomousAgent(
	handlers map[string]AutonomousAgentHandler,
	behaviours []AutonomousAgentBehaviour,
	ctx context.Context,
) *AutonomousAgent {

	outbox := make(chan Msg)
	errCh := make(chan error)

	agent := &AutonomousAgent{
		handlers:   handlers,
		behaviours: behaviours,
		ctx:        ctx,
		errCh:      errCh,
		inbox:      nil,
		outbox:     outbox,
	}

	return agent
}

func (agent *AutonomousAgent) Run() error {

	if agent.inbox == nil {
		return fmt.Errorf("inbox channel not set please call SetInbox() on the agent to set the inbox channel")
	}

	go agent.consumeInbox()
	go agent.runBehaviours()

	return <-agent.errCh
}

func (agent *AutonomousAgent) SetInbox(inbox chan Msg) *AutonomousAgent {
	agent.inbox = inbox

	return agent
}

func (agent *AutonomousAgent) GetOutbox() chan Msg {
	return agent.outbox
}

func (agent *AutonomousAgent) GetInbox() chan Msg {
	return agent.inbox
}

// Consumes inbox and emits to outbox
func (agent *AutonomousAgent) consumeInbox() {
	for msg := range agent.inbox {

		log.Printf("recived message of type: %s and message: %s", msg.Type, msg.Message)

		go func() {
			// Get the handler for this particular message type
			if handler, ok := agent.handlers[msg.Type]; ok {
				select {
				case <-agent.ctx.Done():
					log.Printf("unsubcribing from inbox")
					return

				default:
					go handler(msg, agent.errCh, agent.ctx)
				}
			}
		}()
	}
}

func (agent *AutonomousAgent) runBehaviours() {
	for _, b := range agent.behaviours {
		select {
		case <-agent.ctx.Done():
			return

		default:
			go b(agent.outbox, agent.errCh, agent.ctx)

		}
	}
}
