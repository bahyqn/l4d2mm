package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

var DefaultSwitchConfig = map[string]gui.SwitchCfg{
	"gnome": {
		ID:     "",
		Width:  gui.SomeF(33),
		Height: gui.SomeF(20),
		Colors: gui.ColorSet{
			Base:   theme.DefaultLightGNOME().ViewBackground,
			Border: theme.DefaultLightGNOME().ButtonActive,
		},
		ColorSelect:   gui.RGBA(0, 0, 0, 70),
		ColorUnselect: gui.RGBA(0, 0, 0, 15),
	},
}

func SwitchDsiableMod(mod *schema.Mod) gui.View {
	cfg, ok := DefaultSwitchConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("[Switch] Invalid select style key")
	}

	cfg.ID = "mod-switch-" + mod.Id
	cfg.Selected = mod.IsEnable
	cfg.OnClick = func(ec gui.EventCtx) {
		internal.GLOBALAPP.DI.Vpk.DisableMod(mod)
	}

	return gui.Switch(cfg)
}
