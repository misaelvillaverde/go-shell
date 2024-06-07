package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		parseCommand(command)
	}
}

var (
	builtins = map[string]struct{}{
		"exit": {},
		"echo": {},
		"type": {},
	}
)

func typeof(cmd string) string {
	if _, ok := builtins[cmd]; ok {
		return cmd + " is a shell builtin"
	}

	if path, err := exec.LookPath(cmd); err == nil {
		return cmd + " is " + path
	}

	return cmd + ": not found"
}

func parseCommand(command string) {
	predicate := strings.Split(command[:len(command)-1], " ")
	cmd := predicate[0]
	args := predicate[1:]

	switch cmd {
	case "exit":
		os.Exit(0)
		return
	case "echo":
		fmt.Println(strings.Join(args, " "))
		return
	case "type":
		fmt.Println(typeof(args[0]))
		return
	case "cd":
		if err := os.Chdir(args[0]); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", args[0])
		}
		return
	}

	if output, err := exec.Command(cmd, args...).Output(); err == nil {
		fmt.Print(string(output))
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}
}
