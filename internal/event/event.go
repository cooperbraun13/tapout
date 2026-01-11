package event

type Event struct {
	Name   string  `toml:"name"`
	Date   string  `toml:"date"`
	Fights []Fight `toml:"fights"`
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
