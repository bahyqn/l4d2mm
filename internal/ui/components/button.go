package components

import (
	"fmt"
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

		Color:       theme.DefaultLightGNOME().ButtonDefault,
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		SizeBorder:  gui.NoBorder,
		ColorHover:  theme.DefaultLightGNOME().ButtonActive,
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

func Button(k string, clickFunc func(w *gui.Window)) gui.View {
	cfg, ok := DefaultButtonConfig[k]

	if !ok {
		panic("Invalid select style key")
	}

	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		clickFunc(w)
	}
	return gui.Button(cfg)
}

func chooseFolder(w *gui.Window) {
	w.NativeFolderDialog(gui.NativeFolderDialogCfg{
		Title:    "Select Folder",
		StartDir: "",
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

			fmt.Println("Selected folder:", folder)
		},
	})
}

func ButtonChooseFolder(k string) gui.View {
	cfg, ok := DefaultButtonConfig[k]

	if !ok {
		panic("Invalid select style key")
	}

	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		chooseFolder(w)
	}
	return gui.Button(cfg)
}
