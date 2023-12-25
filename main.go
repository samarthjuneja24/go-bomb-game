package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number in the range of 1-10
	// The timeout will be for that many seconds
	randomNumber := rand.Intn(10) + 1

	// Creating two channels
	// inputChannel takes input, switches the sender and receiver and sends it to the outputChannel.
	// outputChannel checks for context timeout. If timed out, it shows who had the bomb when it exploded,
	// otherwise it passes the context back to inputChannel
	inputChannel := make(chan context.Context)
	outputChannel := make(chan context.Context)

	//inputChannel listener go routine started
	go Play(inputChannel, outputChannel)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(randomNumber)*time.Second)
	ctx = context.WithValue(ctx, "sender", "player1")
	ctx = context.WithValue(ctx, "receiver", "player2")
	defer cancel()
	inputChannel <- ctx

	for {
		//Waiting for response sent from inputChannel for loop
		processedCtx := <-outputChannel
		if err := processedCtx.Err(); err != nil {
			fmt.Printf("Context finished with error: %v\n", processedCtx.Value("receiver"))
			break
		}
		inputChannel <- processedCtx.(context.Context)
	}
	close(inputChannel)
}

func Play(inputChannel <-chan context.Context, outputChannel chan<- context.Context) {
	for ctx := range inputChannel {
		newSender := ctx.Value("receiver").(string)
		newReceiver := ctx.Value("sender").(string)
		newCtx := context.WithValue(ctx, "sender", newSender)
		newCtx = context.WithValue(newCtx, "receiver", newReceiver)
		fmt.Printf("Passing to player %s\n", newReceiver)
		outputChannel <- newCtx
	}
	close(outputChannel)
}
