package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)
 
var DefaultButtonConfig = map[string]gui.ButtonCfg{
	"gnome": {
		HAlign: gui.Some(gui.HAlignCenter),
		VAlign: gui.Some(gui.VAlignMiddle),
		Sizing: gui.FixedFixed,
		Width:  25,
		Height: 25,

		Color: theme.DefaultLightGNOME().ButtonDefault,
		// v0.51.0
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		// v0.51.0
		ColorHover: theme.DefaultLightGNOME().ButtonActive,
		SizeBorder: gui.NoBorder,
		// Padding:     gui.NoPadding,
		Radius: gui.SomeF(8),
		Content: []gui.View{
			gui.Svg(gui.SvgCfg{
				Width:    15,
				Height:   15,
				Sizing:   gui.FixedFixed,
				FileName: "assets/icons/folder_open.svg",
			}),
		},
	},
}

func Button(bthID string, clickFunc func(w *gui.Window)) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Invalid select style key")
	}

	cfg.ID = bthID
	// v0.51.0
	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		clickFunc(w)
	}
	// v0.61.0
	// cfg.OnClick = func(ec gui.EventCtx) {
	// 	clickFunc(ec.Window)
	// }
	return gui.Button(cfg)
}

func chooseFolder(w *gui.Window) {
	w.NativeFolderDialog(gui.NativeFolderDialogCfg{
		Title: "Select Folder",
		// v0.51.0
		// StartDir: "",
		OnDone: func(ndr gui.NativeDialogResult, w *gui.Window) {
			if ndr.Status != gui.DialogOK {
				return
			}

			paths := ndr.PathStrings()

			if len(paths) == 0 {
				return
			}

			folder := paths[0]

			vpk := internal.GLOBALAPP.DI.Vpk

			vpk.SetupPath(folder)
		},
	})
}

func ButtonChooseFolder() gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Invalid select style key")
	}

	cfg.ID = "choose-addons-dir"
	// v0.51.0
	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		chooseFolder(w)
	}
	// v0.61.0
	// cfg.OnClick = func(ec gui.EventCtx) {
	// 	chooseFolder(ec.Window)
	// }
	return gui.Button(cfg)
}
