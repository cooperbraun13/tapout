ufc-picker/
|- main.go           # Entry point, starts TUI
|- go.mod
|- go.sum
|- events/           # TOML config files (manually updated)
|  |- ufc324.toml
|
|- picks/            # Saved picks (program-written, JSON or TOML)
|  |- ufc324.json    
|
|- internal/
   |- event/
   |  |- event.go    # Event, Fight, Fighter structs + loading
   |
   |- picks/
   |  |- picks.go    # Load/save user picks
   |
   |- tui/
      |- model.go    # Bubble Tea model
      |- update.go   # Handle keypresses
      |- view.go     # Render to terminal
