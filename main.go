package main

import (
	"bootdev/gopoke/command"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	ftd := strings.Trim(text, " ")
	ftd = strings.Trim(ftd, "\n")
	ftd = strings.ToLower(ftd)
	ret := strings.Split(ftd, " ")

	return ret

}
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		obj, ok := command.Register[text[0]]
		if ok {
			obj.Callback()
		} else {
			fmt.Println("Unknown Command")
		}
	}
}
