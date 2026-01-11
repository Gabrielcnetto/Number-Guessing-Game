package services

import (
	"fmt"
	"math/rand"
)

var diff = map[int]string{1: "Easy", 2: "Medium", 3: "Hard"}
var difContent = map[int]string{
	1: "10 chances",
	2: "5 chances",
	3: "3 chances",
}
var difAttempts = map[int]int{
	1: 10,
	2: 5,
	3: 3,
}

func Game() error {
	var difficulty int
	var responseValue int

	thinkedValue := rand.Intn(100) + 1
	attempts := 0

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("\nEnter your choice: ")

	fmt.Scanln(&difficulty)
	maxAttempts, ok := difAttempts[difficulty]
	if !ok {
		return fmt.Errorf("invalid difficulty")
	}
	fmt.Printf("\nGreat! You selected %s mode.\n", diff[difficulty])
	fmt.Println("Let's start the game!")
	for attempts < maxAttempts {
		fmt.Printf("\nAttempt %d/%d - Enter your guess: ", attempts+1, maxAttempts)
		fmt.Scanln(&responseValue)

		attempts++

		switch {
		case responseValue < thinkedValue:
			fmt.Println("Incorrect! The number is greater.")
		case responseValue > thinkedValue:
			fmt.Println("Incorrect! The number is smaller.")
		default:
			fmt.Printf("\n🎉 Congratulations! You guessed the number in %d attempts.\n", attempts)
			return nil
		}
	}
	fmt.Printf("\n❌ Game over! The correct number was %d\n", thinkedValue)
	return nil
}
