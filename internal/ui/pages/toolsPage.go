package pages

import (
	"fmt"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"

	"github.com/go-gui-org/go-gui/gui"
)

var cardInfo = []schema.ToolCardInfo{
	{
		SvgFilename: "library_add.svg",
		Title:       "Add Collection",
		Describe:    "Create a brand new collection",
	},
	{
		SvgFilename: "delete_forever.svg",
		Title:       "Delete Collection",
		Describe:    "Remove an existing collection",
	},
	{
		SvgFilename: "edit_note.svg",
		Title:       "Modify Collection",
		Describe:    "Add or remove mods in existing ones",
	},
	{
		SvgFilename: "category.svg",
		Title:       "Classify Mods",
		Describe:    "Auto-categorize & adjust mod types",
	},
	{
		SvgFilename: "warning.svg",
		Title:       "Conflict Detection",
		Describe:    "Detect conflicts by category or collection",
	},
}

var s string

func ToolsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		ID:      "page-tools",
		Sizing:  gui.FillFill,
		Spacing: gui.NoSpacing,
		Content: []gui.View{
			gui.Breadcrumb(gui.BreadcrumbCfg{
				ID: "tools-nav",
				TextStyle: gui.TextStyle{
					Size:     16,
					Color:    gui.RGBA(0, 0, 0, 255),
					Typeface: gui.DefaultTextStyle.Italic().Typeface,
				},
				Items: []gui.BreadcrumbItemCfg{
					gui.NewBreadcrumbItem("tool", "Tools", nil),
					gui.NewBreadcrumbItem("docs", "Docs", nil),
					gui.NewBreadcrumbItem("api", "API Reference", nil),
				},
				Selected: s,
				OnSelect: func(s string, ec gui.EventCtx) {
					fmt.Println(s)
				},
			}),

			components.VerticalGap(3, gui.FixedFixed),
			components.HDivider(2),

			sectionHeader("COLLECTION MANAGEMENT"),
			rowCard(cardInfo[:3]),
			sectionHeader("UTILITIES & DIAGNOSTICS"),
			rowCard(cardInfo[3:]),
		},
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
	})
}
