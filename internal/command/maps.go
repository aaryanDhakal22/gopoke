package command

import (
	"bootdev/gopoke/internal/appstate"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var Counter int = 0

func MapGen() {
	locs := LocType{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 20*(Counter-1))
	data, ok := appstate.GlobalCache.Get(url)
	if !ok {
		log.Println("\n\n Getting a new copy ")
		res, err := http.Get(url)
		if err != nil {
			log.Fatal("Unable to create a request", err)
			return
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Unable to convert to bytes", err)
		}
		appstate.GlobalCache.Add(url, data)
		defer res.Body.Close()
	} else {
		log.Println("Found a copy on cache")
	}
	if err := json.Unmarshal(data, &locs); err != nil {
		log.Fatal("Unable to decode the json", err)
		return
	}
	for idx := range locs.Results {
		fmt.Println(locs.Results[idx].Name)
	}
	return
}
func MapFactory(change int) func() {
	if change != 1 && change != -1 {
		panic("The world is upside down. It should be -1 or 1")
	}
	return func() {
		rem := Counter + change
		if rem <= 0 {
			fmt.Println("Already on first page")
			return
		}
		Counter = rem
		MapGen()
	}

}
