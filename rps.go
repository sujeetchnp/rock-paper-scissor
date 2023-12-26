package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sujeetchnp/rock-paper-scissor/model"
	"github.com/sujeetchnp/rock-paper-scissor/service"
)

// RockPaperScissor struct
type RockPaperScissor struct {
	CountPlayer1  int
	CountPlayer2  int
	Player1Choice model.Choice
	Player2Choice model.Choice
	WinningChoice model.Choice
	PlayerName    string

	CacheService service.CacheService
	GameService  service.GameService
}

func NewRockPaperScissor(cacheService service.CacheService, gameService service.GameService) *RockPaperScissor {
	return &RockPaperScissor{
		CacheService: cacheService,
		GameService:  gameService,
	}
}

func (rps *RockPaperScissor) computerChoice() {
	// Seed the random number generator
	(time.Now().UnixNano())

	// Generate a random number between 0 and 2
	choice := model.Choice(rand.Intn(3))

	// Assign the random choice to player2Choice
	rps.Player2Choice = choice

	// Optionally, you can print out the computer's choice
	fmt.Println("Computer chooses:", choice.String())
}

func (rps *RockPaperScissor) Run() {
	var playerInput string

	for {
		fmt.Println("Enter your choice (rock, paper, scissor) or 'q' to quit:")

		// Get player input
		fmt.Scanln(&playerInput)

		// Check for quit command
		if playerInput == "q" {
			break
		}

		// Convert player input to Choice
		player1Choice, err := model.GetChoice(playerInput)
		if err != nil {
			fmt.Println("Invalid choice. Please enter rock, paper, or scissor.")
			continue
		}

		// Computer choice (random)
		rps.computerChoice()

		// Compare choices and determine winner
		rps.Player1Choice = player1Choice
		winningChoice := rps.GameService.ComparePlayerChoice(rps.Player1Choice, rps.Player2Choice)

		// Update and display round results
		if winningChoice == rps.Player1Choice {
			fmt.Println("You win this round!")
			rps.CountPlayer1++
		} else if winningChoice == rps.Player2Choice {
			fmt.Println("Computer wins this round!")
			rps.CountPlayer2++
		} else {
			fmt.Println("This round is a draw!")
		}

		// Display scores
		fmt.Printf("Current Score - You: %d, Computer: %d\n", rps.CountPlayer1, rps.CountPlayer2)
	}

	// End of game summary
	fmt.Printf("Final Score - You: %d, Computer: %d\n", rps.CountPlayer1, rps.CountPlayer2)
	if rps.CountPlayer1 > rps.CountPlayer2 {
		fmt.Println("You are the final winner!")
	} else if rps.CountPlayer1 < rps.CountPlayer2 {
		fmt.Println("Computer is the final winner!")
	} else {
		fmt.Println("The game ended in a draw!")
	}
}
