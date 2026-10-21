package pages

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"

	"github.com/go-gui-org/go-gui/gui"
)

var cardInfo = []schema.ToolCardInfo{
	{
		SvgFilename: "library_add",
		Title:       "Add Collection",
		Describe:    "Create a brand new collection.",
		OnClick: func() {
			updateToolPage(1, "Add Collection")
		},
	},
	{
		SvgFilename: "delete_forever",
		Title:       "Delete Collections",
		Describe:    "Remove an existing collection. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(2, "Delete Collections")
		},
	},
	{
		SvgFilename: "edit_note",
		Title:       "Modify Collection",
		Describe:    "Add or remove mods in existing ones. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(3, "Modify Collection")
		},
	},
	{
		SvgFilename: "category",
		Title:       "Classify Mods",
		Describe:    "Auto-categorize & adjust mod types. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(4, "Classify Mods")
		},
	},
	{
		SvgFilename: "warning",
		Title:       "Conflict Detection",
		Describe:    "Detect conflicts by category or collection. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(5, "Conflict Detective")
		},
	},
	{
		SvgFilename: "edit_document",
		Title:       "Edit GameInfo",
		Describe:    "Configure game search paths & mod loading. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(6, "Edit GameInfo")
		},
	},
	{
		SvgFilename: "arrow_left_alt",
		Title:       "Sync DB to Addonlist",
		Describe:    "Overwrite addonlist.txt using app database states to resolve Cloud Sync issues. (UNAVAILABLE)",
		OnClick: func() {
			// updateToolPage(7, "Sync DB to Addonlist")
		},
	},
	{
		SvgFilename: "arrow_right_alt",
		Title:       "Sync Addonlist to DB",
		Describe:    "Update app database with in-game toggle states and scan for newly added VPKs.",
		OnClick: func() {
			// updateToolPage(8, "Sync Addonlist to DB")
			internal.GLOBALAPP.DI.Vpk.SyncAddonlistToDB()
		},
	},
	{
		SvgFilename: "drive_file_move",
		Title:       "Move VPKs",
		Describe:    "Batch move all VPKs from Steam Workshop folder to the app managed folder. (UNAVAILABLE)",
		OnClick: func() {
			updateToolPage(9, "Move VPKs")
		},
	},
}

func updateToolPage(idx int, s string) {
	internal.GLOBALAPP.ComponentStatus.SubAsideIndices[internal.GLOBALAPP.ComponentStatus.AsideIdx] = idx
	internal.AppendBreadCrumb(s)
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
			Size:     14,
			Color:    gui.RGBA(0, 0, 0, 255),
			Typeface: gui.DefaultTextStyle.Italic().Typeface,
		},
		Items:    bread,
		Selected: s,
		OnSelect: func(s string, ec gui.EventCtx) {
			if s == "Tools" {
				internal.GLOBALAPP.ComponentStatus.SubAsideIndices[internal.GLOBALAPP.ComponentStatus.AsideIdx] = 0
				internal.GLOBALAPP.ComponentStatus.PageTools.BreadCrumb[1] = ""
			}
		},
	}))

	comps = append(comps, components.VerticalGap(3, gui.FixedFixed))
	comps = append(comps, components.HDivider(2))

	// if app.SubAsideIndices.
	switch app.SubAsideIndices[3] {
	case 0:
		comps = append(comps, sectionHeader("COLLECTION MANAGEMENT"))
		comps = append(comps, rowCard(cardInfo[:3]))
		comps = append(comps, sectionHeader("UTILITIES & DIAGNOSTICS"))
		comps = append(comps, rowCard(cardInfo[3:6]))
		comps = append(comps, sectionHeader("MOD & CONFIG SYNCHRONIZATION"))
		comps = append(comps, rowCard(cardInfo[6:]))
	case 1:
		comps = append(comps, addCollection())
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
		Padding: gui.NewPadding(20, 14, 14, 14),
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
		OnClick: func(ec gui.EventCtx) {},
	})
}

// AsideIdx:        3,
// SubAsideIndices: [6]int{0, 0, 0, 1, 0, 0},
func addCollection() gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:      "add-collection",
		Sizing:  gui.FillFit,
		Padding: gui.NewPadding(14, 200, 14, 200),
		HAlign:  gui.HAlignCenter,
		Content: []gui.View{
			gui.Row(gui.ContainerCfg{
				ID:     "",
				Sizing: gui.FillFill,
				Content: []gui.View{
					components.InputCollectionName(),
					components.TextModsAcount(),
					components.VerticalSpacer(),
					components.ButtonAddCollection(),
				},
			}),
			gui.Row(gui.ContainerCfg{
				ID:     "add-collection-describe-block",
				Sizing: gui.FillFit,
				Content: []gui.View{
					gui.Toggle(gui.ToggleCfg{
						ID:         "toggle-enable-collection",
						TextSelect: "✔",
						Selected:   internal.GLOBALAPP.ComponentStatus.PageTools.ApplyCollection,
						Size:       gui.SomeF(30),
						OnClick: func(ec gui.EventCtx) {
							internal.GLOBALAPP.ComponentStatus.PageTools.ApplyCollection = !internal.GLOBALAPP.ComponentStatus.PageTools.ApplyCollection
						},
					}),
					components.HorizontalGap(10, gui.FixedFill),
					components.SettingRow(schema.SettingRowConfig{
						ColumnID:           "column-enable-collection",
						PrimaryTextID:      "toggle-enable-collection-title",
						PrimaryTextLabel:   "Apply collection",
						SecondaryTextID:    "toggle-enable-collection-description",
						SecondaryTextLabel: "Check this collection first to apply it when added",
					}),
				},
			}),
		},
	})
}
