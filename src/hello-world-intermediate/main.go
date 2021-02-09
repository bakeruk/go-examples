/*
	An intermediate "Hello world" example, demonstrating Go's syntax and
	concurrency via its async/await equivalent, channels.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HelloWorld struct {
	hello string
	world string
}

/*
	Main
*/
func main() {
	outputLength := 10
	// Make an array with n (outputLength) indices
	results := make([]string, outputLength)
	// Example typed object
	msgObj := HelloWorld{hello: "Hello", world: "world"}
	// Example array
	strs := []string{
		fmt.Sprintf("%s %s", msgObj.hello, msgObj.world),
		msgObj.hello,
		msgObj.world,
		"hElLo",
		"WoRlD",
	}
	// Create channel
	worldChan := make(chan string)

	// Update the deterministic state of rand
	rand.Seed(time.Now().UnixNano())

	// Use a goroutine for each outputLength that
	for i := 0; i < outputLength; i++ {
		go pickRandomValue(worldChan, strs)
	}

	// Comile the results of worldChan into a simple array
	for i := range results {
		results[i] = <-worldChan

		// Output our results, line by line
		fmt.Println(results[i])
	}

	// Closes the worldChan channel
	close(worldChan)

	// Output our results
	fmt.Println(results)
}

/*
	Picks a random value from the given strs array and sets the value to the given
	channel.
*/
func pickRandomValue(worldChan chan string, strs []string) {
	randStrIndex := rand.Intn(len(strs))

	// Send the value to the channel
	worldChan <- strs[randStrIndex]
}
