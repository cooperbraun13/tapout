package tui

import "github.com/cooperbraun13/tapout/internal/event"

type Screen int

const (
	ScreenEventList  Screen = iota // 0
	ScreenFightPicks               // 1
	ScreenSummary                  // 2
)

type Model struct {
	Screen       Screen            // Which screen we are currently on
	Events       []event.Event     // Events listed
	CurrentEvent int               // Cursor on what event we are selecting
	CurrentFight int               // Cursor on what fight we are selecting
	Picks        map[string]string // Fight -> Winner

}
