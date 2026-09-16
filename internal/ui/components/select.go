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
		MaxWidth:    100,
		MinWidth:    100,
		Sizing:      gui.FillFit,
		ID:          "",
		Placeholder: "",
		Selected:    []string{},
		Options:     []string{},

		Invisible:        false,
		Color:            theme.DefaultLightGNOME().ViewBackground,
		ColorBorder:      theme.DefaultLightGNOME().ButtonActive,
		ColorBorderFocus: gui.RGBA(0, 0, 0, 90),
		Radius:           gui.SomeF(6),
		SizeBorder:       gui.SomeF(1),
		// v0.51.0
		Padding: gui.NewPadding(6, 8, 6, 8),
		// v0.61.0
		// Padding: gui.NewPadding(6, 10, 6, 10),
		TextStyle: gui.TextStyle{
			Size:  10,
			Color: gui.RGBA(150, 150, 150, 255),
		},
		PlaceholderStyle: gui.TextStyle{
			Size:  10,
			Color: gui.RGB(150, 150, 150),
		},
		ColorFocus: theme.DefaultLightGNOME().WindowBackground,
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
		panic("[Select] Invalid select style key")
	}

	cfg.ID = selectConfig.ID
	cfg.MaxWidth = selectConfig.MaxWidth
	cfg.Placeholder = selectConfig.Options[0]
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

	switch internal.GLOBALAPP.ComponentStatus.PageMods.SourceMode[0] {
	case internal.AllSourceModes[0]:
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp

	case internal.AllSourceModes[1]: // Workshop First
		sort.SliceStable(tmp, func(i, j int) bool {
			if tmp[i].IsFromWorkshop != tmp[j].IsFromWorkshop {
				return tmp[i].IsFromWorkshop
			}
			return tmp[i].Idx < tmp[j].Idx
		})
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp

	case internal.AllSourceModes[2]: // Local First
		sort.SliceStable(tmp, func(i, j int) bool {
			if tmp[i].IsFromWorkshop != tmp[j].IsFromWorkshop {
				return !tmp[i].IsFromWorkshop //
			}
			return tmp[i].Idx < tmp[j].Idx
		})
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = tmp
	case internal.AllSourceModes[3]: // Workshop only
		ttmap := []schema.Mod{}

		for _, el := range src {
			if el.IsFromWorkshop {
				ttmap = append(ttmap, el)
			}
		}
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = ttmap
	case internal.AllSourceModes[4]: // Local only
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
