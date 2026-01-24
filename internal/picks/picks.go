package picks

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cooperbraun13/tapout/internal/event"
)

type Picks struct {
	Event   *event.Event
	Results map[int]event.Fighter // Maps fight order to the winning fighter
}

// Save picks to json with pretty formatting
func (p *Picks) Save() error {
	// Directory where we create our result files
	dir := "picks"

	// Ensure the picks directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Convert our struct into a pretty-formatted JSON byte slice
	data, err := json.MarshalIndent(p, "", "  ")
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

// LoadPicks loads picks from json for a given event slug
// Returns nil with no error if the file doesn't exist (new event)
func LoadPicks(slug string) (*Picks, error) {
	dir := "picks"
	filename := slug + ".json"
	fullPath := filepath.Join(dir, filename)

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, nil // No picks file yet, not an error
	}

	// Read the file
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal into Picks struct
	var picks Picks
	if err := json.Unmarshal(data, &picks); err != nil {
		return nil, err
	}

	// Ensure Results map is initialized even if empty in JSON
	if picks.Results == nil {
		picks.Results = make(map[int]event.Fighter)
	}

	return &picks, nil
}
