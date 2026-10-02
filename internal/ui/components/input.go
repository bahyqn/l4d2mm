package components

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

// func OnTextChanged needs to be parameter!!!
var DefaultInputConfig = map[string]gui.InputCfg{
	"gnome": {
		ID:       "mods-searchinput",
		Width:    180,
		MaxWidth: 180,
		// Height:      38,
		Height:      theme.DefaultTheme.HeaderMaxHeight,
		MaxHeight:   theme.DefaultTheme.HeaderMaxHeight,
		Sizing:      gui.FillFill,
		Placeholder: "Search mods...",
		SpellCheck:  true,

		Color: theme.DefaultLightGNOME().ViewBackground,
		Colors: gui.ColorSet{
			Base:        theme.DefaultLightGNOME().ViewBackground,
			Hover:       theme.DefaultLightGNOME().ViewBackground,
			Border:      theme.DefaultLightGNOME().ButtonActive,
			BorderFocus: gui.RGBA(0, 0, 0, 90),
		},
		Radius:     gui.RadiusPx(6),
		SizeBorder: gui.BorderPx(1),
		Padding:    gui.NewPadding(6, 10, 6, 10),

		TextStyle: gui.TextStyle{
			Size:  theme.DefaultTheme.HeaderFontSize,
			Color: gui.RGBA(150, 150, 150, 255),
		},
		PlaceholderStyle: gui.TextStyle{
			Size:  theme.DefaultTheme.HeaderFontSize,
			Color: gui.RGB(150, 150, 150),
		},
		// OnTextChanged: func(s string, ec gui.EventCtx) {
		// 	internal.GLOBALAPP.ComponentStatus.PageMods.ModsSearchValue = s
		// 	fmt.Println(s)
		// },
	},
}

func Input(text string) gui.View {
	cfg, ok := DefaultInputConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("invalid input theme key")
	}

	cfg.Text = text
	cfg.OnTextChanged = func(s string, ec gui.EventCtx) {
		internal.GLOBALAPP.ComponentStatus.PageMods.ModsSearchValue = s
		fmt.Println(s)
	}
	cfg.OnTextCommit = func(s string, icr gui.InputCommitReason, ec gui.EventCtx) {
		// internal.GLOBALAPP.ComponentStatus.PageMods.ModsSearchValue = s
		fmt.Println("commit: ", s)
	}

	return gui.Input(cfg)
}

// tidx: {title, author, url}
func InputPageModDetails(text string, tidx int, placeHolder string) gui.View {
	cfg, ok := DefaultInputConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Input pageModDetails]invalid input theme key")
	}

	cfg.Width = 200
	cfg.MaxWidth = 200
	cfg.Text = text

	if text == "" {
		cfg.Placeholder = placeHolder
	}
	cfg.OnTextChanged = func(s string, ec gui.EventCtx) {
		internal.GLOBALAPP.ComponentStatus.PageMods.TempModDetailsFields[tidx] = s
		fmt.Println(s)
	}
	cfg.OnTextCommit = func(s string, icr gui.InputCommitReason, ec gui.EventCtx) {
		// internal.GLOBALAPP.ComponentStatus.PageMods.ModsSearchValue = s
		fmt.Println("commit: ", internal.GLOBALAPP.ComponentStatus.PageMods.TempModDetailsFields[tidx])
	}

	return gui.Input(cfg)
}
