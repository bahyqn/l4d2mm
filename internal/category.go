package internal

import "github.com/go-gui-org/go-gui/gui"

var ModLables = map[string][]gui.SelectOption{
	"All":         {},
	"Maps":        {},
	"Collections": {},
	"Rifles": {
		{
			Label: "M16",
			Value: "M16",
		},
		{
			Label: "Scar",
			Value: "Scar",
		},
		{
			Label: "AK47",
			Value: "AK47",
		},
		{
			Label: "SG552",
			Value: "v",
		},
		{
			Label: "M60",
			Value: "M60",
		},
	},
	"Shotguns": {
		{
			Label: "Punmp",
			Value: "Punmp",
		},
		{
			Label: "Chrome",
			Value: "Chrome",
		},
		{
			Label: "Auto",
			Value: "Auto",
		},
		{
			Label: "Spas",
			Value: "Spas",
		},
		{
			Label: "",
			Value: "",
		},
	},
	"Snipers": {
		{
			Label: "Hunting",
			Value: "Hunting",
		},
		{
			Label: "Military",
			Value: "Military",
		},
		{
			Label: "Scount",
			Value: "Scount",
		},
		{
			Label: "AWP",
			Value: "AWP",
		},
	},
	"SMG": {
		{
			Label: "SMG",
			Value: "SMG",
		},
		{
			Label: "Silenced",
			Value: "Silenced",
		},
		{
			Label: "MP5",
			Value: "MP5",
		},
		{
			Label: "",
			Value: "",
		},
	},
	"Pistols": {
		{
			Label: "Pistol",
			Value: "Pistol",
		},
		{
			Label: "Magnum",
			Value: "Magnum",
		},
	},
	"GL": {},
	"Melee": {
		{
			Label: "ChainSaw",
			Value: "ChainSaw",
		},
		{
			Label: "Fireaxe",
			Value: "Fireaxe",
		},
		{
			Label: "HuntingKnkife",
			Value: "HuntingKnkife",
		},
		{
			Label: "Katana",
			Value: "Katana",
		},
		{
			Label: "Cricket Bat",
			Value: "Cricket Bat",
		},
		{
			Label: "Baseball Bat",
			Value: "Baseball Bat",
		},
		{
			Label: "Golfclub",
			Value: "Golfclub",
		},
		{
			Label: "Machete",
			Value: "Machete",
		},
		{
			Label: "Tonfa",
			Value: "Tonfa",
		},
		{
			Label: "Electric Guitar",
			Value: "Electric Guitar",
		},
		{
			Label: "Frying Pan",
			Value: "Frying Pan",
		},
		{
			Label: "Crowbar",
			Value: "Crowbar",
		},
		{
			Label: "Riotshield",
			Value: "Riotshield",
		},
	},
	"Items": {
		{
			Label: "First aid kit",
			Value: "First aid kit",
		},
		{
			Label: "Defibrillator",
			Value: "Defibrillator",
		},
		{
			Label: "Addrenaline",
			Value: "Addrenaline",
		},
		{
			Label: "Pain pills",
			Value: "Pain pills",
		},
		{
			Label: "Molotov",
			Value: "Molotov",
		},
		{
			Label: "Pipebomb",
			Value: "Pipebomb",
		},
		{
			Label: "Vomit Jar",
			Value: "Vomit Jar",
		},
		{
			Label: "Gnome",
			Value: "Gnome",
		},
		{
			Label: "Cola Bottles",
			Value: "Cola Bottles",
		},
	},
	"Survivors": {
		{
			Label: "Coach",
			Value: "Coach",
		},
		{
			Label: "Ellis",
			Value: "Ellis",
		},
		{
			Label: "Nick",
			Value: "Nick",
		},
		{
			Label: "Rochelle",
			Value: "Rochelle",
		},
		{
			Label: "Bill",
			Value: "Bill",
		},
		{
			Label: "Francis",
			Value: "Francis",
		},
		{
			Label: "Louis",
			Value: "Louis",
		},
		{
			Label: "Zoey",
			Value: "Zoey",
		},
	},
	"Zombies": {
		{
			Label: "Boomer",
			Value: "Boomer",
		},
		{
			Label: "Charger",
			Value: "Charger",
		},
		{
			Label: "Hunter",
			Value: "Hunter",
		},
		{
			Label: "Jockey",
			Value: "Jockey",
		},
		{
			Label: "Smoker",
			Value: "Smoker",
		},
		{
			Label: "Spitter",
			Value: "Spitter",
		},
		{
			Label: "Tank",
			Value: "Tank",
		},
		{
			Label: "Witch",
			Value: "Witch",
		},
		{
			Label: "Zombie",
			Value: "Zombie",
		},
	},
	"Scripts": {},
	"Sounds":  {},
	"Effects": {},
}

type SourceMode string

const (
	SourceDefault       SourceMode = "Default"
	SourceWorkshopFirst SourceMode = "Workshop First"
	SourceLocalFirst    SourceMode = "Local First"
	SourceWorkshopOnly  SourceMode = "Workshop only"
	SourceLocalOnly     SourceMode = "Local only"
)

var AllSourceModes = []gui.SelectOption{
	{
		Label: string(SourceDefault),
		Value: string(SourceDefault),
	},
	{
		Label: string(SourceWorkshopFirst),
		Value: string(SourceWorkshopFirst),
	},
	{
		Label: string(SourceLocalFirst),
		Value: string(SourceLocalFirst),
	},
	{
		Label: string(SourceWorkshopOnly),
		Value: string(SourceWorkshopOnly),
	},
	{
		Label: string(SourceLocalOnly),
		Value: string(SourceLocalOnly),
	},
}
