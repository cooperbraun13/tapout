package event

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Event struct {
	Slug     string  // "ufc324" - derived from filename as an identifier for matching output files
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

func LoadAll() ([]Event, error) {
	// Slice of Events where we load our .toml files into
	var events []Event
	// Directory where we store the event data
	dir := "events"

	// Get the list of files from a specified directory. In our case, we will be grabbing the .toml files from
	// the /events directory primarily
	files, err := os.ReadDir(dir) // Returns a a slice of os.DirEntry objects (not file paths)
	if err != nil {
		log.Fatalf("Failed to get the list of files from given directory: %v", err)
	}

	for _, file := range files {
		// Filter only for .toml files right now (because thats the only type of file we expect)
		if filepath.Ext(file.Name()) != ".toml" {
			continue
		}

		// Create fresh event for each file (prevents fights bleeding between events)
		var event Event

		// Extract slug to use as an identifier: "ufc324.toml" -> "ufc324"
		// Can use this so our picks file saves with the same name
		slug := strings.TrimSuffix(file.Name(), ".toml")
		event.Slug = slug

		// toml.DecodeFile expects a string (for the path) so we need to combine the directory we gave
		// and call the .Name method on the file (which is a os.DirEntry object) and combine them
		// to get the full correct path
		// Ex: dir -> "events", file.Name() -> "ufc324.toml", final -> "events/ufc324.toml"
		fullPath := filepath.Join(dir, file.Name())
		if _, err := toml.DecodeFile(fullPath, &event); err != nil {
			log.Fatalf("Failed to decode TOML file: %v", err)
		}

		// Append populated event to the final slice
		events = append(events, event)
	}

	return events, nil
}
