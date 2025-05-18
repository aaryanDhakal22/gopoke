package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Counter int = 0

func MapGen() {
	locs := LocType{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 20*(Counter-1))
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to create a request", err)
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locs); err != nil {
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
