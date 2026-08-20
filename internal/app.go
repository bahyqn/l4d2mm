package internal

import (
	"l4d2mm/internal/schema"
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
			AsideIdx:         0,
			ModsSearchValue:  "",
			Categories:       []string{""},
			SubLabels:        []string{"Fireaxe", "Katana"},
			SelectedSubLabel: []string{"Fireaxe"},

			PageSize:        50,
			ModsBySelectIdx: []schema.Mod{},
			PageEnd:         []string{""},
		},
		AppConfig: schema.AppConfig{
			Width:      700,
			Height:     500,
			AsideWidth: 130,

			Theme: "gnome",
		},
	}
	return GLOBALAPP
}

func (app *App) RegisterAllDependencies(w *gui.Window) {
	// app.DI.Window = w
	app.DI.Vpk = NewVpk(app)
}

func (app *App) VpkRegister(vpk *Vpk) {
	app.DI.Vpk = vpk
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

func (app *App) DynamicPageSelect() []string {
	modLen := len(app.DI.Vpk.Mods)
	if modLen == 0 {
		return []string{""}
	}

	t := make([]string, 0, modLen/app.ComponentStatus.PageSize+1)

	for i := app.ComponentStatus.PageSize; i < modLen; i += app.ComponentStatus.PageSize {
		t = append(t, strconv.Itoa(i))
	}

	if (modLen % app.ComponentStatus.PageSize) > 0 {
		t = append(t, strconv.Itoa(modLen))
	}

	if app.ComponentStatus.PageEnd[0] == "" {
		app.ComponentStatus.PageEnd = t[:1]

		app.DynamicMods()
	}

	return t
}

func (app *App) DynamicMods() {

	idx, err := strconv.Atoi(GLOBALAPP.ComponentStatus.PageEnd[0])

	if err != nil {
		panic("xxxxxxxxxxxxx")
	}

	startIdx := idx % app.ComponentStatus.PageSize

	if startIdx == 0 {
		GLOBALAPP.ComponentStatus.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-app.ComponentStatus.PageSize : idx]
	} else {
		GLOBALAPP.ComponentStatus.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-startIdx : idx]
	}
}
