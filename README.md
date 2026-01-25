# Tapout

A terminal UI for making UFC fight picks. Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Usage

```bash
go run .
```

Navigate with arrow keys, press Enter to select, and make your picks for upcoming events.

## Adding Events

Create a TOML file in the `events/` directory:

```toml
name = "UFC 324"
date = "2026-01-24"
location = "Las Vegas, Nevada"
venue = "T-Mobile Arena"

[[fights]]
order = 1

[fights.fighter_a]
name = "Fighter One"
record = "20-5-0"
ranking = "C"
lastFive = "W W W W L"
division = "Lightweight"
odds = -150

[fights.fighter_b]
name = "Fighter Two"
record = "18-3-0"
ranking = "3"
lastFive = "W W W W W"
division = "Lightweight"
odds = +130
```

## Project Structure

```
tapout/
├── main.go                     # Entry point
├── go.mod
├── go.sum
├── events/                     # Event TOML files (manually updated)
│   ├── ufc324.toml
│   └── ufc325.toml
├── picks/                      # Saved picks (auto-generated JSON)
│   └── ufc324.json
├── internal/
│   ├── event/
│   │   └── event.go            # Event, Fight, Fighter structs + loading
│   ├── picks/
│   │   └── picks.go            # Load/save user picks
│   └── tui/
│       ├── model.go            # Bubble Tea model
│       ├── update.go           # Handle keypresses
│       ├── view.go             # Render screens
│       ├── styles.go           # Lipgloss styling
│       └── delegate.go         # List delegates
└── .github/
    └── workflows/
        └── ci.yml              # GitHub Actions CI
```

## Maintenance

Events and fights will be updated up to the day of by myself for the foreseeable future. Any mistakes or suggestions should be written in a pull request.
