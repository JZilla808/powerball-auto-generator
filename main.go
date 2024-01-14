package main

import (
	"fmt"

	"github.com/JZilla808/powerball-auto-generator/internal/random"
)

func main() {
	// Test the random number generator
	// fmt.Println("Testing the random number generator...")
	// go random.TestGenerateAndCommit()

	// Start the random number generator as a goroutine
	fmt.Println("Starting the random number generator...")
	go random.StartRandomNumberGenerator()

}
