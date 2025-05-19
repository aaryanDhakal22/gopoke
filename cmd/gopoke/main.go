package main

import (
	"bootdev/gopoke/internal/appstate"
	"bootdev/gopoke/internal/command"
	"bootdev/gopoke/internal/logger"
	"bootdev/gopoke/internal/utils"
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Init()
	logger.SetLevel(slog.LevelError)

	cache := appstate.NewCache(ctx)

	processor := command.NewProcessor(cache)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := utils.CleanInput(scanner.Text())
		if len(text) == 1 {
			processor.Execute(text[0])
		} else if len(text) == 2 {
			processor.ExecuteWithArgs(text[0], text[1])
		}
	}
}
