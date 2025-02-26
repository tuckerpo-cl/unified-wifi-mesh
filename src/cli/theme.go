package main

import (
	"strings"
	"errors"
)

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	AppStyle         lipgloss.Style
	TitleStyle       lipgloss.Style
	MenuBodyStyle    lipgloss.Style
	CanvasStyle      lipgloss.Style
	JSONStyle        lipgloss.Style
	ListItemStyle    lipgloss.Style
	ActiveItemStyle  lipgloss.Style
	ButtonStyle      lipgloss.Style
	ActiveButtonStyle lipgloss.Style
}

type ThemeSelection int
const (
	Unknown ThemeSelection = iota
	Light
	Dark
)

func NewTheme(themeType ThemeSelection) Theme {
	if themeType == Dark {
		return Theme{
			AppStyle: lipgloss.NewStyle().Padding(1, 2),
			TitleStyle: lipgloss.NewStyle().
				Foreground(lipgloss.Color("#c0c0c0")). // Lighter for dark mode
				Bold(true),
			MenuBodyStyle: lipgloss.NewStyle().
				Background(lipgloss.Color("#202020")), // Dark background
			CanvasStyle: lipgloss.NewStyle().
				Background(lipgloss.Color("#303030")),
			JSONStyle: lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Background(lipgloss.Color("#202020")).
				Foreground(lipgloss.Color("#c0c0c0")),
			ListItemStyle: lipgloss.NewStyle().
				Foreground(lipgloss.Color("#aaaaaa")),
			ActiveItemStyle: lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff")).
				Align(lipgloss.Center).
				Bold(true),
			ButtonStyle: lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff")).
				Background(lipgloss.Color("#505050")).
				Padding(0, 1).
				MarginRight(3).
				Width(25).
				Align(lipgloss.Center),
			ActiveButtonStyle: lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff")).
				Padding(0, 1).
				MarginRight(3).
				Width(25).
				Align(lipgloss.Center),
				Background(lipgloss.Color("#c0c0c0")).
				Bold(true),
		}
	}

	// NOTE: add additionally styled themes here

	// Light theme (default)
	return Theme{
		AppStyle: lipgloss.NewStyle().Padding(1, 2),
		TitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#606060")).
			Bold(true),
		MenuBodyStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("#ffffff")),
		CanvasStyle: lipgloss.NewStyle().
			Background(lipgloss.Color("#ebebeb")),
		JSONStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Background(lipgloss.Color("#ffffff")).
			Foreground(lipgloss.Color("#606060")),
		ListItemStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#aaaaaa")),
		ActiveItemStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#606060")).
			Align(lipgloss.Center).
			Bold(true),
		ButtonStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#bfbfbf")).
			Padding(0, 1).
			MarginRight(3).
			Width(25).
			Align(lipgloss.Center).
			MarginBackground(lipgloss.Color("#ebebeb")),
		ActiveButtonStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Padding(0, 1).
			MarginRight(3).
			Width(25).
			Align(lipgloss.Center).
			MarginBackground(lipgloss.Color("#ebebeb")).
			Background(lipgloss.Color("#606060")).
			Bold(true),
	}
}

func ParseTheme(theme string) (ThemeSelection, error) {
	match := strings.ToLower(theme)
	switch {
	case match == "light":
			return Light, nil
	case match == "dark":
			return Dark, nil
	}
	return Unknown, errors.New("Invalid theme.")
}
