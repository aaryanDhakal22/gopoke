package command

type cliCommand struct {
	Name        string
	Description string
	Callback    func()
}
type LocType struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

var Register = map[string]cliCommand{}
