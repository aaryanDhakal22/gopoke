package command

import (
	"bootdev/gopoke/internal/pokecache"
	"fmt"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func()
}
type Processor struct {
	commands   map[string]cliCommand
	cache      *pokecache.PokeCache
	mapCounter int
}

func NewProcessor(cache *pokecache.PokeCache) *Processor {
	p := &Processor{
		commands:   make(map[string]cliCommand),
		cache:      cache,
		mapCounter: 0,
	}
	p.registerCommands()
	return p
}

func (p *Processor) Execute(cmdName string) {
	if cmd, ok := p.commands[cmdName]; ok {
		cmd.Callback()
	} else {
		fmt.Println("Unknown Command")
	}
}
func (p *Processor) registerCommands() {
	p.commands["exit"] = cliCommand{
		Name:        "exit",
		Description: "Exit the program",
		Callback:    p.commandExit,
	}
	p.commands["help"] = cliCommand{
		Name:        "Help",
		Description: "Print the Help message",
		Callback:    p.commandHelp,
	}
	p.commands["mapb"] = cliCommand{
		Name:        "mapb",
		Description: "Go to previous Map",
		Callback:    p.mapFactory(-1),
	}
	p.commands["map"] = cliCommand{
		Name:        "map",
		Description: "Go to next Map",
		Callback:    p.mapFactory(1),
	}

}

func (p *Processor) commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for k, v := range p.commands {
		fmt.Printf("%v: %v\n", k, v.Description)
	}
	return
}

func (p *Processor) commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return
}
