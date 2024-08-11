package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"fmt"
	"log"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicHealth(snapshot agent.GameSnapshot) int {
	totalHealth := 0
	allSnakeStats := ""
	for _, allySnake := range snapshot.YourTeam() {
		totalHealth += allySnake.Health()
		allSnakeStats += fmt.Sprintf("[ID: %s, Health: %d] ", allySnake.ID(), allySnake.Health())
	}
	log.Printf("Turn %d - Allies [ID, Health]*: %s\n", snapshot.Turn(), allSnakeStats)
	return totalHealth
}
