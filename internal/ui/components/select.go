package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"sort"

	"github.com/go-gui-org/go-gui/gui"
)

var DefaultSelectConfig = map[string]gui.SelectCfg{
	"gnome": {
		MaxWidth:    180,
		MinWidth:    180,
		Sizing:      gui.FillFit,
		ID:          "",
		Placeholder: "",
		Selected:    []string{},
		Options:     []gui.SelectOption{},

		Invisible:  false,
		Color:      theme.DefaultLightGNOME().ViewBackground,
		Radius:     gui.RadiusPx(6),
		SizeBorder: gui.BorderPx(1),
		// v0.51.0
		Padding: gui.NewPadding(9, 8, 9, 8),
		// v0.61.0
		// Padding: gui.NewPadding(6, 10, 6, 10),
		TextStyle: gui.TextStyle{
			Size:  theme.DefaultTheme.HeaderFontSize,
			Color: gui.RGBA(150, 150, 150, 255),
		},
		PlaceholderStyle: gui.TextStyle{
			Size:  theme.DefaultTheme.HeaderFontSize,
			Color: gui.RGB(150, 150, 150),
		},
		// v0.51.0
		SubheadingStyle: gui.TextStyle{
			Color: gui.RGBA(150, 150, 150, 255),
		},
		ColorSelect: theme.DefaultLightGNOME().BorderColor,
	},
}

func Select(selectConfig schema.TemplateSelect) gui.View {
	cfg, ok := DefaultSelectConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Select] Invalid select theme key")
	}

	cfg.ID = selectConfig.ID
	cfg.MaxWidth = selectConfig.MaxWidth
	cfg.Placeholder = selectConfig.Options[0].Label
	cfg.Selected = selectConfig.Selected
	cfg.Options = selectConfig.Options
	cfg.Invisible = selectConfig.Invisible
	cfg.OnSelect = selectConfig.OnSelectFunc

	return gui.Select(cfg)
}

func SelectSourceMode() {

	src := internal.GLOBALAPP.DI.Vpk.Mods
	tmp := make([]schema.Mod, len(src))
	copy(tmp, src)

	switch internal.GLOBALAPP.ComponentStatus.PageMods.SourceMode {
	case internal.AllSourceModes[0].Value:
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp

	case internal.AllSourceModes[1].Value: // Workshop First
		sort.SliceStable(tmp, func(i, j int) bool {
			if tmp[i].IsFromWorkshop != tmp[j].IsFromWorkshop {
				return tmp[i].IsFromWorkshop
			}
			return tmp[i].Idx < tmp[j].Idx
		})
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp

	case internal.AllSourceModes[2].Value: // Local First
		sort.SliceStable(tmp, func(i, j int) bool {
			if tmp[i].IsFromWorkshop != tmp[j].IsFromWorkshop {
				return !tmp[i].IsFromWorkshop //
			}
			return tmp[i].Idx < tmp[j].Idx
		})
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp
	case internal.AllSourceModes[3].Value: // Workshop only
		ttmap := []schema.Mod{}

		for _, el := range src {
			if el.IsFromWorkshop {
				ttmap = append(ttmap, el)
			}
		}
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = ttmap
	case internal.AllSourceModes[4].Value: // Local only
		ttmap := []schema.Mod{}

		for _, el := range src {
			if !el.IsFromWorkshop {
				ttmap = append(ttmap, el)
			}
		}

		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = ttmap
	}
	internal.GLOBALAPP.DynamicPageSelect()
}
