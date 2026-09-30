package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

func HeaderModDetails() gui.View {
	tabs := []gui.View{
		ButtonModDetailTab("Overview", 1),
		ButtonModDetailTab("Addoninfo", 2),
		ButtonModDetailTab("Files", 3),
	}

	if len(internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails.Missions) > 0 {
		tabs = append(tabs, ButtonModDetailTab("Missons", 4))
		tabs = append(tabs, ButtonModDetailTab("Map", 5))

	}

	// tabs = append(tabs, VerticalSpacer())
	tabs = append(tabs, HorizontalGap(20, gui.FixedFixed))
	tabs = append(tabs, ButtonReturnModsView())

	return gui.Row(gui.ContainerCfg{
		ID:      "tabs-mod",
		Sizing:  gui.FitFit,
		Padding: gui.NoPadding,
		// Color:   theme.DefaultLightGNOME().ButtonActive,
		Color:   theme.DefaultLightGNOME().ButtonActive,
		Content: tabs,
	})
}
