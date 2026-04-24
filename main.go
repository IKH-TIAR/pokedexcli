package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct{
	Name string
	Description string
	callback func() error
}

var commands map[string]Command

func main() {
	commands = map[string]Command{
		"exit": {
			Name: "exit",
			Description: "Exit the Pokedex CLI",
			callback: commandExit,
		},
		"help": {
			Name: "help",
			Description: "Show this help message",
			callback: commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := strings.ToLower(strings.TrimSpace(text))

		words := strings.Fields(cleaned)

		if len(words) == 0 {
			continue
		}
		cmdName := words[0]
		
		cmd, exists := commands[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error executing command '%s': %v\n", cmdName, err)
		}
	}
}

func commandExit() error {
	fmt.Println("Exiting Pokedex CLI. Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to Pokerdex!")
	fmt.Println("Usage:")

	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}