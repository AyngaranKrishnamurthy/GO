package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameState interface {
	executeState(*GameContext) bool //Accepts one of the contexts from GameContext and returns true if the game has finished, false if not.
}

type GameContext struct { //Contains all the game states
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState //Contains the current state of the game
}

type StarterState struct{}

func (s *StarterState) executeState(c *GameContext) bool { //Implements GameState struct
	c.Next = &AskState{} //Asks the player for a number to guess

	rand.Seed(time.Now().UnixNano()) //Generates random numbers of type int64
	c.SecretNumber = rand.Intn(10)   //Returns as integer between 0 and specified number(10)

	fmt.Println("Enter the number of retries to set difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", c.Retries) //io.Reader implemented to fetch data

	return true
}

type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool { //Implements GameState struct
	fmt.Printf("Guess a number between 0 and 10. Tries Left: %d", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n) //Guess of the player
	c.Retries = c.Retries - 1

	if n == c.SecretNumber { //Checks if the guess is right
		c.Won = true
		c.Next = &FinishState{}
	}
	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true //To continue the game it returns true to the executeState method
}

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		fmt.Println("Congratulations, you won.")
	} else {
		fmt.Println("Sorry, You didn't win the game \n")
		fmt.Println("The correcct number was: %d", c.SecretNumber)
	}
	return false
}

func main() {
	start := StarterState{}
	game := GameContext{
		Next: &start, //Setting next state as pointer to the start variable
	}
	for game.Next.executeState(&game) {
	} //A loop without any statement, executes as long as executeState is true
}
