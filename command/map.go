package command

func init() {
	newMap := MapFactory(1)
	Register["map"] = cliCommand{
		Name:        "map",
		Description: "Go to next Map",
		Callback:    newMap,
	}
}
