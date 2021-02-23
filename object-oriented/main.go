// generates a cat object using OOP equivalent methodology
package main

import (
	"fmt"
)

func main() {
	// Creates a new Cat
	Tom := *New("Tom", "Jangles")
	Tim := *New("Joe", "Black")

	// Take one life away from Tom
	Tom.TakeOneLife()
	// Take 2 lives away from Tom
	Tom.TakeLives(2)

	// Output our Cats
	fmt.Printf("%s %s has %d lives left and %s %s has %d lives left\n", Tom.FirstName, Tom.LastName, Tom.LivesLeft, Tim.FirstName, Tim.LastName, Tim.LivesLeft)
}
