package pages

import (
	"encoding/json"
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"l4d2mm/internal/ui/components"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

func ModsPage() gui.View {
	if internal.GLOBALAPP.ComponentStatus.SubAsideIndices[0] == 0 {
		return ModsView()
	}
	// idx: 1
	return ModDetails()
}

func ModsView() gui.View {
	return gui.Column(gui.ContainerCfg{
		// Width:     internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// MaxWidth:  internal.GLOBALAPP.AppConfig.Width - internal.GLOBALAPP.AppConfig.AsideWidth - 1,
		// Height:    internal.GLOBALAPP.AppConfig.Height,
		// MaxHeight: internal.GLOBALAPP.AppConfig.Height,
		// Spacing: gui.SomeF(1),
		Spacing:      gui.SpacingPx(1),
		Sizing:       gui.FillFill,
		Clip:         true,
		ClipContents: true,
		// Padding: gui.NoPadding,

		Content: []gui.View{
			components.ModsHeader(),
			components.VerticalGap(8, gui.FillFixed),
			components.HDivider(1),
			components.VerticalGap(8, gui.FillFixed),

			gui.Wrap(gui.ContainerCfg{
				ID:     "mods-wrap",
				Sizing: gui.FitFill,
				Wrap:   true,
				// Spacing: gui.SomeF(16),
				Spacing: gui.SpacingPx(16),
				OnScroll: func(ec gui.EventCtx) {

				},
				Scrollable: true,
				ScrollMode: gui.ScrollVerticalOnly,
				Overflow:   false,
				Padding:    gui.NewPadding(6, 6, 0, 6),
				// Wrap:       true,
				// Overflow:   true,
				Clip:         true,
				ClipContents: true,
				Content:      renderCards(),
			}),
		},
	})
}

func ModDetails() gui.View {

	comps := []gui.View{
		components.HeaderModDetails(),
		components.VerticalGap(10, gui.FixedFixed),
		components.HDivider(1),
		components.VerticalGap(10, gui.FixedFixed),
	}

	switch internal.GLOBALAPP.ComponentStatus.SubAsideIndices[0] {
	case 1:
		comps = append(comps, overview())
	case 2:
		// internal.GLOBALAPP.DI.Window.SetView(renderWindow)
		comps = append(comps, renderWindow(internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails.Addoninfo))
	case 4:
		// internal.GLOBALAPP.DI.Window.SetView(renderWindow)
		comps = append(comps, renderWindow(internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails.Missions[0]))
	}

	return gui.Column(gui.ContainerCfg{
		// Spacing: gui.SomeF(1),
		Spacing: gui.SpacingPx(1),
		Sizing:  gui.FillFill,
		HAlign:  gui.HAlignCenter,
		Content: comps,
	})
}

func renderCards() []gui.View {
	// Uncomment this to mock data
	// fakeData()

	t := []gui.View{}

	endIdx, err := strconv.Atoi(internal.GLOBALAPP.ComponentStatus.PageMods.PageEnd[0])
	if err != nil {
		return t
	}

	if endIdx <= 0 {
		return t
	}

	startIdx := (endIdx - 1) / internal.GLOBALAPP.ComponentStatus.PageMods.PageAmount * internal.GLOBALAPP.ComponentStatus.PageMods.PageAmount

	if len(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx) > 0 {
		for _, item := range internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[startIdx:endIdx] {
			// Control the number of elements rendered on the page
			// if idx >= 50 {
			// 	break
			// }
			// fmt.Printf("%+v\n", item)
			t = append(t, components.ModCard(item))
		}
	}

	return t
}

func fakeData() {
	for i := 0; i < 100; i++ {
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, schema.Mod{
			Id: strconv.Itoa(i),
		})
	}
}

// Page modDetails
func overview() gui.View {
	mod := internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails
	return gui.Column(gui.ContainerCfg{
		ID:       "profile-mod-" + mod.Id,
		Width:    750,
		MaxWidth: 750,
		Sizing:   gui.FillFit,
		Padding:  gui.NoPadding,
		// Color:   theme.DefaultLightGNOME().ButtonHover,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		Spacing:     gui.NoSpacing,
		Content: []gui.View{
			// id
			// components.VerticalGap(5, gui.FillFixed),
			modProfileRow1("Id", mod.Id),
			components.HDivider(1),
			modProfileRow2("AddonTitle", internal.GetValueFromMap(mod.Addoninfo, "addontitle", "No title"), 0),
			components.HDivider(1),
			modProfileRow2("AddonAuthor", internal.GetValueFromMap(mod.Addoninfo, "addonauthor", "No author"), 1),
			components.HDivider(1),
			modProfileRow1("AddonVersion", internal.GetValueFromMap(mod.Addoninfo, "addonversion", "No version")),
			// components.HDivider(1),
			// modProfileRow1("Verison", strconv.Itoa(mod.VpkVersion)),
			components.HDivider(1),
			modProfileRow2("URL", internal.GetValueFromMap(mod.Addoninfo, "addonurl0", "No url"), 2),
			components.HDivider(1),
			modProfileRow1("Description", internal.GetValueFromMap(mod.Addoninfo, "addondescription", "No description")),
			// components.HDivider(1),
			// modProfileRow2("Describe", mod.Describe, 2, "No URL"),
			// components.VerticalGap(5, gui.FillFixed),
		},
	})
}

func modProfileRow1(key string, value string) gui.View {
	return gui.Row(gui.ContainerCfg{
		ID:      "page-mod-details-" + key,
		Sizing:  gui.FillFit,
		HAlign:  gui.HAlignLeft,
		VAlign:  gui.VAlignMiddle,
		Padding: gui.NewPadding(8, 12, 8, 12),
		Content: []gui.View{
			components.HorizontalGap(25, gui.FixedFixed),
			gui.Text(gui.TextCfg{
				ID:       "row-" + key,
				Text:     key,
				MinWidth: 100,
				TextStyle: gui.TextStyle{
					CellWidth: 100,
					Size:      theme.DefaultTheme.FontSize,
					Color:     gui.RGBA(40, 40, 40, 255),
				},
			}),
			components.HorizontalSpacer(),
			gui.Text(gui.TextCfg{
				ID:       "row-" + value,
				Text:     value,
				MinWidth: 500,
				Mode:     gui.TextModeWrap,
				TextStyle: gui.TextStyle{
					CellWidth: 190,
					Size:      theme.DefaultTheme.FontSize,
					Color:     gui.RGBA(40, 40, 40, 255),
				},
			}),
			components.HorizontalGap(25, gui.FixedFixed),
		},
	})
}

func modProfileRow2(k string, v string, tidx int) gui.View {
	return gui.Row(gui.ContainerCfg{
		ID:      "page-mod-details-" + k,
		Sizing:  gui.FillFit,
		HAlign:  gui.HAlignLeft,
		VAlign:  gui.VAlignMiddle,
		Padding: gui.NewPadding(8, 12, 8, 12),
		Content: []gui.View{
			components.HorizontalGap(25, gui.FixedFixed),
			gui.Text(gui.TextCfg{
				ID:       "row-" + k,
				Text:     k,
				MinWidth: 100,
				TextStyle: gui.TextStyle{
					CellWidth: 100,
					Size:      theme.DefaultTheme.FontSize,
					Color:     gui.RGBA(40, 40, 40, 255),
				},
			}),
			components.HorizontalSpacer(),
			components.InputPageModDetails(v, tidx),
			components.HorizontalGap(25, gui.FixedFixed),
		},
	})
}

// func renderWindow(w *gui.Window) gui.View {
func renderWindow(data map[string]any) gui.View {
	// jsonBytes, err := json.MarshalIndent(internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails.Addoninfo, "", "    ")
	jsonBytes, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Println("json.MarshalIndent failed:", err)
		jsonBytes = []byte("{}")
	}

	w := internal.GLOBALAPP.DI.Window
	mdSource := fmt.Sprintf("```json\n%s\n```", string(jsonBytes))
	return gui.Column(gui.ContainerCfg{
		ID:         "markdown-addoninfo",
		Sizing:     gui.FillFill,
		Focusable:  true,
		Scrollable: true,
		Clip:       true,
		Content: []gui.View{
			w.Markdown(gui.MarkdownCfg{
				ID:                  "rendeJson",
				Source:              mdSource,
				Style:               gui.DefaultMarkdownStyle(),
				DisableExternalAPIs: true,
				Mode:                gui.Some(gui.TextModeWrap),
				Clip:                true,
			}),
		},
	})
}
