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
		Width:  38,
		Height: 38,

		// Color: theme.DefaultLightGNOME().ButtonDefault,
		Colors: gui.ColorSet{
			Base:   theme.DefaultLightGNOME().ButtonDefault,
			Border: theme.DefaultLightGNOME().BorderColor,
			Click:  theme.DefaultLightGNOME().ButtonActive,
			Focus:  theme.DefaultLightGNOME().ButtonActive,
			Hover:  theme.DefaultLightGNOME().ButtonHover,
		},

		Padding: gui.NewPadding(5, 5, 5, 5),
		// Padding:     gui.NoPadding,
		SizeBorder: gui.NoBorder,
		// Radius:     gui.SomeF(8),
		Radius:  gui.RadiusPx(8),
		Content: []gui.View{},
	},
}

func Button(bthID string, clickFunc func(w *gui.Window)) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Button: Invalid select theme key")
	}

	cfg.ID = bthID
	cfg.OnClick = func(ec gui.EventCtx) {
		clickFunc(ec.Window)
	}

	return gui.Button(cfg)
}

func ButtonWithIcon(cfg *gui.ButtonCfg, bthID string, svgId string, iconName string, svgSize [2]float32) {

	cfg.ID = bthID
	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			ID:       svgId,
			Width:    svgSize[0],
			Height:   svgSize[1],
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
		panic("[Button chooseFolder]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "choose-addons-dir", "icon-folder-open", "folder_open.svg", theme.DefaultTheme.HeaderIconSize)

	cfg.OnClick = func(ec gui.EventCtx) {
		chooseFolder(ec.Window)
	}
	return gui.Button(cfg)
}

func ButtonShowModInfo(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button save addoninfo]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "bth-show-mod-info-"+mod.Id, "icon-file-save", "file_save.svg", theme.DefaultTheme.HeaderIconSize)

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
		panic("[Button disable mod]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "bth-disable-mod-"+mod.Id, "icon-disable-mod", "block.svg", theme.DefaultTheme.ModCardIconSize)

	cfg.OnClick = func(ec gui.EventCtx) {
		ec.Event.IsHandled = true
		// internal.GLOBALAPP.DI.Vpk.ReadVpkInfo(mod)
	}
	return gui.Button(cfg)
}

func ButtonDeleteMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button delete Mod]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "bth-delete-mod-"+mod.Id, "icon-delete-mod", "delete.svg", theme.DefaultTheme.ModCardIconSize)

	cfg.OnClick = func(ec gui.EventCtx) {
		ec.Event.IsHandled = true
	}
	return gui.Button(cfg)
}

func ScanButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button delete Mod]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "more_horiz-"+mod.Id, "icon-more", "more_horiz.svg", theme.DefaultTheme.ModCardIconSize)

	cfg.OnClick = func(ec gui.EventCtx) {}

	cfg.OnHover = func(ec gui.EventCtx) {
		fmt.Println("hovering --->", mod.Id)
	}

	return gui.Button(cfg)
}

func DisplayModProfileButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button length scan]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "view-overview-"+mod.Id, "icon-overview", "find_in_page.svg", theme.DefaultTheme.ModCardIconSize)
	// cfg.ID = "view-profile-" + mod.Id

	// cfg.Content = []gui.View{
	// 	gui.Svg(gui.SvgCfg{
	// 		ID:       "profile-mod-" + mod.Id,
	// 		FileName: "assets/icons/find_in_page.svg",
	// 		Width:    theme.DefaultTheme.ModCardIconSize[0],
	// 		Height:   theme.DefaultTheme.ModCardIconSize[1],
	// 	}),
	// }

	cfg.OnClick = func(ec gui.EventCtx) {
		// fmt.Printf("len --- > %s was clicked.", mod.Id)
		internal.GLOBALAPP.ComponentStatus.PageMods.ModDetails = mod
		internal.SwitchSubPage(0, 1)
	}

	return gui.Button(cfg)
}

func CRCScanButton(mod *schema.Mod) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button CRC]]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "CRC-scan-"+mod.Id, "icon-crc-scan", "fingerprint.svg", theme.DefaultTheme.ModCardIconSize)
	// cfg.ID = "CRC-scan" + mod.Id

	// cfg.Content = []gui.View{
	// 	gui.Svg(gui.SvgCfg{
	// 		ID:       "crc-mod-" + mod.Id,
	// 		FileName: "assets/icons/fingerprint.svg",
	// 		Width:    15,
	// 		Height:   15,
	// 	}),
	// }
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
		// Spacing:  gui.SomeF(1),
		Spacing: gui.SpacingPx(1),
		Color:   theme.DefaultLightGNOME().ButtonDefault,
		Content: []gui.View{
			DisplayModProfileButton(mod),
			CRCScanButton(mod),
		},
		OnHover: func(ec gui.EventCtx) {
			hoverBthId = mod.Id
		},
	})
}

func SelectAllButton() gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button selectAll]]Invalid select theme key")
	}

	cfg.ID = "select-all-mods"

	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			ID:       "select-all-svg",
			FileName: "assets/icons/check_box.svg",
			Width:    theme.DefaultTheme.HeaderIconSize[0],
			Height:   theme.DefaultTheme.HeaderIconSize[1],
		}),
	}
	cfg.OnClick = func(ec gui.EventCtx) {
		if len(internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx) == 0 {
			return
		}

		for _, el := range internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx {
			internal.GLOBALAPP.DI.Task.AddTask(el.Id)
		}
	}
	return gui.Button(cfg)
}

func CancelAllButton() gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button cancelAll]]Invalid select theme key")
	}

	cfg.ID = "remove-all-mods"

	cfg.Content = []gui.View{
		gui.Svg(gui.SvgCfg{
			ID:       "remove-all-svg",
			FileName: "assets/icons/check_box_outline_blank.svg",
			Width:    theme.DefaultTheme.HeaderIconSize[0],
			Height:   theme.DefaultTheme.HeaderIconSize[1],
		}),
	}
	cfg.OnClick = func(ec gui.EventCtx) {
		if len(internal.GLOBALAPP.DI.Task.Selected) == 0 {
			return
		}

		for _, el := range internal.GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx {
			internal.GLOBALAPP.DI.Task.RemoveTask(el.Id)
		}
	}
	return gui.Button(cfg)
}

func ButtonModDetailTab(tabName string, subPage int) gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button mod detail tab]]Invalid select theme key")
	}

	cfg.ID = "tab-" + tabName + "-mod"

	cfg.Width = 80
	cfg.Height = 36

	cfg.Content = []gui.View{
		gui.Text(gui.TextCfg{
			Text: tabName,
			TextStyle: gui.TextStyle{
				Size:  theme.DefaultTheme.HeaderTabFontSize,
				Color: gui.RGBA(40, 40, 40, 255),
			},
		}),
	}
	status := internal.GLOBALAPP.ComponentStatus

	if status.AsideIdx == 0 && status.SubAsideIndices[0] == subPage {
		cfg.Colors.Base = gui.RGBA(255, 255, 255, 255)
		cfg.Colors.Click = gui.RGBA(255, 255, 255, 255)
		cfg.Colors.Focus = gui.RGBA(255, 255, 255, 255)
		cfg.Colors.Hover = gui.RGBA(255, 255, 255, 255)
	} else {
		cfg.Colors.Base = gui.RGBA(0, 0, 0, 0)
		cfg.Colors.Click = gui.RGBA(0, 0, 0, 0)
		cfg.Colors.Focus = gui.RGBA(0, 0, 0, 0)
		cfg.Colors.Hover = gui.RGBA(0, 0, 0, 0)
	}

	cfg.OnClick = func(ec gui.EventCtx) {
		internal.SwitchSubPage(0, subPage)
	}

	return gui.Button(cfg)
}

func ButtonReturnModsView() gui.View {
	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Button cancelAll]]Invalid select theme key")
	}

	ButtonWithIcon(&cfg, "return-mods-view", "return-mods-view-svg", "close.svg", theme.DefaultTheme.ModCardIconSize)
	// cfg.ID = "return-mods-view"

	cfg.Colors.Base = theme.DefaultLightGNOME().ButtonActive

	// cfg.Content = []gui.View{
	// 	gui.Svg(gui.SvgCfg{
	// 		ID:       "return-mods-view-svg",
	// 		FileName: "assets/icons/close.svg",
	// 		Width:    15,
	// 		Height:   15,
	// 	}),
	// }
	cfg.OnClick = func(ec gui.EventCtx) {
		internal.SwitchSubPage(0, 0)
	}
	return gui.Button(cfg)
}
