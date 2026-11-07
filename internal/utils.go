package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"l4d2mm/internal/schema"

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

// GetSubCategories Get subLabels based on option string
func GetSubCategories(option string) {
	pageMods := &GLOBALAPP.ComponentStatus.PageMods
	// clear(pageMods.SubLabels)
	// pageMods.SubLabels = pageMods.SubLabels[:0]
	pageMods.SelectedSubType[0] = ""

	// fmt.Println("")
	sub, ok := ModLables[option]
	if !ok {
		return
	}

	if len(sub) == 0 {
		return
	}

	pageMods.SubTypes = ModLables[option]
	pageMods.SelectedSubType[0] = sub[0].Label
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

// nil: Only len(Header) == len(RowData[i])
func ValidateDimensions(tablleConfig *schema.TableConfig) error {
	headerLen := len(tablleConfig.Header)

	for _, el := range tablleConfig.RowData {
		if len(el) != headerLen {
			return fmt.Errorf("PLease check the raw data.")
		}
	}

	return nil
}

// GenerateSelectOptions Just for ModCategoris & ModLables
func GenerateSelectOptions(key string, appendData []gui.SelectOption) {
	_, ok := ModLables[key]
	if !ok {
		return
	}

	ModLables[key] = append(ModLables[key], appendData...)
}

func IsRelatedPart(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	diff := len1 - len2
	if diff < 0 {
		diff = -diff
	}

	if diff > 2 {
		return false
	}

	minLength := len1
	if len2 < minLength {
		minLength = len2
	}

	matchCount := 0
	for i := 0; i < minLength; i++ {
		if str1[i] == str2[i] {
			matchCount++
		}
	}

	if float64(matchCount)/float64(minLength) >= 0.8 {
		return true
	}

	return false
}

// QueryVPKs query vpks based on category & sublabel
func QueryVPKs() {
	pageMods := &GLOBALAPP.ComponentStatus.PageMods

	clear(pageMods.ModsBySelectIdx)
	pageMods.ModsBySelectIdx = pageMods.ModsBySelectIdx[:0]

	if pageMods.SelectedategoryLabel[0] == "All" {
		mods := GLOBALAPP.DI.Vpk.Mods
		for i := range GLOBALAPP.DI.Vpk.Mods {
			pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &mods[i])
		}

		GLOBALAPP.DynamicPageSelect()
		return
	}

	if pageMods.SelectedategoryLabel[0] == "Collections" {
		if pageMods.SelectedSubType[0] == "" {
			return
		}

		mslice := []string{}
		for _, item := range GLOBALAPP.DI.Vpk.Collections {
			if item.Name == pageMods.SelectedSubType[0] {
				mslice = append(mslice, item.Mods...)
				break
			}
		}

		if len(mslice) > 0 {
			for _, id := range mslice {
				for _, mod := range GLOBALAPP.DI.Vpk.Mods {
					if id == mod.ID {
						pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &mod)
						continue
					}
				}
			}
			GLOBALAPP.DynamicPageSelect()
		}
		return
	}

	if subType, ok := ModLables[pageMods.SelectedategoryLabel[0]]; ok {

		if len(subType) == 0 {
			for _, item := range GLOBALAPP.DI.Vpk.Mods {
				if item.Category == pageMods.SelectedategoryLabel[0] {
					pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &item)
				}
			}
		} else {
			pageMods.SubTypes[0] = subType[0]

			for _, item := range GLOBALAPP.DI.Vpk.Mods {
				if item.Category == pageMods.SelectedategoryLabel[0] && item.SubType == subType[0].Value {
					pageMods.ModsBySelectIdx = append(pageMods.ModsBySelectIdx, &item)
				}
			}
		}
	}
	GLOBALAPP.DynamicPageSelect()
}

func GenerateDynamicCategorySelect() {
	tm := map[string][]string{}

	for _, item := range GLOBALAPP.DI.Vpk.Mods {
		if item.SubType == "" {
			tm[item.Category] = nil
		} else {
			tm[item.Category] = append(tm[item.Category], item.SubType)
		}
	}

	for k, v := range tm {
		if v == nil {
			ModsCategoris = append(ModsCategoris, gui.SelectOption{
				Label: k,
				Value: k,
			})
			continue
		}

		for _, el := range v {
			ModLables[k] = append(ModLables[k], gui.SelectOption{
				Label: el,
				Value: el,
			})
		}
	}
}
