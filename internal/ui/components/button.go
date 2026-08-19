package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
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

		Color:      theme.DefaultLightGNOME().ButtonDefault,
		ColorClick: theme.DefaultLightGNOME().ButtonActive,
		// v0.51.0
		ColorBorder: theme.DefaultLightGNOME().BorderColor,
		// v0.51.0
		ColorHover: theme.DefaultLightGNOME().ButtonHover,
		ColorFocus: theme.DefaultLightGNOME().ButtonActive,
		SizeBorder: gui.NoBorder,
		// Padding:     gui.NoPadding,
		Radius:  gui.SomeF(8),
		Content: []gui.View{},
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

func ButtonWithIcon(cfg *gui.ButtonCfg, bthID string, iconName string) {

	cfg.ID = bthID
	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			Width:    15,
			Height:   15,
			Sizing:   gui.FixedFixed,
			FileName: "assets/icons/" + iconName,
		}),
	}
}
func ButtonWithIconText(cfg *gui.ButtonCfg, bthID string, iconName string, text string) {

	cfg.ID = bthID
	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			Width:    15,
			Height:   15,
			Sizing:   gui.FixedFixed,
			FileName: "assets/icons/" + iconName,
		}),
		gui.Text(gui.TextCfg{}),
	}
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

	ButtonWithIcon(&cfg, "choose-addons-dir", "folder_open.svg")

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

func ButtonShowModInfo(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-show-mod-info-"+mod.Id, "file_save.svg")

	// cfg.ID = "bth-show-mod-info-" + mod.Id

	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		e.IsHandled = true
		internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
	}
	return gui.Button(cfg)
}

func ButtonDisableMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-disable-mod-"+mod.Id, "block.svg")

	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		e.IsHandled = true
		// internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
	}
	return gui.Button(cfg)
}

func ButtonDeleteMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-delete-mod-"+mod.Id, "delete.svg")

	cfg.OnClick = func(l *gui.Layout, e *gui.Event, w *gui.Window) {
		e.IsHandled = true
		// internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
	}
	return gui.Button(cfg)
}
