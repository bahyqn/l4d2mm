package components

import (
	"github.com/go-gui-org/go-gui/gui"
)

var DefaultCheckboxCOnfig = map[string]gui.ToggleCfg{
	"gnome": gui.ToggleCfg{},
}

// func CheckBox(cbId string, clickFunc func()) gui.View {
// 	cfg, ok := DefaultButtonConfig[internal.GLOBALAPP.AppConfig.Theme]

// 	if !ok {
// 		panic("CheckBox: Invalid select style key")
// 	}

// 	cfg.ID = cbId
// 	cfg.OnClick = func(ec gui.EventCtx) {
// 		clickFunc()
// 	}

// 	return gui.Checkbox
// }
