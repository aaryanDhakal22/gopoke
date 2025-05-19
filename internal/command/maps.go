package command

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocType struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (p *Processor) mapGen() {
	locs := LocType{}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 20*(p.mapCounter-1))
	data, ok := p.cache.Get(url)
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
		p.cache.Add(url, data)
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
func (p *Processor) mapFactory(change int) func() {
	if change != 1 && change != -1 {
		panic("The world is upside down. It should be -1 or 1")
	}
	return func() {
		p.mapCounter += change
		if p.mapCounter <= 0 {
			fmt.Println("Already on first page")
			p.mapCounter = 0
			return
		}
		p.mapGen()
	}

}
