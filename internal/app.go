package internal

import (
	"embed"
	"path"
	"strconv"
	"strings"

	"l4d2mm/internal/schema"
	"l4d2mm/internal/theme"

	"github.com/go-gui-org/go-gui/gui"
)

var GLOBALAPP *App

func DefaultAppConfig() schema.AppConfig {
	return schema.AppConfig{
		ID:       1,
		Theme:    "gnome",
		GamePath: "",
	}
}

type DIContainer struct {
	Window *gui.Window
	SQL    *Sqlite3
	Vpk    *Vpk
	Task   *Task
}

type App struct {
	DI              DIContainer            `gorm:"-"`
	ComponentStatus schema.ComponentStatus `gorm:"-"`
	AppConfig       schema.AppConfig       `gorm:"-"`
	assetfile       embed.FS               `gorm:"-"`
}

func NewApp(as embed.FS) *App {
	GLOBALAPP = &App{
		assetfile: as,
		ComponentStatus: schema.ComponentStatus{
			AsideIdx:        0,
			SubAsideIndices: [6]int{0, 0, 0, 0, 0, 0},
			PageMods: schema.PageMods{
				ModsSearchValue:      "",
				Categories:           ModsCategoris,
				SelectedategoryLabel: []string{"All"},
				SubLabels:            []gui.SelectOption{},
				SelectedSubLabel:     []string{""},

				PageAmount:      50,
				ModsBySelectIdx: []schema.Mod{},
				PageEnd:         []string{""},
				SourceMode:      AllSourceModes[0].Value,

				// modDetails: addontile, addonauthor, addonurl0
				TempModDetailsFields: [3]string{"", "", ""},
			},
			PageTools: schema.PageTools{
				BreadCrumb:      [2]string{"Tools"},
				CollectionName:  "",
				ApplyCollection: false,
				DeleteCollection: schema.TableConfig{
					Header:  []schema.TableField{},
					RowData: [][]schema.TableField{},
					Toggled: false,
				},
			},
		},
		AppConfig: schema.AppConfig{
			Svg: make(map[string]string),
			// Width:      700,
			// Height:     500,
			// AsideWidth: 130,
			Width:      theme.DefaultTheme.AppSize[0],
			Height:     theme.DefaultTheme.AppSize[1],
			AsideWidth: theme.DefaultTheme.AsideMaxSize,

			Theme:    "gnome",
			GamePath: "",
		},
		DI: DIContainer{
			SQL:    nil,
			Vpk:    nil,
			Task:   nil,
			Window: nil,
		},
	}

	GLOBALAPP.setupSvg()
	GetSubCategories(GLOBALAPP.ComponentStatus.PageMods.SelectedategoryLabel[0])
	return GLOBALAPP
}

func (app *App) RegisterAllDependencies(w *gui.Window) {
	app.DI.Window = w
	app.DI.SQL = NewSqlite3(app)
	app.DI.Vpk = NewVpk(app)
	app.DI.Task = NewTask()
}

func (app *App) Prepare() {
	// chosed the path, query all mods from .db by local & workshop .vpk files
	if app.AppConfig.GamePath != "" {
		app.DI.Vpk.SyncVPKs(app.AppConfig.GamePath)
		if !app.ComponentStatus.PageTools.DeleteCollection.Toggled {
			app.DI.Vpk.FetchCollection()
		}
	}
}

func (app *App) SetTheme(theme gui.Theme) {
	gui.SetTheme(theme)
}

func (app *App) setupSvg() {
	entries, err := app.assetfile.ReadDir("assets/icons")
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".svg") {
			continue
		}

		filename := strings.TrimSuffix(entry.Name(), ".svg")
		if data, err := app.assetfile.ReadFile(path.Join("assets/icons", entry.Name())); err == nil {
			app.AppConfig.Svg[filename] = string(data)
		}
	}
}

func (app *App) GetSvg(key string) string {
	v, ok := app.AppConfig.Svg[key]
	if ok {
		return v
	}

	return ""
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

	t := make([]gui.SelectOption, 0, modLen/app.ComponentStatus.PageMods.PageAmount)

	for i := app.ComponentStatus.PageMods.PageAmount; i <= modLen; i += app.ComponentStatus.PageMods.PageAmount {
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

// Return the start and end mod indices for the selected dropdown option.
func (app *App) CurrentModsRange() (int, int) {
	start := 0
	end, err := strconv.Atoi(app.ComponentStatus.PageMods.PageEnd[0])
	if err != nil {
		return 0, 0
	}

	for i := 0; i < end-1; i += app.ComponentStatus.PageMods.PageAmount {
		start = i
	}

	return start, end
}

func (app *App) FetchDeleteCollection() {
	dc := &app.ComponentStatus.PageTools.DeleteCollection

	dc.Header = dc.Header[:0]

	clear(dc.RowData)
	dc.RowData = dc.RowData[:0]

	clear(dc.Selected)
	dc.Selected = dc.Selected[:0]

	dc.Toggled = false

	clear(app.DI.Vpk.Collections)
	app.DI.Vpk.Collections = app.DI.Vpk.Collections[:0]

	app.DI.Vpk.FetchCollection()
}

// GetSelectedCollectionsID return []string that picked in delete collection
func (app *App) GetSelectedCollectionsID() []string {
	dc := app.ComponentStatus.PageTools.DeleteCollection

	IDs := []string{}

	for idx, el := range dc.Selected {
		if el {
			IDs = append(IDs, app.DI.Vpk.Collections[idx].ID)
		}
	}
	return IDs
}

// GetSelectedCollections return []schema.Collection that picked in delete collection page
func (app *App) GetSelectedCollections() []schema.Collection {
	dc := app.ComponentStatus.PageTools.DeleteCollection

	colls := []schema.Collection{}

	for idx, el := range dc.Selected {
		if el {
			colls = append(colls, app.DI.Vpk.Collections[idx])
		}
	}
	return colls
}

func SwitchSubPage(idx, subPage int) {
	if idx > len(GLOBALAPP.ComponentStatus.SubAsideIndices) {
		return
	}

	GLOBALAPP.ComponentStatus.SubAsideIndices[idx] = subPage
}

func GetValue[T any](m map[string]any, key string) (T, bool) {
	v, ok := m[key].(T)
	if ok {
		return v, true
	}
	return v, false
}

func GetValueFromMap(m map[string]any, key string, notFoundValue string) string {
	v, ok := m[key].(string)
	if ok {
		return v
	}
	return notFoundValue
}
