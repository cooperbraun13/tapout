package tui

import "github.com/charmbracelet/lipgloss"

// UFC-inspired color palette
var (
	// Primary colors
	RedPrimary   = lipgloss.Color("#E40520") // UFC Red
	BlackPrimary = lipgloss.Color("#1A1A1A") // Dark background
	White        = lipgloss.Color("#FFFFFF")
	GrayLight    = lipgloss.Color("#B0B0B0")
	GrayDark     = lipgloss.Color("#4A4A4A")
	GreenSuccess = lipgloss.Color("#22C55E")
	YellowWarn   = lipgloss.Color("#EAB308")
)

// Styles holds all the lipgloss styles for the TUI
type Styles struct {
	// Title styles
	Title    lipgloss.Style
	Subtitle lipgloss.Style

	// Container styles
	Box         lipgloss.Style
	SelectedBox lipgloss.Style

	// Text styles
	Normal   lipgloss.Style
	Muted    lipgloss.Style
	Accent   lipgloss.Style
	Success  lipgloss.Style
	Warning  lipgloss.Style
	Selected lipgloss.Style

	// Fighter card styles
	FighterCard         lipgloss.Style
	FighterCardSelected lipgloss.Style
	FighterName         lipgloss.Style
	FighterStat         lipgloss.Style
	FighterStatLabel    lipgloss.Style

	// List styles
	ListItem         lipgloss.Style
	ListItemSelected lipgloss.Style

	// Help bar
	HelpBar lipgloss.Style
	HelpKey lipgloss.Style
	HelpDesc lipgloss.Style
}

// DefaultStyles returns the default UFC-themed styles
func DefaultStyles() Styles {
	return Styles{
		// Title styles
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(RedPrimary).
			Padding(0, 1).
			MarginBottom(1),

		Subtitle: lipgloss.NewStyle().
			Foreground(GrayLight).
			Padding(0, 1),

		// Container styles
		Box: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(GrayDark).
			Padding(1, 2),

		SelectedBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(RedPrimary).
			Padding(1, 2),

		// Text styles
		Normal: lipgloss.NewStyle().
			Foreground(White),

		Muted: lipgloss.NewStyle().
			Foreground(GrayLight),

		Accent: lipgloss.NewStyle().
			Foreground(RedPrimary).
			Bold(true),

		Success: lipgloss.NewStyle().
			Foreground(GreenSuccess),

		Warning: lipgloss.NewStyle().
			Foreground(YellowWarn),

		Selected: lipgloss.NewStyle().
			Foreground(White).
			Background(RedPrimary).
			Bold(true).
			Padding(0, 1),

		// Fighter card styles
		FighterCard: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(GrayDark).
			Padding(1, 2).
			Width(32),

		FighterCardSelected: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(RedPrimary).
			Padding(1, 2).
			Width(32),

		FighterName: lipgloss.NewStyle().
			Foreground(White).
			Bold(true).
			MarginBottom(1),

		FighterStat: lipgloss.NewStyle().
			Foreground(White),

		FighterStatLabel: lipgloss.NewStyle().
			Foreground(GrayLight).
			Width(10),

		// List styles
		ListItem: lipgloss.NewStyle().
			Foreground(White).
			Padding(0, 2),

		ListItemSelected: lipgloss.NewStyle().
			Foreground(White).
			Background(RedPrimary).
			Bold(true).
			Padding(0, 2),

		// Help bar
		HelpBar: lipgloss.NewStyle().
			Foreground(GrayLight).
			MarginTop(1),

		HelpKey: lipgloss.NewStyle().
			Foreground(RedPrimary).
			Bold(true),

		HelpDesc: lipgloss.NewStyle().
			Foreground(GrayLight),
	}
}
