package picks

import "github.com/cooperbraun13/tapout/internal/event"

type Picks struct {
	EventName string
	Results   map[int]event.Fighter // Maps fight to the winning fighter
}

// Save picks to json
func (p *Picks) Save(path string) error {

}

// Function to load picks from json (for viewing past picks)
func LoadPicks(path string) (*Picks, error) {

}
