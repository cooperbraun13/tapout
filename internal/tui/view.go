package tui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	switch m.Screen {
	case ScreenEventList:
		return m.viewEventList()
	case ScreenFightList:
		return m.viewFightList()
	case ScreenFightDetail:
		return m.viewFightDetail()
	case ScreenSummary:
		return m.viewSummary()
	}
	return ""
}

func (m Model) viewEventList() string {
	var b strings.Builder

	b.WriteString("╔══════════════════════════════════════╗\n")
	b.WriteString("║            TAPOUT - EVENTS           ║\n")
	b.WriteString("╚══════════════════════════════════════╝\n\n")

	if len(m.Events) == 0 {
		b.WriteString("  No events found.\n")
	} else {
		for i, event := range m.Events {
			cursor := "  "
			if i == m.EventCursor {
				cursor = "> "
			}

			b.WriteString(fmt.Sprintf("%s┌────────────────────────────────────┐\n", cursor))
			b.WriteString(fmt.Sprintf("  │ %-34s │\n", event.Name))
			b.WriteString(fmt.Sprintf("  │ %-34s │\n", event.Date))
			b.WriteString(fmt.Sprintf("  │ %-34s │\n", event.Location))
			b.WriteString("  └────────────────────────────────────┘\n")
		}
	}

	b.WriteString("\n  [j/k] navigate  [enter] select  [q] quit\n")
	return b.String()
}

func (m Model) viewFightList() string {
	var b strings.Builder

	if m.SelectedEvent == nil {
		return "No event selected"
	}

	e := m.SelectedEvent
	b.WriteString("╔══════════════════════════════════════╗\n")
	b.WriteString(fmt.Sprintf("║ %-36s ║\n", e.Name))
	b.WriteString(fmt.Sprintf("║ %-36s ║\n", e.Date+" - "+e.Venue))
	b.WriteString("╚══════════════════════════════════════╝\n\n")

	for i, fight := range e.Fights {
		cursor := "  "
		if i == m.FightCursor {
			cursor = "> "
		}

		// Check if this fight has a pick
		picked := " "
		if _, ok := m.Picks.Results[fight.Order]; ok {
			picked = "✓"
		}

		b.WriteString(fmt.Sprintf("%s[%s] %s vs %s\n",
			cursor,
			picked,
			fight.FighterA.Name,
			fight.FighterB.Name,
		))
	}

	b.WriteString("\n  [j/k] navigate  [enter] view fight  [s] summary  [esc] back\n")
	return b.String()
}

func (m Model) viewFightDetail() string {
	var b strings.Builder

	if m.SelectedEvent == nil || m.FightCursor >= len(m.SelectedEvent.Fights) {
		return "No fight selected"
	}

	fight := m.SelectedEvent.Fights[m.FightCursor]
	a := fight.FighterA
	bFighter := fight.FighterB

	b.WriteString("╔═══════════════════════════════════════════════════════════════════╗\n")
	b.WriteString(fmt.Sprintf("║ %-29s  VS   %29s ║\n", a.Name, bFighter.Name))
	b.WriteString("╠═══════════════════════════════════════════════════════════════════╣\n")
	b.WriteString(fmt.Sprintf("║ Record:   %-21s │ Record:   %21s ║\n", a.Record, bFighter.Record))
	b.WriteString(fmt.Sprintf("║ Ranking:  %-21s │ Ranking:  %21s ║\n", a.Ranking, bFighter.Ranking))
	b.WriteString(fmt.Sprintf("║ Last 5:   %-21s │ Last 5:   %21s ║\n", a.LastFive, bFighter.LastFive))
	b.WriteString(fmt.Sprintf("║ Division: %-21s │ Division: %21s ║\n", a.Division, bFighter.Division))
	b.WriteString(fmt.Sprintf("║ Odds:     %-21d │ Odds:     %21d ║\n", a.Odds, bFighter.Odds))
	b.WriteString("╚═══════════════════════════════════════════════════════════════════╝\n\n")

	// Selection boxes
	boxA := "[ " + a.Name + " ]"
	boxB := "[ " + bFighter.Name + " ]"

	if m.PickCursor == 0 {
		boxA = "[>" + a.Name + "<]"
	} else {
		boxB = "[>" + bFighter.Name + "<]"
	}

	b.WriteString("  Pick your winner:\n\n")
	b.WriteString(fmt.Sprintf("    %s     %s\n", boxA, boxB))

	b.WriteString("\n  [h/l] switch  [enter] confirm pick  [esc] back\n")
	return b.String()
}

func (m Model) viewSummary() string {
	var b strings.Builder

	if m.SelectedEvent == nil {
		return "No event selected"
	}

	e := m.SelectedEvent
	b.WriteString("╔══════════════════════════════════════╗\n")
	b.WriteString("║           YOUR PICKS SUMMARY         ║\n")
	b.WriteString("╚══════════════════════════════════════╝\n\n")

	b.WriteString(fmt.Sprintf("  Event: %s\n\n", e.Name))

	for _, fight := range e.Fights {
		pick, ok := m.Picks.Results[fight.Order]
		if ok {
			b.WriteString(fmt.Sprintf("  %s vs %s\n", fight.FighterA.Name, fight.FighterB.Name))
			b.WriteString(fmt.Sprintf("    → Winner: %s\n\n", pick.Name))
		} else {
			b.WriteString(fmt.Sprintf("  %s vs %s\n", fight.FighterA.Name, fight.FighterB.Name))
			b.WriteString("    → No pick made\n\n")
		}
	}

	b.WriteString(fmt.Sprintf("  Total picks: %d/%d\n", len(m.Picks.Results), len(e.Fights)))
	b.WriteString("\n  [enter] save & quit  [esc] back\n")
	return b.String()
}
