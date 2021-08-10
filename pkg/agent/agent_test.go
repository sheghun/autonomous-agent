package agent

import (
	"context"
	"sync"
	"testing"
)

func TestAutonomousAgent_Integration(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	wg.Add(1)

	msg := Msg{
		Type:    "hello",
		Message: "hello world",
	}

	// Write handler to run test
	 testHandler := func(msg1 Msg, errCh chan error, ctx context.Context) {
	 	defer wg.Done()

		 if msg.Message != msg1.Message {
			 t.Fatalf("actual: %s not equal to expected: %s", msg1.Message, msg.Message)
		 }
	}

	handlers := make(map[string]AutonomousAgentHandler)
	handlers[msg.Type] = testHandler

	var behaviours []AutonomousAgentBehaviour

	agent1 := NewAutonomousAgent(handlers, behaviours, context.Background())


	// Initialize a buffered channel to avoid blocking
	inbox := make(chan Msg, 1)

	agent1.SetInbox(inbox)
	inbox <- msg

		go func() {
			if err := agent1.Run(); err != nil {
				t.Fatalf("failed running agent: %v", err)
			}
		}()

	wg.Wait()
}

func TestUnitAutonomousAgent_GetInboxAndGetOutbox(t *testing.T) {
	handlers := make(map[string]AutonomousAgentHandler)
	var behaviours []AutonomousAgentBehaviour

	agent1 := NewAutonomousAgent(handlers, behaviours, context.Background())

	inbox := make(chan Msg)

	agent1.SetInbox(inbox)

	if agent1.GetInbox() == nil {
		t.Fatalf("SetInbox() failed and wasn't successful")
	}

	if agent1.GetOutbox() == nil {
		t.Fatalf("GetOutbox() failed and wasn't successful")
	}
}
