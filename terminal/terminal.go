package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// Ask asks a question and returns the answer
// Arguments:
//
//	question: The question to ask
//	args: Optional arguments (
//		1. bool: If the answer can be empty
//		2. string: The error message to print when the answer is empty
//	)
func Ask(question string, args ...any) string {
	switch len(args) {
	case 0:
		return ask(question)
	case 1:
		// One argument provided, has to be bool.
		if args[0].(bool) {
			answer := ask(question)
			if answer == "" {
				return Ask(question, args...)
			} //else if strings.EqualFold(answer, "exit") {
			// 	os.Exit(0)
			// }
			return answer
		} else {
			q := ask(question)
			// If the answer is empty, ask again
			// If the answer is exit, exit the program
			// switch q {
			// case "exit":
			// os.Exit(0)
			// }
			return q
		}
	case 2:
		// Two arguments provided, has to be string and bool
		// If the answer is empty, ask again when the first argument is true
		// Second argument is the errormessage to print to the user.
		canBeEmpty := args[0].(bool)
		errmsg := args[1].(string)
		answer := ask(question)
		// If the answer is empty, ask again
		// If the answer is exit, exit the program
		// switch question {
		// case "exit":
		// os.Exit(0)
		// }
		if answer == "" {
			if canBeEmpty {
				return answer
			} else {
				fmt.Println(errmsg)
				return Ask(question, args...)
			}
		} else {
			return answer
		}
	}
	return ""
}

func ask(question string) string {
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
	var s = string(bytePassword)
	if s == "" {
		return AskProtected(question)
	}
	return s
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
