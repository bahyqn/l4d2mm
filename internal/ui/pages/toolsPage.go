package pages

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"

	"github.com/go-gui-org/go-gui/gui"
)

var cardInfo = []schema.ToolCardInfo{
	{
		SvgFilename: "library_add.svg",
		Title:       "Add Collection",
		Describe:    "Create a brand new collection. (UNAVAILABLE)",
	},
	{
		SvgFilename: "delete_forever.svg",
		Title:       "Delete Collection",
		Describe:    "Remove an existing collection. (UNAVAILABLE)",
	},
	{
		SvgFilename: "edit_note.svg",
		Title:       "Modify Collection",
		Describe:    "Add or remove mods in existing ones. (UNAVAILABLE)",
	},
	{
		SvgFilename: "category.svg",
		Title:       "Classify Mods",
		Describe:    "Auto-categorize & adjust mod types. (UNAVAILABLE)",
	},
	{
		SvgFilename: "warning.svg",
		Title:       "Conflict Detection",
		Describe:    "Detect conflicts by category or collection. (UNAVAILABLE)",
	},
	{
		SvgFilename: "edit_document.svg",
		Title:       "Edit GameInfo",
		Describe:    "Configure game search paths & mod loading. (UNAVAILABLE)",
	},
	{
		SvgFilename: "arrow_left_alt.svg",
		Title:       "Sync DB to Addonlist",
		Describe:    "Overwrite addonlist.txt using app database states to resolve Cloud Sync issues. (UNAVAILABLE)",
	},
	{
		SvgFilename: "arrow_right_alt.svg",
		Title:       "Sync Addonlist to DB",
		Describe:    "Update app database with in-game toggle states and scan for newly added VPKs.",
	},
	{
		SvgFilename: "drive_file_move.svg",
		Title:       "Move VPKs",
		Describe:    "Batch move all VPKs from Steam Workshop folder to the app managed folder. (UNAVAILABLE)",
	},
}

var s string

func ToolsPage() gui.View {
	app := internal.GLOBALAPP.ComponentStatus

	bread := []gui.BreadcrumbItemCfg{}
	comps := []gui.View{}

	for _, el := range app.PageTools.BreadCrumb {
		bread = append(bread, gui.NewBreadcrumbItem(el, el, nil))
	}

	comps = append(comps, gui.Breadcrumb(gui.BreadcrumbCfg{
		ID: "tools-nav",
		TextStyle: gui.TextStyle{
			Size:     16,
			Color:    gui.RGBA(0, 0, 0, 255),
			Typeface: gui.DefaultTextStyle.Italic().Typeface,
		},
		Items:    bread,
		Selected: s,
		OnSelect: func(s string, ec gui.EventCtx) {
			fmt.Println(s)
		},
	}))

	// if app.SubAsideIndices.
	if app.SubAsideIndices[3] == 0 {
		comps = append(comps, components.VerticalGap(3, gui.FixedFixed))
		comps = append(comps, components.HDivider(2))
		comps = append(comps, sectionHeader("COLLECTION MANAGEMENT"))
		comps = append(comps, rowCard(cardInfo[:3]))
		comps = append(comps, sectionHeader("UTILITIES & DIAGNOSTICS"))
		comps = append(comps, rowCard(cardInfo[3:6]))
		comps = append(comps, sectionHeader("MOD & CONFIG SYNCHRONIZATION"))
		comps = append(comps, rowCard(cardInfo[6:]))
	}

	return gui.Column(gui.ContainerCfg{
		ID:      "page-tools",
		Sizing:  gui.FillFill,
		Spacing: gui.NoSpacing,
		Content: comps,
	})
}

func sectionHeader(text string) gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:      "section-header-" + text,
		Padding: gui.NewPadding(30, 14, 14, 14),
		Content: []gui.View{
			gui.Text(gui.TextCfg{
				ID:   "text-" + text,
				Text: text,
				TextStyle: gui.TextStyle{
					Size:  16,
					Color: theme.DefaultTheme.SecondaryLabel,
				},
			}),
		},
	})
}

func rowCard(slice []schema.ToolCardInfo) gui.View {
	comps := []gui.View{}

	for _, el := range slice {
		comps = append(comps, components.ToolCard(el))
	}

	return gui.Row(gui.ContainerCfg{
		ID:      "tools-card-warp",
		Sizing:  gui.FitFit,
		Wrap:    true,
		Spacing: gui.SpacingPx(20),
		Padding: gui.NewPadding(0, 14, 0, 14),
		Content: comps,
		OnClick: func(ec gui.EventCtx) {
			internal.GLOBALAPP.DI.Vpk.SyncAddonlistToDB()
		},
	})
}
