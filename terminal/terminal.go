package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Ask(question string) string {
	var input string
	fmt.Print(question)
	// Scan until enter is pressed
	std := bufio.NewScanner(os.Stdin)
	std.Scan()
	input = std.Text()
	return input
}

func AskProtected(question string) string {
	// Hide input
	fmt.Print(question)
	bytePassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(bytePassword)
}

func RepeatAsk(question string, validAnswers []string, cancel bool) (string, error) {
	var input string
	if cancel {
		question += " (or type 'cancel' to cancel)"
	}
	validAnswers = append(validAnswers, "cancel")
	for {
		input = Ask(question)
		for _, v := range validAnswers {
			if strings.EqualFold(input, v) {
				if strings.EqualFold(input, "cancel") {
					return "", fmt.Errorf("cancelled by user")
				}
				return input, nil
			}
		}
	}
}
