package components

import (
	"fmt"
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

var hoverBthId string

var DefaultButtonConfig = map[string]gui.ButtonCfg{
	"gnome": {
		HAlign: gui.Some(gui.HAlignCenter),
		VAlign: gui.Some(gui.VAlignMiddle),
		Sizing: gui.FixedFixed,
		Width:  25,
		Height: 25,

		// Color: theme.DefaultLightGNOME().ButtonDefault,
		Colors: gui.ColorSet{
			Base:   theme.DefaultLightGNOME().ButtonDefault,
			Border: theme.DefaultLightGNOME().BorderColor,
			Click:  theme.DefaultLightGNOME().ButtonActive,
			Focus:  theme.DefaultLightGNOME().ButtonActive,
			Hover:  theme.DefaultLightGNOME().ButtonHover,
		},

		Padding: gui.NewPadding(3, 3, 3, 3),
		// Padding:     gui.NoPadding,
		SizeBorder: gui.NoBorder,
		Radius:     gui.SomeF(8),
		Content:    []gui.View{},
	},
}

func Button(bthID string, clickFunc func(w *gui.Window)) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Button: Invalid select style key")
	}

	cfg.ID = bthID
	cfg.OnClick = func(ec gui.EventCtx) {
		clickFunc(ec.Window)
	}

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
		panic("[Button chooseFolder]Invalid select style key")
	}

	ButtonWithIcon(&cfg, "choose-addons-dir", "folder_open.svg")

	cfg.OnClick = func(ec gui.EventCtx) {
		chooseFolder(ec.Window)
	}
	return gui.Button(cfg)
}

func ButtonShowModInfo(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button save addoninfo]Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-show-mod-info-"+mod.Id, "file_save.svg")

	cfg.OnClick = func(ec gui.EventCtx) {
		ec.Event.IsHandled = true
		// internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
		internal.GLOBALAPP.DI.Vpk.SaveVpkInfo(mod)
	}
	return gui.Button(cfg)
}

func ButtonDisableMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button disable mod]Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-disable-mod-"+mod.Id, "block.svg")

	cfg.OnClick = func(ec gui.EventCtx) {
		ec.Event.IsHandled = true
		// internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
	}
	return gui.Button(cfg)
}

func ButtonDeleteMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button delete Mod]Invalid select style key")
	}

	ButtonWithIcon(&cfg, "bth-delete-mod-"+mod.Id, "delete.svg")

	cfg.OnClick = func(ec gui.EventCtx) {
		ec.Event.IsHandled = true
	}
	return gui.Button(cfg)
}

func ScanButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button delete Mod]Invalid select style key")
	}

	ButtonWithIcon(&cfg, "more_horiz-"+mod.Id, "more_horiz.svg")

	cfg.OnClick = func(ec gui.EventCtx) {}

	cfg.OnHover = func(ec gui.EventCtx) {
		fmt.Println("hovering --->", mod.Id)
	}

	return gui.Button(cfg)
}

func DisplayModProfileButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button length scan]Invalid select style key")
	}

	cfg.ID = "view-profile-" + mod.Id

	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			ID:       "profile-mod-" + mod.Id,
			FileName: "assets/icons/find_in_page.svg",
			Width:    15,
			Height:   15,
		}),
	}

	cfg.OnClick = func(ec gui.EventCtx) {
		fmt.Printf("len --- > %s was clicked.", mod.Id)
	}

	return gui.Button(cfg)
}

func CRCScanButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button CRC]]Invalid select style key")
	}

	cfg.ID = "CRC-scan" + mod.Id

	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			ID:       "crc-mod-" + mod.Id,
			FileName: "assets/icons/fingerprint.svg",
			Width:    15,
			Height:   15,
		}),
	}
	cfg.OnClick = func(ec gui.EventCtx) {
		fmt.Printf("CRC --- > %s was clicked.", mod.Id)
	}

	return gui.Button(cfg)
}

func MoreHoriz(mod *schema.Mod) gui.View {
	return gui.Row(gui.ContainerCfg{
		Width:    50,
		MaxWidth: 50,
		Sizing:   gui.FillFill,
		Padding:  gui.NoPadding,
		Spacing:  gui.SomeF(1),
		Color:    theme.DefaultLightGNOME().ButtonDefault,
		Content: []gui.View{
			DisplayModProfileButton(mod),
			CRCScanButton(mod),
		},
		OnHover: func(ec gui.EventCtx) {
			hoverBthId = mod.Id
		},
	})
}
