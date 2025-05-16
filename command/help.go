package command

import (
	"fmt"
)

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for k, v := range Register {
		fmt.Printf("%v: %v\n", k, v.Description)
	}
	return
}
func init() {

	Register["help"] = cliCommand{
		Name:        "Help",
		Description: "Print the Help message",
		Callback:    commandHelp,
	}
}
