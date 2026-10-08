package main

import (
	"slices"

	"github.com/aungshanbo/a9r/configs"
	"github.com/aungshanbo/a9r/models"
	"github.com/aungshanbo/a9r/ui"
	"github.com/rivo/tview"
)

func main() {

	profiles, profileRegions := configs.Loadconfig()
	regions := configs.GetRegions()
	resources := configs.GetResources()

	for _, region := range profileRegions {

		if region != "" &&
			!slices.Contains(regions, region) {

			regions = append(regions, region)
		}
	}

	app := tview.NewApplication()

	state := &models.AppState{}

	var selectedProfile string
	var selectedRegion string
	var selectedResource string

	autoRefresh := true

	table := ui.NewTable()

	statusBar := tview.NewTextView()

	statusBar.SetBorder(true)
	statusBar.SetTitle("Help")

	leftPanel,
		profileDropDown,
		regionDropDown,
		resourceDropDown,
		statusView := ui.NewForm(
		app,
		table,
		profiles,
		profileRegions,
		regions,
		resources,
		&selectedProfile,
		&selectedRegion,
		&selectedResource,
	)

	layout := ui.NewLayout(
		leftPanel,
		table,
		statusBar,
	)

	pages := tview.NewPages()

	pages.AddPage(
		"main",
		layout,
		true,
		true,
	)

	ui.BindTableKeys(
		app,
		table,
		pages,
		state,
		&selectedProfile,
		&selectedRegion,
		&selectedResource,
	)

	ui.BindGlobalKeys(
		app,
		table,
		profileDropDown,
		regionDropDown,
		resourceDropDown,
		statusBar,
		statusView,
		state,
		&selectedProfile,
		&selectedRegion,
		&selectedResource,
		&autoRefresh,
	)

	ui.StartAutoRefresh(
		app,
		table,
		statusView,
		state,
		&selectedProfile,
		&selectedRegion,
		&selectedResource,
		&autoRefresh,
	)

	app.SetFocus(profileDropDown)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
