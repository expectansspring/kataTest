package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabicNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var romanNumbers = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var signs = []string{"+", "-", "*", "/"}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n') // обработать ошибку
	text = strings.TrimSpace(text)
	return text
}

func fromRomanToArabic(str string) int {
	var result int
	for i := 0; i < len(romanNumbers); i++ {
		if str == romanNumbers[i] {
			result = i + 1
			break
		}
	}
	return result
}

func doesContain(inputArray []string, str string) bool {
	var result = false
	for i := 0; i < len(inputArray); i++ {
		if str == inputArray[i] {
			result = true
			break
		}
	}
	return result
}

func isCorrectInput(inputArray []string) {
	var result = 0

	for i := 0; i < len(inputArray); i++ {
		if i%2 == 1 && !doesContain(signs, inputArray[i]) {
			result = 1
		} else if i%2 == 0 && !(doesContain(arabicNumbers, inputArray[i]) || doesContain(romanNumbers, inputArray[i])) {
			result = 1
		}
	}
	if result == 1 {
		panic("The string is not a mathematical operation")
	}

	if len(inputArray) < 3 {
		result = 1
	} else if len(inputArray) > 3 {
		result = 2
	} else if len(inputArray) == 3 && !(doesContain(arabicNumbers, inputArray[0]) && doesContain(arabicNumbers, inputArray[2]) || doesContain(romanNumbers, inputArray[0]) && doesContain(romanNumbers, inputArray[2])) {
		result = 3
	}

	switch result {
	case 1:
		panic("The string is not a mathematical operation")
	case 2:
		panic("The format of the mathematical operation does not satisfy the task — two operands and one operator (+, -, /, *).")
	case 3:
		panic("Different number systems are used simultaneously.")
	}
}

func isRoman(str string) bool {
	var result = true
	if doesContain(arabicNumbers, str) {
		result = false
	}
	return result
}

func fromArabicToRoman(number int) string {
	var result = ""
	var romanDigits = map[string]int{
		"C": 100, "L": 50, "X": 10, "V": 5, "I": 1,
	}
	for key, value := range romanDigits {
		if number == 0 {
			break
		} else if number < 11 {
			result += romanNumbers[number-1]
			break
		} else if number >= value-10 && number < value {
			result += "X" + key
			number -= value - 10
		} else {
			for i := 0; i < number/value; i++ {
				result += key
			}
			number %= value
		}
	}
	return result
}

func calculatingOperation(a, b int, sign string) int {
	var result = 0
	switch sign {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	return result
}

func calculatingResult(inputArray []string, roman bool) int {
	var result = 0
	if roman == false {
		a, _ := strconv.Atoi(inputArray[0])
		b, _ := strconv.Atoi(inputArray[2])
		result = calculatingOperation(a, b, inputArray[1])
	} else {
		var a, b = fromRomanToArabic(inputArray[0]), fromRomanToArabic(inputArray[2])
		result = calculatingOperation(a, b, inputArray[1])
		if result < 1 {
			panic("There are no negative numbers in the Roman system.")
		}
	}
	return result
}

func output(value int, currentType bool) {
	if currentType == false {
		fmt.Println(value)
	} else {
		fmt.Println(fromArabicToRoman(value))
	}
}

func main() {
	var inputString = input()
	var inputArray = strings.Split(inputString, " ")
	isCorrectInput(inputArray)
	var currentType = isRoman(inputArray[0])
	output(calculatingResult(inputArray, currentType), currentType)
}
