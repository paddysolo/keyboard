// Package keyboard reads user input from a keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetUserInput function reads a floating-point number from a keyboard
// It returns the number read and any error encountered
func GetUserInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	userInput = strings.TrimSpace(userInput)
	number, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
