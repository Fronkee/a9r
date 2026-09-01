package utils

import "github.com/gdamore/tcell/v2"

func Getstatecolor(state string) tcell.Color {
	switch state {
	case "running":
		return tcell.ColorGreen
	case "stopped":
		return tcell.ColorRed
	case "pending":
		return tcell.ColorYellow
	case "stopping":
		return tcell.ColorOrange
	default:
		return tcell.ColorWhite
	}
}

func stringValue(v *string) string {
	if v == nil {
		return "-"
	}
	return *v
}

func int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func boolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
