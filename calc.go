package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	number1  float64
	number2  float64
	operator string
}

func isValidNumber(input string) (float64, error) {
	if input == "" {
		return 0, fmt.Errorf("Input cannot be empty.")
	}
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Please enter a valid number.")
	}
	return num, nil
}

func isValidSymbol(input string) bool {
	return input == "+" || input == "-" || input == "*" || input == "/"
}

func getValidatedNumber(prompt string, scanner *bufio.Scanner) float64 {
	for {
		input := getInput(prompt, scanner)
		num, err := isValidNumber(input)
		if err == nil {
			return num
		}
		fmt.Println("Invalid input.", err)
	}
}

func getValidatedOperator(prompt string, scanner *bufio.Scanner) string {
	for {
		operator := getInput(prompt, scanner)
		if operator == "" {
			fmt.Println("Invalid input. Input cannot be empty.")
			continue
		}
		if isValidSymbol(operator) {
			return operator
		}
		fmt.Println("Invalid operator. Please enter a valid operator [+ - * /].")
	}
}

func getInput(prompt string, scanner *bufio.Scanner) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func getYesNoAnswer(prompt string, scanner *bufio.Scanner) bool {
	for {
		fmt.Print(prompt + " (Y/N): ")
		scanner.Scan()
		response := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(response, "Y") {
			return true
		} else if strings.EqualFold(response, "N") {
			return false
		}
		fmt.Print("Please enter Y or N.")
	}
}

func (c *Calculator) add() float64 {
	return c.number1 + c.number2
}

func (c *Calculator) minus() float64 {
	return c.number1 - c.number2
}

func (c *Calculator) multiply() float64 {
	return c.number1 * c.number2
}

func (c *Calculator) divide() (float64, error) {
	if c.number2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return c.number1 / c.number2, nil
}

func main() {
	var calc Calculator
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=====================")
		fmt.Println("Simple Calculator App")
		fmt.Println("=====================")

		calc.number1 = getValidatedNumber("Enter 1st number: ", scanner)
		calc.number2 = getValidatedNumber("Enter 2nd number: ", scanner)
		calc.operator = getValidatedOperator("Enter operator [+ - * /] : ", scanner)

		fmt.Printf("Is `%g %s %g =` the Mathematical Equation to calculate?", calc.number1, calc.operator, calc.number2)
		if !getYesNoAnswer("", scanner) {
			fmt.Println("Current operation cancelled.")
			continue
		}

		switch calc.operator {
		case "+":
			fmt.Printf("Result: %g\n", calc.add())
		case "-":
			fmt.Printf("Result: %g\n", calc.minus())
		case "*":
			fmt.Printf("Result: %g\n", calc.multiply())
		case "/":
			result, err := calc.divide()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %g\n", result)
			}
		default:
			fmt.Println("Invalid operator.")
		}

		if !getYesNoAnswer("Exit calculator app?", scanner) {
			fmt.Println("Resetting calculator...")
			calc.number1 = 0
			calc.number2 = 0
			calc.operator = ""
			continue
		}

		fmt.Println("Exiting calculator app...")
		break
	}

	fmt.Println("Calculator app closed.")
}