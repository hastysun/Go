// Simple math game written by Hasty in Go
// Version 1

// github.com/hastysun

/*
	It is able to generate three different types
	of mathematical problems for the user to solve.
	The three different operands are addition,
	subtraction, and multiplication. Addition and
	subtraction problems use two randomly generated
	numbers between 0 and 50 while the multiplication
	operand only generates two numbers between 0 and 10,
	to not make it too difficult.
*/

// Packages
package main

// Imports
import "fmt"
import "time"
import "math/rand"

// Variables
var AnswerType int
var RandomMathOperand string

/*
	Function that delays printing lines for fluidity,
	it is used as slp(n) whereas "n" is the amount
	of milliseconds to sleep for
*/
func slp(x time.Duration) time.Duration {
	time.Sleep(x * time.Millisecond)
	return x
}

// Function for clearing the screen
func cls() {
	fmt.Println("\033c")
}

// Intro that plays once when the program is run
func intro() {

	cls()
	fmt.Println()
	slp(500)
	fmt.Println("Welcome to MathWiz!")
	slp(1000)
	cls()
	/*
		slp(500)
		fmt.Println("Use 00 to quit")
		slp(500)
		cls()
	*/

}

/*
	game() is basically main but this had to be a separate function
	so that the ability to run the intro only once would be
	possible, a proper fix for this will be implented soon
*/
func game() {

	rand.Seed(time.Now().UnixNano())

	cls()

	LongNumberA := rand.Intn(50)
	LongNumberB := rand.Intn(50)

	ShortNumberA := rand.Intn(10)
	ShortNumberB := rand.Intn(10)

	AdditionAnswer := LongNumberA + LongNumberB
	SubtractionAnswer := LongNumberA - LongNumberB
	MultiplicationAnswer := ShortNumberA * ShortNumberB

	RandomMathOperand := rand.Intn(3)

	/*
		if RandomMathOperand == 0 {
			fmt.Println("It's 0")
			RandomMathOperand = rand.Intn(3)
		}
	*/

	if RandomMathOperand == 0 {
		AnswerType = AdditionAnswer
		fmt.Println("What is", LongNumberA, "+", LongNumberB, "?")
	}

	if RandomMathOperand == 1 {
		AnswerType = SubtractionAnswer
		fmt.Println("What is", LongNumberA, "-", LongNumberB, "?")
	}

	if RandomMathOperand == 2 {
		AnswerType = MultiplicationAnswer
		fmt.Println("What is", ShortNumberA, "*", ShortNumberB, "?")
	}

	var input int
	fmt.Scanf("%d", &input)

	if input == AnswerType {
		cls()
		fmt.Println("Correct!")
		slp(800)

	} else if input == 00 {
		cls()
		panic("Invalid Answer")

	} else {
		cls()
		fmt.Println("Wrong, answer was", AnswerType)
		slp(1000)
	}

	game() // this keeps generating new problems

}

func main() {

	defer game()
	intro()

}
