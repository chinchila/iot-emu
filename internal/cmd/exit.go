package cmd

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EXIT_LITERAL = "exit"
	EXIT_USAGE   = EXIT_LITERAL + " [code]"
)

func Exit(args []string) int {
	code := 0
	if len(args) >= 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("You need to specify an integer exit code")
		}
		code = i
	}
	os.Exit(code)
	return 0
}
