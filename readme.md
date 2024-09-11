# Powerball Auto Generator

This project is a Go-based application that automatically generates and records Powerball lottery numbers. It's designed to run on a schedule, generating random numbers and committing them to a JSON file in a Git repository.

## Features

- Generates random Powerball lottery numbers (5 numbers from 1-69 and 1 Powerball number from 1-26)
- Schedules number generation four times daily: 9 AM, 1 PM, 5 PM, and 9 PM Pacific Time
- Stores generated numbers in a JSON file
- Automatically commits and pushes changes to a Git repository
- Supports manual testing and immediate number generation

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/JZilla808/powerball-auto-generator.git
   ```

2. Navigate to the project directory:
   ```
   cd powerball-auto-generator
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

To run the application:

```
go run main.go
```

This will start the scheduled number generator, which will run four times daily at 9 AM, 1 PM, 5 PM, and 9 PM Pacific Time.

## Commands and Functions

The project supports the following main functions:

1. `StartRandomNumberGenerator()`: Starts the scheduled random number generator.
   
```go
// StartRandomNumberGenerator initializes and starts the lottery number generator
func StartRandomNumberGenerator() {
	c := cron.New()
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Define multiple run times throughout the day
	runTimes := []string{
		"CRON_TZ=America/Los_Angeles 0 9 * * *",  // 9 AM
		"CRON_TZ=America/Los_Angeles 0 13 * * *", // 1 PM
		"CRON_TZ=America/Los_Angeles 0 17 * * *", // 5 PM
		"CRON_TZ=America/Los_Angeles 0 21 * * *", // 9 PM
	}

	for _, timeToRun := range runTimes {
		fmt.Printf("Job is scheduled to run at %s every day in America/Los_Angeles timezone.\n", timeToRun)
		c.AddFunc(timeToRun, func() {
			generateAndCommit(location)
		})
	}

	c.Start()
}
```

2. `StartRandomNumberGeneratorNow()`: Generates numbers immediately without waiting for the schedule.
   
```go
func StartRandomNumberGeneratorNow() {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	generateAndCommit(location)
}
```

3. `TestGenerateAndCommit()`: A test function to manually generate one set of numbers and commit them.
   
```go
// testGenerateAndCommit is a manual test function for generateAndCommit
func TestGenerateAndCommit() {
	fmt.Println("TestGenerateAndCommit Called...")
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Manually set the random count to 1 for testing
	originalGenerateCount := generateCount
	defer func() { generateCount = originalGenerateCount }()
	generateCount = func() int { return 1 }

	// Call generateAndCommit for testing
	generateAndCommit(location)
}
```

To use these functions, uncomment the relevant lines in the `main()` function:

```go
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
```

## Configuration

- The schedule is set to run four times daily at 9 AM, 1 PM, 5 PM, and 9 PM Pacific Time. You can modify this in the `StartRandomNumberGenerator()` function.
- The number of commits per run is randomly generated between 5 and 18. This can be adjusted in the `generateCount()` function.

## File Structure

- `main.go`: Entry point of the application
- `internal/random/random.go`: Contains the core logic for number generation and Git operations
- `internal/random/numbers.json`: Stores the generated lottery numbers

## Dependencies

- github.com/robfig/cron/v3: Used for scheduling the number generation

## Note

Ensure that you have proper Git credentials set up on your machine for the automatic commit and push functionality to work.