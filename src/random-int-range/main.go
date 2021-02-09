/*
	A truly random integer range. Any number generated through math/rand is not
	really random by default, as being deterministic it will always print the
	same value each time.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Main
*/
func main() {
	// Update the deterministic state of rand
	rand.Seed(time.Now().UnixNano())
	// Generate a random integer with the given range
	randomInt := randomRangedInt(1, 10)

	// Output our results
	fmt.Println(randomInt)
}

/*
	Picks a random integer between the given min/max range
*/
func randomRangedInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
