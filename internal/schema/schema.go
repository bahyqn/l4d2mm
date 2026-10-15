package schema

import (
	"github.com/go-gui-org/go-gui/gui"
)

type ModLocation string

const (
	LocationGame ModLocation = "game"
	LocationApp  ModLocation = "app"
)

type Mod struct {
	Idx            int              `gorm:"-"`
	Id             string           `gorm:"primaryKey;column:id"`
	Filename       string           `gorm:"column:filename"`
	ModName        string           `gorm:"column:mod_name"`
	Author         string           `gorm:"column:author"`
	Category       string           `gorm:"column:category"`
	SubType        string           `gorm:"column:sub_type"`
	Url            string           `gorm:"column:url"`
	Remark         string           `gorm:"column:remark"`
	Location       ModLocation      `gorm:"column:location"`
	IsEnable       bool             `gorm:"column:is_enable"`
	IsRemoved      bool             `gorm:"column:is_removed"`
	IsFromWorkshop bool             `gorm:"column:is_from_workshop"`
	HasConflict    bool             `gorm:"column:has_conflict"`
	Addoninfo      map[string]any   `gorm:"column:addinfo;serializer:json"`
	Missions       []map[string]any `gorm:"column:missions;serializer:json"`
	VpkVersion     uint32           `gorm:"column:vpk_version"`
	Treesize       uint32           `gorm:"column:tree_size"`
}

func (mod *Mod) TableName() string {
	return "mods"
}

type Collection struct {
	Id             string   `gorm:"column:id"`
	Name           string   `gorm:"column:name"`
	Url            string   `gorm:"column:url"`
	Mods           []string `gorm:"column:mods;serializer:json"`
	Remark         string   `gorm:"column:remark"`
	IsEnable       bool     `gorm:"column:is_enable"`
	IsRemoved      bool     `gorm:"column:is_removed"`
	IsFromWorkshop bool     `gorm:"column:is_from_workshop"`
}

func (collection *Collection) TableName() string {
	return "collections"
}

type AppConfig struct {
	Id         int     `gorm:"column:id"`
	Width      float32 `gorm:"-"`
	Height     float32 `gorm:"-"`
	AsideWidth float32 `gorm:"-"`

	Theme string `gorm:"column:theme"`

	// automatic get value by FetchTables(sqlite.go)
	GamePath string `gorm:"column:game_path"`
}

func (cf *AppConfig) TableName() string {
	return "app_config"
}

// no gorm

type ComponentStatus struct {
	AsideIdx        int
	SubAsideIndices [6]int // 0: mods, 1: workshop, 2: tasks, 3: tools,  4: servers, 5: settings
	PageMods        PageMods
	PageTools       PageTools
}

type PageMods struct {
	AsideModExpand       bool
	ModsSearchValue      string
	Categories           []gui.SelectOption // Labels under the `Mods` by dynamic generrate accoding to all of your installed mods
	SelectedategoryLabel []string           // Such as: map, weapon, ..., you  can pick one
	SubLabels            []gui.SelectOption // Dynamic generate. Such as: fireaxe, katana, ...
	SelectedSubLabel     []string           // Such as: you pick the melee first, labels will disaplay: fireaxe, katana, ... (you just can pick one)
	DynamicPageSelect    []gui.SelectOption

	PageAmount      int
	ModsBySelectIdx []Mod
	PageEnd         []string
	SourceMode      string

	// details
	ModDetails           *Mod
	TempModDetailsFields [3]string
}

type PageTools struct {
	BreadCrumb []string
}

type TemplateSelect struct {
	ID           string
	MaxWidth     float32
	Placeholder  string
	Selected     []string
	Options      []gui.SelectOption
	Invisible    bool
	OnSelectFunc func([]string, gui.EventCtx)
}

type VpkInfo struct {
	Addoninfo  map[string]any
	Missions   []map[string]any
	VpkVersion uint32
	Treesize   uint32
}

type ToolCardInfo struct {
	SvgFilename string
	Title       string
	Describe    string
}
