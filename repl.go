package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/zspekt/pokedexcli/internal/pokeapi"
)

var CmdWritten []string = []string{}

type command struct {
	name        string
	description string
	call        func(*pokeapi.Config) error
}

// fmt.Scanln(&poop)
// stops after whitespace. not useful

func caller(cmdList map[string]command, cmd string) {
	// should include error return
	cmdToRun := cmdList[cmd]
	cmdToRun.call(Cfg)
}

func repl() {
	prompt := "Pokedex CLI >> "
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)

	for scanner.Scan() {
		var err error
		CmdWritten, err = cleanInput(scanner.Text())
		if err != nil {
			fmt.Print(prompt)
			continue
		}

		if cmd, ok := getCommand()[CmdWritten[0]]; ok {
			cmd.call(Cfg)
		} else {
			fmt.Printf("The command you have entered: <%v> is invalid.\n", CmdWritten)
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
		"explore": {
			name:        "explore",
			description: "Explore the selected area",
			call:        explore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			call:        catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon",
			call:        inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display your Pokedex",
			call:        pokedex,
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
