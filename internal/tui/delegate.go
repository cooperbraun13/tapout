package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cooperbraun13/tapout/internal/event"
)

// EventItem wraps an event for the list component
type EventItem struct {
	Event event.Event
}

func (i EventItem) Title() string       { return i.Event.Name }
func (i EventItem) Description() string { return fmt.Sprintf("%s - %s", i.Event.Date, i.Event.Location) }
func (i EventItem) FilterValue() string { return i.Event.Name }

// FightItem wraps a fight for the list component
type FightItem struct {
	Fight  event.Fight
	Picked bool
	Winner string // Name of picked winner, empty if not picked
}

func (i FightItem) Title() string {
	return fmt.Sprintf("%s vs %s", i.Fight.FighterA.Name, i.Fight.FighterB.Name)
}
func (i FightItem) Description() string {
	if i.Picked {
		return fmt.Sprintf("Pick: %s", i.Winner)
	}
	return "No pick yet"
}
func (i FightItem) FilterValue() string { return i.Title() }

// EventDelegate is a custom delegate for rendering events
type EventDelegate struct {
	Styles Styles
}

func (d EventDelegate) Height() int                             { return 4 }
func (d EventDelegate) Spacing() int                            { return 1 }
func (d EventDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d EventDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	e, ok := item.(EventItem)
	if !ok {
		return
	}

	// Choose style based on selection
	var style lipgloss.Style
	if index == m.Index() {
		style = d.Styles.SelectedBox.Width(40)
	} else {
		style = d.Styles.Box.Width(40)
	}

	// Build the card content
	name := d.Styles.FighterName.Render(e.Event.Name)
	date := d.Styles.Muted.Render(e.Event.Date)
	location := d.Styles.Muted.Render(e.Event.Location)

	content := lipgloss.JoinVertical(lipgloss.Left, name, date, location)
	fmt.Fprint(w, style.Render(content))
}

// FightDelegate is a custom delegate for rendering fights
type FightDelegate struct {
	Styles Styles
	Picks  map[int]event.Fighter
}

func (d FightDelegate) Height() int                             { return 2 }
func (d FightDelegate) Spacing() int                            { return 0 }
func (d FightDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d FightDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	f, ok := item.(FightItem)
	if !ok {
		return
	}

	// Check if this fight has been picked
	_, picked := d.Picks[f.Fight.Order]

	// Build the fight display
	var indicator string
	if picked {
		indicator = d.Styles.Success.Render("✓")
	} else {
		indicator = d.Styles.Muted.Render("○")
	}

	fightText := fmt.Sprintf("%s vs %s", f.Fight.FighterA.Name, f.Fight.FighterB.Name)

	var line string
	if index == m.Index() {
		line = fmt.Sprintf("  %s %s", indicator, d.Styles.Selected.Render(fightText))
	} else {
		line = fmt.Sprintf("  %s %s", indicator, d.Styles.Normal.Render(fightText))
	}

	fmt.Fprint(w, line)
}

// NewEventList creates a new list configured for events
func NewEventList(events []event.Event, styles Styles, width, height int) list.Model {
	items := make([]list.Item, len(events))
	for i, e := range events {
		items[i] = EventItem{Event: e}
	}

	delegate := EventDelegate{Styles: styles}
	l := list.New(items, delegate, width, height)
	l.Title = "TAPOUT"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = styles.Title
	l.Styles.TitleBar = lipgloss.NewStyle().Padding(0, 0, 1, 2)

	return l
}

// NewFightList creates a new list configured for fights
func NewFightList(fights []event.Fight, picks map[int]event.Fighter, styles Styles, width, height int) list.Model {
	items := make([]list.Item, len(fights))
	for i, f := range fights {
		winner := ""
		picked := false
		if p, ok := picks[f.Order]; ok {
			winner = p.Name
			picked = true
		}
		items[i] = FightItem{
			Fight:  f,
			Picked: picked,
			Winner: winner,
		}
	}

	delegate := FightDelegate{Styles: styles, Picks: picks}
	l := list.New(items, delegate, width, height)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)

	return l
}
