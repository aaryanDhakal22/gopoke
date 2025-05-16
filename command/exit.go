package command

import (
	"fmt"
	"os"
)

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return
}

func init() {
	Register["exit"] = cliCommand{
		Name:        "exit",
		Description: "Exit the program",
		Callback:    commandExit,
	}
}
