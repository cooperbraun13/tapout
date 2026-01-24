package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cooperbraun13/tapout/internal/event"
	"github.com/cooperbraun13/tapout/internal/picks"
)

type Screen int

const (
	ScreenEventList   Screen = iota // List of events to choose from
	ScreenFightList                 // List of fights for selected event
	ScreenFightDetail               // Viewing fighter info and making a pick
	ScreenSummary                   // Review all picks before saving
)

// KeyMap defines the keybindings for the app
type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Left    key.Binding
	Right   key.Binding
	Enter   key.Binding
	Back    key.Binding
	Summary key.Binding
	Quit    key.Binding
}

// DefaultKeyMap returns the default keybindings
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc", "backspace"),
			key.WithHelp("esc", "back"),
		),
		Summary: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "summary"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

// ShortHelp returns keybindings shown in the mini help view
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Back, k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Enter, k.Back, k.Summary, k.Quit},
	}
}

type Model struct {
	// Navigation
	Screen Screen // Current screen

	// Event list state
	Events    []event.Event // All loaded events
	EventList list.Model    // Bubbles list for events

	// Fight list state
	SelectedEvent *event.Event // The event user selected
	FightList     list.Model   // Bubbles list for fights
	FightCursor   int          // Keep for fight detail navigation

	// Fight detail state
	PickCursor int // 0 = Fighter A, 1 = Fighter B

	// Picks state
	Picks picks.Picks // User's picks (maps fight order -> chosen fighter)

	// UI components
	Styles Styles
	Keys   KeyMap
	Help   help.Model

	// Dimensions
	Width  int
	Height int
}

func New() Model {
	events, _ := event.LoadAll()
	styles := DefaultStyles()
	keys := DefaultKeyMap()

	// Default dimensions (will be updated on WindowSizeMsg)
	width := 80
	height := 24

	// Create event list
	eventList := NewEventList(events, styles, width, height-4)

	return Model{
		Screen:    ScreenEventList,
		Events:    events,
		EventList: eventList,
		Picks: picks.Picks{
			Results: make(map[int]event.Fighter),
		},
		Styles: styles,
		Keys:   keys,
		Help:   help.New(),
		Width:  width,
		Height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
