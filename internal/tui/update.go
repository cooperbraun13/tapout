package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cooperbraun13/tapout/internal/event"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.EventList.SetSize(msg.Width, msg.Height-4)
		// Only resize FightList if an event has been selected
		if m.SelectedEvent != nil {
			m.FightList.SetSize(msg.Width, msg.Height-8)
		}
		return m, nil

	case tea.KeyMsg:
		// Global quit
		if key.Matches(msg, m.Keys.Quit) && m.Screen == ScreenEventList {
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

	return m, cmd
}

func (m Model) updateEventList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.Keys.Enter):
		// Select current event
		if item, ok := m.EventList.SelectedItem().(EventItem); ok {
			m.SelectedEvent = &item.Event
			m.Picks.Event = m.SelectedEvent
			m.Picks.Results = make(map[int]event.Fighter)
			m.FightCursor = 0

			// Create fight list for selected event
			m.FightList = NewFightList(
				m.SelectedEvent.Fights,
				m.Picks.Results,
				m.Styles,
				m.Width,
				m.Height-8,
			)
			m.Screen = ScreenFightList
		}
		return m, nil

	default:
		// Delegate to list component for navigation
		m.EventList, cmd = m.EventList.Update(msg)
		return m, cmd
	}
}

func (m Model) updateFightList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.Keys.Back):
		m.Screen = ScreenEventList
		return m, nil

	case key.Matches(msg, m.Keys.Summary):
		m.Screen = ScreenSummary
		return m, nil

	case key.Matches(msg, m.Keys.Enter):
		// Select current fight for detail view
		m.FightCursor = m.FightList.Index()
		m.PickCursor = 0
		m.Screen = ScreenFightDetail
		return m, nil

	default:
		// Delegate to list component for navigation
		m.FightList, cmd = m.FightList.Update(msg)
		return m, cmd
	}
}

func (m Model) updateFightDetail(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.Keys.Back):
		m.Screen = ScreenFightList
		return m, nil

	case key.Matches(msg, m.Keys.Left):
		m.PickCursor = 0
		return m, nil

	case key.Matches(msg, m.Keys.Right):
		m.PickCursor = 1
		return m, nil

	case key.Matches(msg, m.Keys.Enter):
		if m.SelectedEvent != nil && m.FightCursor < len(m.SelectedEvent.Fights) {
			fight := m.SelectedEvent.Fights[m.FightCursor]
			if m.PickCursor == 0 {
				m.Picks.Results[fight.Order] = fight.FighterA
			} else {
				m.Picks.Results[fight.Order] = fight.FighterB
			}

			// Rebuild fight list to reflect new pick
			m.FightList = NewFightList(
				m.SelectedEvent.Fights,
				m.Picks.Results,
				m.Styles,
				m.Width,
				m.Height-8,
			)
			m.FightList.Select(m.FightCursor)

			m.Screen = ScreenFightList
		}
		return m, nil
	}

	return m, nil
}

func (m Model) updateSummary(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.Keys.Back):
		m.Screen = ScreenFightList
		return m, nil

	case key.Matches(msg, m.Keys.Enter):
		// Save picks and quit
		m.Picks.Save()
		return m, tea.Quit
	}

	return m, nil
}
