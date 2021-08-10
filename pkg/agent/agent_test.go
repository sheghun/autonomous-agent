package agent

import (
	"context"
	"testing"
)

func TestAutonomousAgent_Integration(t *testing.T) {
	t.Parallel()

	handlers := make(map[string]AutonomousAgentHandler)
	var behaviours []AutonomousAgentBehaviour

	agent1 := NewAutonomousAgent(handlers, behaviours, context.Background())

	// Initialize a buffered channel to avoid blocking

	inbox := make(chan Msg, 1)
	msg := Msg{
		Type:    "hello",
		Message: "hello world",
	}

	agent1.SetInbox(inbox)
	inbox <- msg


	retrievedInbox := agent1.GetInbox()
	msg1 := <-retrievedInbox

	if msg.Message != msg1.Message {
		t.Fatalf("actual: %s not equal to expected: %s", msg1.Message, msg.Message)
	}

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
