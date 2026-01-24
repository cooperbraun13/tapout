package tui

import (
	"github.com/cooperbraun13/tapout/internal/event"
	"github.com/cooperbraun13/tapout/internal/picks"

	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	ScreenEventList   Screen = iota // List of events to choose from
	ScreenFightList                 // List of fights for selected event
	ScreenFightDetail               // Viewing fighter info and making a pick
	ScreenSummary                   // Review all picks before saving
)

type Model struct {
	// Navigation
	Screen Screen // Current screen

	// Event list state
	Events      []event.Event // All loaded events
	EventCursor int           // Cursor position on event list

	// Fight list state
	SelectedEvent *event.Event // The event user selected
	FightCursor   int          // Cursor position on fight list

	// Fight detail state
	PickCursor int // 0 = Fighter A, 1 = Fighter B

	// Picks state
	Picks picks.Picks // User's picks (maps fight order -> chosen fighter)
}

func New() Model {
	events, _ := event.LoadAll()

	return Model{
		Screen:      ScreenEventList,
		Events:      events,
		EventCursor: 0,
		FightCursor: 0,
		Picks: picks.Picks{
			Results: make(map[int]event.Fighter),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
