package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sheghun/autonomous-agent/pkg/agent"
	"github.com/sheghun/autonomous-agent/pkg/behaviour"
	"github.com/sheghun/autonomous-agent/pkg/handler"
)

func main() {
	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithCancel(context.Background())

	handlers := make(map[string]agent.AutonomousAgentHandler)
	handlers["hello"] = handler.ListenForHello

	behaviours := []agent.AutonomousAgentBehaviour{
		behaviour.GenerateRandomMessages,
	}

	agent1 := agent.NewAutonomousAgent(handlers, behaviours, ctx)
	agent2 := agent.NewAutonomousAgent(handlers, behaviours, ctx)

	// Set the outbox of each agents as their inbox
	agent1 = agent1.SetInbox(agent2.GetOutbox())
	agent2 = agent2.SetInbox(agent1.GetOutbox())

	go func() {
		log.Printf("starting agent1")
		if err := agent1.Run(); err != nil {
			log.Printf("error running agent 1 failed: %v", err)
			cancel()
		}
	}()

	go func() {
		log.Printf("starting agent2")
		if err := agent2.Run(); err != nil {
			log.Printf("error running agent2 failed: %v", err)
			cancel()
		}
	}()

	handleExit(ctx, cancel)

}

func handleExit(ctx context.Context, cancel context.CancelFunc) {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)

	// Wait for one of the channels to receive a message
	select {
	case <-ctx.Done():

	case <-signals:
	}

	log.Printf("Exiting applicatons and shutting down agents")

	cancel()

	time.Sleep(1 * time.Second ) // Wait for go routines to exit.
}
