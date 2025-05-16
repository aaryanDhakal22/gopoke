package command

func init() {
	newMap := MapFactory(-1)
	Register["mapb"] = cliCommand{
		Name:        "mapb",
		Description: "Go to previous Map",
		Callback:    newMap,
	}
}
