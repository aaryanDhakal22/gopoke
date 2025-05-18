package main

import (
	"bootdev/gopoke/internal/command"
	"bootdev/gopoke/internal/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := utils.CleanInput(scanner.Text())
		obj, ok := command.Register[text[0]]
		if ok {
			obj.Callback()
		} else {
			fmt.Println("Unknown Command")
		}
	}
}
