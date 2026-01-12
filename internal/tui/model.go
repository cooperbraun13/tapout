package tui

import (
	"github.com/cooperbraun13/tapout/internal/event"
	"github.com/cooperbraun13/tapout/internal/picks"
)

type Screen int

const (
	ScreenEventList  Screen = iota // 0
	ScreenFightPicks               // 1
	ScreenSummary                  // 2
)

type Model struct {
	Screen       Screen        // Which screen we are currently on
	Events       []event.Event // Events listed
	Picks        picks.Picks   // User's selection
	CurrentEvent int           // Cursor on what event we are selecting
}
