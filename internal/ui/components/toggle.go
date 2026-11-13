package components

import (
	"l4d2mm/internal"
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

var DefaultToggleConfig = map[string]gui.ToggleCfg{
	"gnome": {
		ID:         "",
		Label:      "",
		Size:       gui.SomeF(26),
		TextSelect: "✔",
		// SizeBorder: gui.NoBorder,
	},
}

func Toggle(cbId string, tcfg *schema.TableConfig, idx int, minWidth float32) gui.View {
	cfg, ok := DefaultToggleConfig[internal.GLOBALAPP.AppConfig.Theme]

	if !ok {
		panic("Toggle [delete collection]: Invalid select style key")
	}

	cfg.ID = cbId
	cfg.Selected = tcfg.Selected[idx]
	cfg.MinWidth = minWidth
	cfg.OnClick = func(ec gui.EventCtx) {
		tcfg.Selected[idx] = !(tcfg.Selected[idx])
		tcfg.Toggled = true
	}

	return gui.Toggle(cfg)
}
