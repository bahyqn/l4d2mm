package schema

import (
	"github.com/go-gui-org/go-gui/gui"
)

type ComponentStatus struct {
	AsideModExpand        bool
	AsideIdx              int
	ModsSearchValue       string
	Categories            []string // Labels under the `Mods` by dynamic generrate accoding to all of your installed mods
	SelecCtedategoryLabel string   // Such as: map, weapon, ..., you  can pick one
	SubLabels             []string // Dynamic generate. Such as: fireaxe, katana, ...
	SelectedSubLabel      []string // Such as: you pick the melee first, labels will disaplay: fireaxe, katana, ... (you just can pick one)

	PageSize        int
	ModsBySelectIdx []Mod
	PageEnd         []string
}

type TemplateSelect struct {
	ID           string
	Placeholder  string
	Selected     []string
	Options      []string
	OnSelectFunc func(s []string, e *gui.Event, w *gui.Window)
}

type AppConfig struct {
	Id         int
	Width      float32
	Height     float32
	AsideWidth float32

	Theme string
}

type Mod struct {
	Id             string `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	Category       string `gorm:"column:category"`
	SubType        string `gorm:"column:sub_type"`
	Url            string `gorm:"column:url"`
	Remark         string `gorm:"column:remark"`
	IsEnable       bool   `gorm:"column:is_enable"`
	IsRemoved      bool   `gorm:"column:is_removed"`
	IsFromWorkshop bool   `gorm:"column:is_from_workshop"` // true: workshop, false: local
	HasConflict    bool   `horm:"column:has_conflict"`
}

type Collection struct {
	Id        string   `gorm:"column:id"`
	Name      string   `gorm:"column:name"`
	Url       string   `gorm:"column:url"`
	Mods      []string `gorm:"column:mods;serializer:json"`
	Remark    string   `gorm:"column:remark"`
	IsEnable  bool     `gorm:"column:is_enable"`
	IsRemoved bool     `gorm:"column:is_removed"`
}
