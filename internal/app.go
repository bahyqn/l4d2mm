package internal

import (
	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

var GLOBALAPP *App

func DefaultAppConfig() schema.AppConfig {
	return schema.AppConfig{
		Id:         1,
		Width:      700,
		Height:     500,
		AsideWidth: 130,
	}
}

type DIContainer struct {
	Sql    *Sqlite3
	Vpk    *Vpk
	Task   *Task
	Window *gui.Window
}

type App struct {
	DI              DIContainer
	ComponentStatus schema.ComponentStatus
	AppConfig       schema.AppConfig
}

func NewApp() *App {
	GLOBALAPP = &App{
		ComponentStatus: schema.ComponentStatus{
			AsideIdx:        0,
			SubAsideIndices: [6]int{0, 0, 0, 0, 0, 0},
			PageMods: schema.PageMods{
				ModsSearchValue: "",
				Categories: []gui.SelectOption{
					{
						Label: "All",
						Value: "All",
					},
					{
						Label: "Collections",
						Value: "Collections",
					},
					{
						Label: "Maps",
						Value: "Maps",
					},
					{
						Label: "Rifles",
						Value: "Rifles",
					},
					{
						Label: "Shotguns",
						Value: "Shotguns",
					},
					{
						Label: "Snipers",
						Value: "Snipers",
					},
					{
						Label: "SMG",
						Value: "SMG",
					},
					{
						Label: "Pistols",
						Value: "Pistols",
					},
					{
						Label: "GL",
						Value: "GL",
					},
					{
						Label: "Melee",
						Value: "Melee",
					},
					{
						Label: "Items",
						Value: "Items",
					},
					{
						Label: "Scripts",
						Value: "Scripts",
					},
					{
						Label: "Sounds",
						Value: "Sounds",
					},
					{
						Label: "Effects",
						Value: "Effects",
					},
				},
				SelectedategoryLabel: []string{"All"},
				SubLabels:            []gui.SelectOption{},
				SelectedSubLabel:     []string{},

				PageAmount:      50,
				ModsBySelectIdx: []schema.Mod{},
				PageEnd:         []string{""},
				SourceMode:      AllSourceModes[0].Value,

				// modDetails
				TempModDetailsFields: [3]string{"", "", ""},
			},
			PageTools: schema.PageTools{},
		},
		AppConfig: schema.AppConfig{
			// Width:      700,
			// Height:     500,
			// AsideWidth: 130,
			Width:      theme.DefaultTheme.AppSize[0],
			Height:     theme.DefaultTheme.AppSize[1],
			AsideWidth: theme.DefaultTheme.AsideMaxSize,

			Theme: "gnome",
		},
		DI: DIContainer{
			Sql:    nil,
			Vpk:    nil,
			Task:   nil,
			Window: nil,
		},
	}

	GetSubCategories(GLOBALAPP.ComponentStatus.PageMods.SelectedategoryLabel[0])
	return GLOBALAPP
}

func (app *App) RegisterAllDependencies(w *gui.Window) {
	app.DI.Window = w
	app.DI.Vpk = NewVpk(app)
	app.DI.Task = NewTask()
}

func (app *App) VpkRegister(vpk *Vpk) {
	app.DI.Vpk = vpk
}

func (app *App) TaskRegister(task *Task) {
	app.DI.Task = task
}

func (app *App) Sql3Register(sql3 *Sqlite3) {
	app.DI.Sql = sql3
}

func (app *App) SetTheme(theme gui.Theme) {
	gui.SetTheme(theme)
}

func (app *App) SetAsideIdx(cidx int) {
	app.ComponentStatus.AsideIdx = cidx
}

func (app *App) ButtonColorChoice(cidx int, color1 gui.Color, color2 gui.Color) gui.Color {
	if app.ComponentStatus.AsideIdx == cidx {
		app.SetAsideIdx(cidx)
		return color1
	}
	return color2
}

func (app *App) DynamicPageSelect() {
	modLen := len(app.ComponentStatus.PageMods.ModsBySelectIdx)

	if modLen == 0 {
		return
	}

	t := make([]gui.SelectOption, 0, modLen/app.ComponentStatus.PageMods.PageAmount+1)

	for i := app.ComponentStatus.PageMods.PageAmount; i < modLen; i += app.ComponentStatus.PageMods.PageAmount {
		t = append(t, gui.SelectOption{
			Label: strconv.Itoa(i),
			Value: strconv.Itoa(i),
		})
	}

	if (modLen % app.ComponentStatus.PageMods.PageAmount) > 0 {
		t = append(t, gui.SelectOption{
			Label: strconv.Itoa(modLen),
			Value: strconv.Itoa(modLen),
		})
	}

	app.ComponentStatus.PageMods.DynamicPageSelect = t

	if len(t) > 0 {
		app.ComponentStatus.PageMods.PageEnd = []string{t[:1][0].Value}
	}
	if len(app.ComponentStatus.PageMods.PageEnd) == 0 ||
		app.ComponentStatus.PageMods.PageEnd[0] == "" {
		app.ComponentStatus.PageMods.PageEnd[0] = ""
	}
}

func SwitchSubPage(idx, subPage int) {
	if idx > len(GLOBALAPP.ComponentStatus.SubAsideIndices) {
		return
	}

	GLOBALAPP.ComponentStatus.SubAsideIndices[idx] = subPage
}

// func (app *App) DynamicMods() {
// 	vpkCount := len(app.DI.Vpk.Mods)
// 	// fmt.Println("vpkCount: ", vpkCount)

// 	if vpkCount == 0 {
// 		return
// 	}

// 	if GLOBALAPP.ComponentStatus.PageMods.PageEnd[0] == "" {
// 		if len(app.DI.Vpk.Mods) > app.ComponentStatus.PageMods.PageAmount {
// 			GLOBALAPP.ComponentStatus.PageMods.PageEnd[0] = strconv.Itoa(app.ComponentStatus.PageMods.PageAmount)
// 		} else if vpkCount < app.ComponentStatus.PageMods.PageAmount {
// 			GLOBALAPP.ComponentStatus.PageMods.PageEnd[0] = strconv.Itoa(vpkCount)
// 		}
// 	}

// 	idx, err := strconv.Atoi(GLOBALAPP.ComponentStatus.PageMods.PageEnd[0])
// 	// fmt.Println("idx: ", idx)

// 	if err != nil {
// 		panic("DyamicMods xxxxxxxxxxxxx")
// 	}

// 	startIdx := idx % app.ComponentStatus.PageMods.PageAmount
// 	// fmt.Println("startIdx: ", startIdx)

// 	if startIdx == 0 {
// 		GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-app.ComponentStatus.PageMods.PageAmount : idx]
// 	} else {
// 		GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-startIdx : idx]
// 	}

// 	// fmt.Printf("%+v\n", GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[0])
// 	// fmt.Printf("%+v\n\n", GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[1])
// }

func GetValue[T any](m map[string]any, key string) (T, bool) {
	v, ok := m[key].(T)
	if ok {
		return v, true
	}
	return v, false
}
