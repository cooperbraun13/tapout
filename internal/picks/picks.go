package picks

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cooperbraun13/tapout/internal/event"
)

type Picks struct {
	Event   *event.Event
	Results map[int]event.Fighter // Maps fight to the winning fighter
}

// Save picks to json
func (p *Picks) Save() error {
	// Directory where we create our result files
	dir := "picks"
	// Convert our struct into a JSON-formatted byte slice
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// Create the file at the specified full path, i.e., picks/ufc324.json
	filename := p.Event.Slug + ".json"
	fullPath := filepath.Join(dir, filename)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	// Make sure file closes when we're done
	defer file.Close()

	// Write the JSON data to the newly created file
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Function to load picks from json (for viewing past picks)
func LoadPicks(path string) (*Picks, error) {
	return nil, nil
}
