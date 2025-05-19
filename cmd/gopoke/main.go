package main

import (
	"bootdev/gopoke/internal/appstate"
	"bootdev/gopoke/internal/command"
	"bootdev/gopoke/internal/logger"
	"bootdev/gopoke/internal/utils"
	"bufio"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Init()

	appstate.InitCache(ctx)

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
