package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/cooperbraun13/tapout/internal/event"
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
	// Render the list
	content := m.EventList.View()

	// Help bar
	help := m.renderHelp("↑/↓ navigate", "enter select", "q quit")

	return lipgloss.JoinVertical(lipgloss.Left, content, help)
}

func (m Model) viewFightList() string {
	if m.SelectedEvent == nil {
		return "No event selected"
	}

	// Header with event info
	header := m.Styles.Title.Render(m.SelectedEvent.Name)
	subtitle := m.Styles.Muted.Render(fmt.Sprintf("%s • %s", m.SelectedEvent.Date, m.SelectedEvent.Venue))

	headerBlock := lipgloss.JoinVertical(lipgloss.Left, header, subtitle, "")

	// Fight list
	content := m.FightList.View()

	// Progress
	progress := m.Styles.Muted.Render(
		fmt.Sprintf("Picks: %d/%d", len(m.Picks.Results), len(m.SelectedEvent.Fights)),
	)

	// Help bar
	help := m.renderHelp("↑/↓ navigate", "enter view", "s summary", "esc back")

	return lipgloss.JoinVertical(lipgloss.Left, headerBlock, content, "", progress, help)
}

func (m Model) viewFightDetail() string {
	if m.SelectedEvent == nil || m.FightCursor >= len(m.SelectedEvent.Fights) {
		return "No fight selected"
	}

	fight := m.SelectedEvent.Fights[m.FightCursor]
	a := fight.FighterA
	b := fight.FighterB

	// Title
	title := m.Styles.Title.Render("FIGHT DETAILS")

	// Fighter A card
	fighterAStyle := m.Styles.FighterCard
	if m.PickCursor == 0 {
		fighterAStyle = m.Styles.FighterCardSelected
	}
	fighterACard := m.renderFighterCard(a, fighterAStyle)

	// VS separator
	vs := lipgloss.NewStyle().
		Foreground(RedPrimary).
		Bold(true).
		Padding(4, 2).
		Render("VS")

	// Fighter B card
	fighterBStyle := m.Styles.FighterCard
	if m.PickCursor == 1 {
		fighterBStyle = m.Styles.FighterCardSelected
	}
	fighterBCard := m.renderFighterCard(b, fighterBStyle)

	// Join fighters horizontally
	comparison := lipgloss.JoinHorizontal(lipgloss.Center, fighterACard, vs, fighterBCard)

	// Selection prompt
	prompt := m.Styles.Muted.Render("Pick your winner:")

	// Selection boxes
	boxA := m.renderPickBox(a.Name, m.PickCursor == 0)
	boxB := m.renderPickBox(b.Name, m.PickCursor == 1)
	boxes := lipgloss.JoinHorizontal(lipgloss.Center, boxA, "    ", boxB)

	// Help bar
	help := m.renderHelp("←/→ switch", "enter confirm", "esc back")

	return lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		comparison,
		"",
		prompt,
		boxes,
		"",
		help,
	)
}

func (m Model) viewSummary() string {
	if m.SelectedEvent == nil {
		return "No event selected"
	}

	// Title
	title := m.Styles.Title.Render("YOUR PICKS")
	subtitle := m.Styles.Subtitle.Render(m.SelectedEvent.Name)

	// Build picks list
	var picksContent string
	for _, fight := range m.SelectedEvent.Fights {
		fightLine := fmt.Sprintf("%s vs %s", fight.FighterA.Name, fight.FighterB.Name)

		pick, ok := m.Picks.Results[fight.Order]
		var pickLine string
		if ok {
			indicator := m.Styles.Success.Render("✓")
			winner := m.Styles.Accent.Render(pick.Name)
			pickLine = fmt.Sprintf("  %s %s\n    → %s\n", indicator, fightLine, winner)
		} else {
			indicator := m.Styles.Warning.Render("○")
			pickLine = fmt.Sprintf("  %s %s\n    → %s\n", indicator, fightLine, m.Styles.Muted.Render("No pick"))
		}
		picksContent += pickLine + "\n"
	}

	// Stats
	total := len(m.SelectedEvent.Fights)
	picked := len(m.Picks.Results)
	stats := m.Styles.Muted.Render(fmt.Sprintf("Total: %d/%d picks", picked, total))

	// Container
	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		subtitle,
		"",
		picksContent,
		stats,
	)

	box := m.Styles.Box.Render(content)

	// Help bar
	help := m.renderHelp("enter save & quit", "esc back")

	return lipgloss.JoinVertical(lipgloss.Left, box, "", help)
}

// Helper functions

func (m Model) renderFighterCard(f Fighter, style lipgloss.Style) string {
	name := m.Styles.FighterName.Render(f.Name)

	stats := []string{
		m.renderStat("Record", f.Record),
		m.renderStat("Ranking", f.Ranking),
		m.renderStat("Last 5", f.LastFive),
		m.renderStat("Division", f.Division),
		m.renderStat("Odds", fmt.Sprintf("%d", f.Odds)),
	}

	content := lipgloss.JoinVertical(lipgloss.Left, append([]string{name, ""}, stats...)...)
	return style.Render(content)
}

func (m Model) renderStat(label, value string) string {
	l := m.Styles.FighterStatLabel.Render(label + ":")
	v := m.Styles.FighterStat.Render(value)
	return l + " " + v
}

func (m Model) renderPickBox(name string, selected bool) string {
	style := lipgloss.NewStyle().
		Padding(0, 2).
		Border(lipgloss.RoundedBorder())

	if selected {
		style = style.
			BorderForeground(RedPrimary).
			Foreground(White).
			Background(RedPrimary).
			Bold(true)
	} else {
		style = style.
			BorderForeground(GrayDark).
			Foreground(GrayLight)
	}

	return style.Render(name)
}

func (m Model) renderHelp(bindings ...string) string {
	var parts []string
	for _, b := range bindings {
		parts = append(parts, m.Styles.HelpDesc.Render(b))
	}

	help := ""
	for i, p := range parts {
		if i > 0 {
			help += m.Styles.Muted.Render(" • ")
		}
		help += p
	}

	return m.Styles.HelpBar.Render(help)
}

// Import Fighter type for use in helper
type Fighter = event.Fighter
