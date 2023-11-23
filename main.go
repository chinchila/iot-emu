package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chinchila/iot-emu/internal/cmd"
)

func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func parse(fullCommand string) bool {
	args := strings.Split(fullCommand, " ")
	command := args[0]
	v, ok := cmd.COMMANDS_FUNCTION_MAP[command]
	if !ok {
		fmt.Fprintf(os.Stderr, "This command does not exist.\n")
		return false
	}
	if v.(func([]string) int)(args[1:]) == 1 {
		fmt.Fprintf(os.Stderr, "\nThere was an error running the last command, here is the usage:\n%s\n", cmd.COMMANDS_USAGE_MAP[command])
	}
	return true
}

func main() {
	for {
		command := prompt(">")
		parse(command)
	}
}
