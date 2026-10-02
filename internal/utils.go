package internal

import (
	"os"
)

// func ButtonColorChoice(app *internal.App, cidx int, color1 gui.Color, color2 gui.Color) gui.Color {
// 	if app.ComponentStatus.AsideIdx == cidx {
// 		app.ComponentStatus.AsideIdx = cidx
// 		return color1
// 	}
// 	return color2
// }

func StatPath(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else if err == nil {
		return true
	} else {
		return false
	}
}

func GetSubCategories(option string) {
	GLOBALAPP.ComponentStatus.PageMods.SubLabels = ModLables[option]

	if len(GLOBALAPP.ComponentStatus.PageMods.SubLabels) == 0 {
		GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel = []string{}
		return
	}
	GLOBALAPP.ComponentStatus.PageMods.SelectedSubLabel = []string{option}
}
