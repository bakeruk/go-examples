package main

// Cat defines the structure of a cat
type Cat struct {
	FirstName string
	LastName  string
	LivesLeft int
}

// New generates a new cat
func New(firstName string, lastName string) *Cat {
	return &Cat{firstName, lastName, 9}
}

// TakeOneLife removes 1 life from the cat
func (c *Cat) TakeOneLife() {
	c.LivesLeft--
}

// TakeLives removes n live from the cat
func (c *Cat) TakeLives(n int) {
	c.LivesLeft -= n
}
