package internal

import (
	"l4d2mm/internal/schema"
	"strconv"

	"github.com/go-gui-org/go-gui/gui"
)

var GLOBALAPP *App
var ModLables = map[string][]string{
	"All":         {},
	"Maps":        {},
	"Collections": {},
	"Rifles":      {"M16", "Scar", "AK47", "SG552", "M60"},
	"Shotguns":    {"Punmp", "Chrome", "Auto", "Spas"},
	"Snipers":     {"Hunting", "Military", "Scount", "AWP"},
	"SMG":         {"SMG", "Silenced", "MP5"},
	"Pistols":     {"Pistol", "Magnum"},
	"GL":          {},
	"Melee":       {"ChainSaw", "Fireaxe", "HuntingKnkife", "Katana", "Cricket Bat", "Baseball Bat", "Golfclub", "Machete", "Tonfa", "Electric Guitar", "Frying Pan", "Crowbar", "Riotshield"},
	"Items":       {"First aid kit", "Defibrillator", "Addrenaline", "Pain pills", "Molotov", "Pipebomb", "Vomit Jar", "Gnome", "Cola Bottles"},
	"Survivors":   {"Coach", "Ellis", "Nick", "Rochelle", "Bill", "Francis", "Louis", "Zoey"},
	"Zombies":     {"Boomer", "Charger", "Hunter", "Jockey", "Smoker", "Spitter", "Tank", "Witch", "Zombie"},
	"Scripts":     {},
	"Sounds":      {},
	"Effects":     {},
}

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
			AsideIdx:             0,
			ModsSearchValue:      "",
			Categories:           []string{"All", "Collections", "Maps", "Rifles", "Shotguns", "Snipers", "SMG", "Pistols", "GL", "Melee", "Items", "Scripts", "Sounds", "Effects"},
			SelectedategoryLabel: []string{"All"},
			SubLabels:            []string{},
			SelectedSubLabel:     []string{},

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

	GetSubCategories(GLOBALAPP.ComponentStatus.SelectedategoryLabel[0])
	return GLOBALAPP
}

func (app *App) RegisterAllDependencies(w *gui.Window) {
	// app.DI.Window = w
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
	vpkCount := len(app.DI.Vpk.Mods)

	if vpkCount == 0 {
		return
	}

	if GLOBALAPP.ComponentStatus.PageEnd[0] == "" {
		if len(app.DI.Vpk.Mods) > app.ComponentStatus.PageSize {
			GLOBALAPP.ComponentStatus.PageEnd[0] = strconv.Itoa(app.ComponentStatus.PageSize)
		} else if vpkCount < app.ComponentStatus.PageSize {
			GLOBALAPP.ComponentStatus.PageEnd[0] = strconv.Itoa(vpkCount)
		}

	}

	idx, err := strconv.Atoi(GLOBALAPP.ComponentStatus.PageEnd[0])

	if err != nil {
		panic("DyamicMods xxxxxxxxxxxxx")
	}

	startIdx := idx % app.ComponentStatus.PageSize

	if startIdx == 0 {
		GLOBALAPP.ComponentStatus.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-app.ComponentStatus.PageSize : idx]
	} else {
		GLOBALAPP.ComponentStatus.ModsBySelectIdx = GLOBALAPP.DI.Vpk.Mods[idx-startIdx : idx]
	}

}

func GetValue[T any](m map[string]any, key string) (T, bool) {
	v, ok := m[key].(T)
	if ok {
		return v, true
	}
	return v, false
}
