package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	call        func() error
}

// fmt.Scanln(&poop)
// stops after whitespace. not useful

func caller(cmdList map[string]command, cmd string) {
	// should include error return
	cmdToRun := cmdList[cmd]
	cmdToRun.call()
}

func repl() {
	prompt := "Pokedex CLI >> "
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)

	for scanner.Scan() {
		cmdWritten, err := cleanInput(scanner.Text())
		if err != nil {
			fmt.Print(prompt)
			continue
		}

		if cmd, ok := getCommand()[cmdWritten[0]]; ok {
			cmd.call()
		} else {
			fmt.Printf("The command you have entered: <%v> is invalid.\n", cmdWritten)
		}

		fmt.Print(prompt)
	}
}

func getCommand() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Prints a help message",
			call:        commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the REPL",
			call:        commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			call:        mapn,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			call:        mapb,
		},
	}
}

func cleanInput(str string) ([]string, error) {
	output := strings.TrimSpace(str)

	// if after removing whitespace, string is empty, return err
	if output == "" {
		return []string{}, errors.New("String only containd whitespace.")
	}

	// splits string into a slice of strings. whitespace
	return strings.Fields(strings.ToLower(output)), nil
}
