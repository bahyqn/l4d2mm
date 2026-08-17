package internal

import (
	"l4d2mm/internal/schema"

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
		},
		AppConfig: schema.AppConfig{
			Width:      700,
			Height:     500,
			AsideWidth: 130,
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
