package event

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
)

type Event struct {
	Name     string  `toml:"name"`
	Date     string  `toml:"date"`
	Location string  `toml:"location"`
	Venue    string  `toml:"venue"`
	Fights   []Fight `toml:"fights"`
}

type Fight struct {
	Order    int     `toml:"order"`
	FighterA Fighter `toml:"fighter_a"`
	FighterB Fighter `toml:"fighter_b"`
}

type Fighter struct {
	Name     string `toml:"name"`
	Record   string `toml:"record"`
	Ranking  string `toml:"ranking"`
	LastFive string `toml:"lastFive"`
	Division string `toml:"division"`
	Type     string `toml:"type"`
	Odds     int    `toml:"odds"`
}

func hello() {
	// load from TOML
	var event Event
	toml.DecodeFile("events/ufc324.toml", &event)

	// save as JSON
	file, _ := os.Create("picks/ufc324.json")
	json.NewEncoder(file).Encode(event)
}
