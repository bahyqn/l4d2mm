package internal

import (
	"os"
	"path/filepath"

	"github.com/go-gui-org/go-gui/gui"
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

func GetFileSize(filePath string) (uint64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	return uint64(info.Size()), nil
}

func SetupIcon(iconName string) string {
	return filepath.Join("assets/icons", iconName)
}

func ToastSuccess(title, body string) {
	GLOBALAPP.DI.Window.Toast(gui.ToastCfg{
		Title:    title,
		Body:     body,
		Severity: gui.ToastSuccess,
	})
}

func ToastWarning(title, body string) {
	GLOBALAPP.DI.Window.Toast(gui.ToastCfg{
		Title:    title,
		Body:     body,
		Severity: gui.ToastWarning,
	})
}

func ToastFailed(title, body string) {
	GLOBALAPP.DI.Window.Toast(gui.ToastCfg{
		Title:    title,
		Body:     body,
		Severity: gui.ToastError,
	})
}

func ToastMsg(title, body string) {
	GLOBALAPP.DI.Window.Toast(gui.ToastCfg{
		Title:    title,
		Body:     body,
		Severity: gui.ToastSeverity(gui.ToastTopRight),
	})
}

func AppendBreadCrumb(s string) {
	// GLOBALAPP.ComponentStatus.PageTools.BreadCrumb = append(GLOBALAPP.ComponentStatus.PageTools.BreadCrumb, s)
	GLOBALAPP.ComponentStatus.PageTools.BreadCrumb[1] = s
}

func AreSlicesEqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int, len(a))
	for _, v := range a {
		counts[v]++
	}

	for _, v := range b {
		count, exists := counts[v]
		if !exists || count == 0 {
			return false
		}
		counts[v]--
	}
	return true
}

func BackPageTools() {
	GLOBALAPP.ComponentStatus.SubAsideIndices[GLOBALAPP.ComponentStatus.AsideIdx] = 0
	GLOBALAPP.ComponentStatus.PageTools.BreadCrumb[1] = ""
}
