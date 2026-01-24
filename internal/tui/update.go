package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cooperbraun13/tapout/internal/event"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Global quit
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.Screen {
		case ScreenEventList:
			return m.updateEventList(msg)
		case ScreenFightList:
			return m.updateFightList(msg)
		case ScreenFightDetail:
			return m.updateFightDetail(msg)
		case ScreenSummary:
			return m.updateSummary(msg)
		}
	}

	return m, nil
}

func (m Model) updateEventList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case "up", "k":
		if m.EventCursor > 0 {
			m.EventCursor--
		}
	case "down", "j":
		if m.EventCursor < len(m.Events)-1 {
			m.EventCursor++
		}
	case "enter":
		if len(m.Events) > 0 {
			m.SelectedEvent = &m.Events[m.EventCursor]
			m.Picks.Event = m.SelectedEvent
			m.Picks.Results = make(map[int]event.Fighter) // Reset picks for new event
			m.FightCursor = 0
			m.Screen = ScreenFightList
		}
	}
	return m, nil
}

func (m Model) updateFightList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.Screen = ScreenEventList
	case "up", "k":
		if m.FightCursor > 0 {
			m.FightCursor--
		}
	case "down", "j":
		if m.SelectedEvent != nil && m.FightCursor < len(m.SelectedEvent.Fights)-1 {
			m.FightCursor++
		}
	case "enter":
		if m.SelectedEvent != nil && len(m.SelectedEvent.Fights) > 0 {
			m.PickCursor = 0 // Reset to Fighter A
			m.Screen = ScreenFightDetail
		}
	case "s":
		m.Screen = ScreenSummary
	}
	return m, nil
}

func (m Model) updateFightDetail(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q":
		m.Screen = ScreenFightList
	case "left", "h":
		m.PickCursor = 0 // Fighter A
	case "right", "l":
		m.PickCursor = 1 // Fighter B
	case "enter":
		if m.SelectedEvent != nil {
			fight := m.SelectedEvent.Fights[m.FightCursor]
			if m.PickCursor == 0 {
				m.Picks.Results[fight.Order] = fight.FighterA
			} else {
				m.Picks.Results[fight.Order] = fight.FighterB
			}
			m.Screen = ScreenFightList
		}
	}
	return m, nil
}

func (m Model) updateSummary(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q":
		m.Screen = ScreenFightList
	case "enter":
		// Save picks and quit
		m.Picks.Save()
		return m, tea.Quit
	}
	return m, nil
}
