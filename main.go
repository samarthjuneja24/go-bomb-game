package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Assigning the value as a Random number, can be changed as per convenience
	numberOfSecondsGameShouldRunFor := rand.Intn(10) + 1
	totalNumberOfPlayers := 10

	var listOfPlayers []int
	for i := 1; i <= totalNumberOfPlayers; i++ {
		listOfPlayers = append(listOfPlayers, i)
	}

	// Creating two channels
	// inputChannel takes input, switches the sender and receiver and sends it to the outputChannel.
	// outputChannel checks for context timeout. If timed out, it shows who had the bomb when it exploded,
	// otherwise it passes the context back to inputChannel
	inputChannel := make(chan context.Context)
	outputChannel := make(chan context.Context)

	//inputChannel listener go routine started
	go Play(inputChannel, outputChannel, listOfPlayers)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(numberOfSecondsGameShouldRunFor)*time.Second)
	ctx = context.WithValue(ctx, "sender", 1)
	ctx = context.WithValue(ctx, "receiver", PlayerToPassTo(1, listOfPlayers))
	defer cancel()
	inputChannel <- ctx

	for {
		//Waiting for response sent from inputChannel for loop
		processedCtx := <-outputChannel
		if err := processedCtx.Err(); err != nil {
			fmt.Printf("Player eliminated: %v\n", processedCtx.Value("receiver"))
			break
		}
		inputChannel <- processedCtx.(context.Context)
	}
	close(inputChannel)
}

func Play(inputChannel <-chan context.Context, outputChannel chan<- context.Context, listOfPlayers []int) {
	for ctx := range inputChannel {
		newSender := ctx.Value("receiver").(int)
		newReceiver := PlayerToPassTo(newSender, listOfPlayers)
		newCtx := context.WithValue(ctx, "sender", newSender)
		newCtx = context.WithValue(newCtx, "receiver", newReceiver)
		fmt.Printf("Passing to player %v\n", newReceiver)
		outputChannel <- newCtx
	}
	close(outputChannel)
}

func PlayerToPassTo(passingPlayer int, playersList []int) int {
	randomPosition := rand.Intn(len(playersList))

	//If the same player comes up, do repeat the function else return the new player's number
	if randomPosition != passingPlayer {
		return randomPosition
	}
	return PlayerToPassTo(passingPlayer, playersList)
}
