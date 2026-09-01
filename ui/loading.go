package ui

import (
	"github.com/rivo/tview"
)

func ShowLoading(
	app *tview.Application,
	pages *tview.Pages,
	table *tview.Table,
	title string,
) {

	loading := tview.NewTextView()

	loading.
		SetTextAlign(tview.AlignCenter).
		SetText("Loading " + title + "...")

	loading.SetBorder(true)
	loading.SetTitle("A9R")

	center := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(loading, 30, 0, false).
		AddItem(nil, 0, 1, false)

	overlay := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(center, 3, 0, false).
		AddItem(nil, 0, 1, false)

	pages.AddPage(
		"loading",
		overlay,
		true,
		false,
	)

	pages.ShowPage("loading")

	app.SetFocus(loading)
}
