package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		command = strings.TrimSpace(command)

		predicate := strings.Split(command, " ")

		if parse, ok := commands[predicate[0]]; ok {
			parse(predicate[1:]...)
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}

func typeof(cmd string) string {
	if _, ok := commands[cmd]; ok {
		return cmd + " is a shell builtin"
	}

	return cmd + " not found"
}

var commands map[string]func(args ...string)

func init() {
	commands = map[string]func(args ...string){
		"exit": func(args ...string) {
			os.Exit(0)
		},
		"echo": func(args ...string) {
			fmt.Println(strings.Join(args, " "))
		},
		"type": func(args ...string) {
			fmt.Println(typeof(args[0]))
		},
	}
}
