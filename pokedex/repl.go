package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/api"
	"pokedex/cache"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := api.Config{}
	config.Cache = cache.NewCache()
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()

		commands := strings.Split(text, " ")

		if len(commands) == 0 {
			continue
		} else if len(commands) > 1 {

			fmt.Printf("Invalid command %s\n", text)
		} else {
			handleCommand(&config, commands[0])
		}
	}
}

func handleCommand(config *api.Config, command string) {
	commands := getCommands()
	for _, cmd := range commands {
		if cmd.name == command {
			cmd.handlerFunc(config)
			return
		}
	}
}
