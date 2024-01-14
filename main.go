package main

import (
	"fmt"
	"time"

	"github.com/JZilla808/powerball-auto-generator/internal/random"
)

func main() {
	// Test the random number generator right now
	// fmt.Println("Testing the random number generator now...")
	// random.StartRandomNumberGeneratorNow()

	// Test the random number generator to make one commit
	// fmt.Println("Testing the random number generator to make one commit...")
	// random.TestGenerateAndCommit()

	// Start the scheduled random number generator
	fmt.Println("Starting the random number generator...")
	random.StartRandomNumberGenerator()

	// Keep the program running
	for {
		time.Sleep(time.Hour)
	}

}
