package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	assertArgsValid()
	originUnit = strings.ToUpper(os.Args[1])
	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err := fmt.Scanln(&originValue)
		assertInputValid(err)
		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		shouldConvertAgain = strings.ToUpper(shouldConvertAgain)
		assertInputValid(err)

		if shouldConvertAgain != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func assertInputValid(err error) {
	if err != nil {
		printError(errReadingInput)
	}
}

func assertArgsValid() {
	commandLineArguments := os.Args
	if len(commandLineArguments) != 2 {
		printError(errInvalidArguments)
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
