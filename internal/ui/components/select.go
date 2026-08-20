package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

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

		Color:            theme.DefaultLightGNOME().ViewBackground,
		ColorBorder:      theme.DefaultLightGNOME().ButtonActive,
		ColorBorderFocus: gui.RGBA(0, 0, 0, 90),
		Radius:           gui.SomeF(6),
		SizeBorder:       gui.SomeF(1),
		// v0.51.0
		Padding: gui.NewPadding(6, 10, 6, 10),
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

func Select(maxWidth float32, selectConfig schema.TemplateSelect) gui.View {
	cfg, ok := DefaultSelectConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Select] Invalid select style key")
	}

	cfg.MaxWidth = maxWidth
	cfg.ID = selectConfig.ID
	cfg.Placeholder = selectConfig.Options[0]
	cfg.Selected = selectConfig.Selected
	cfg.Options = selectConfig.Options
	cfg.OnSelect = selectConfig.OnSelectFunc

	return gui.Select(cfg)
}
