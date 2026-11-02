package internal

import (
	"l4d2mm/internal/schema"

	"github.com/go-gui-org/go-gui/gui"
)

var ModsCategoris = []gui.SelectOption{}
var ModLables = map[string][]gui.SelectOption{}
var MDL = map[string]schema.VPKCategory{}

var LABELS = []schema.VPKLabel{
	{
		Label:    "All",
		SubLabel: []schema.VPKSubLabel{},
	},
	{
		Label:    "Collections",
		SubLabel: []schema.VPKSubLabel{},
	},
	{
		Label:    "Maps",
		SubLabel: []schema.VPKSubLabel{},
	},
	{
		Label: "Rifles",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "M16",
				Key:   "v_rifle_m16a2",
			},
			{
				Label: "M16",
				Key:   "w_rifle_m16a2",
			},
			{
				Label: "Scar",
				Key:   "v_desert_rifle",
			},
			{
				Label: "Scar",
				Key:   "w_desert_rifle",
			},
			{
				Label: "AK47",
				Key:   "v_rifle_ak47",
			},
			{
				Label: "AK47",
				Key:   "w_rifle_ak47",
			},
			{
				Label: "SG552",
				Key:   "v_rifle_sg552",
			},
			{
				Label: "SG552",
				Key:   "w_rifle_sg553",
			},
			{
				Label: "M60",
				Key:   "v_60",
			},
			{
				Label: "M60",
				Key:   "w_60",
			},
		},
	},
	{
		Label: "Shotguns",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "Pump",
				Key:   "v_pumpshotgun",
			},
			{
				Label: "Pump",
				Key:   "w_pumpshotgun",
			},
			{
				Label: "Chrome",
				Key:   "v_shotgun_chrome",
			},
			{
				Label: "Chrome",
				Key:   "w_shotgun_chrome",
			},
			{
				Label: "Auto",
				Key:   "v_autoshotgun",
			},
			{
				Label: "Auto",
				Key:   "w_autoshotgun",
			},
			{
				Label: "Spass",
				Key:   "v_shotgun_spas",
			},
			{
				Label: "Spass",
				Key:   "w_shotgun_spas",
			},
		},
	},
	{
		Label: "Snipers",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "Hunting",
				Key:   "v_huntingrifle",
			},
			{
				Label: "Hunting",
				Key:   "w_huntingrifle",
			},
			{
				Label: "Military",
				Key:   "v_sniper_military",
			},
			{
				Label: "Military",
				Key:   "w_sniper_military",
			},
			{
				Label: "Scout",
				Key:   "v_sniper_scout",
			},
			{
				Label: "Scout",
				Key:   "w_sniper_scout",
			},
			{
				Label: "AWP",
				Key:   "v_sniper_awp",
			},
			{
				Label: "AWP",
				Key:   "w_sniper_awp",
			},
		},
	},
	{
		Label: "SMG",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "SMG(Uzi)",
				Key:   "v_smg_uzi",
			},
			{
				Label: "SMG(Uzi)",
				Key:   "w_smg_uzi",
			},
			{
				Label: "Silenced (MAC-10)",
				Key:   "v_smg_a",
			},
			{
				Label: "Silenced (MAC-10)",
				Key:   "w_smg_a",
			},
			{
				Label: "MP5",
				Key:   "v_smg_mp5",
			},
			{
				Label: "MP5",
				Key:   "w_smg_mp5",
			},
		},
	},
	{
		Label: "Pistols",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "Pistol",
				Key:   "v_pistola",
			},
			{
				Label: "Pistol",
				Key:   "w_pistola",
			},
			{
				Label: "Magnum",
				Key:   "v_desert_eagle",
			},
			{
				Label: "Magnum",
				Key:   "w_desert_eagle",
			},
		},
	},
	{
		Label: "GL",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "",
				Key:   "v_grenadelaunche",
			},
			{
				Label: "",
				Key:   "W_grenadelaunche",
			},
		},
	},
	{
		Label: "Melee",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "ChainSaw",
				Key:   "v_chainsaw",
			},
			{
				Label: "ChainSaw",
				Key:   "w_chainsaw",
			},
			{
				Label: "Fireaxe",
				Key:   "v_fireaxe",
			},
			{
				Label: "Fireaxe",
				Key:   "w_fireaxe",
			},
			{
				Label: "HuntingKnife",
				Key:   "v_hunting_knife",
			},
			{
				Label: "HuntingKnife",
				Key:   "w_hunting_knife",
			},
			{
				Label: "Katana",
				Key:   "v_katana",
			},
			{
				Label: "Katana",
				Key:   "w_katana",
			},
			{
				Label: "Cricket Bat",
				Key:   "v_cricket_bat",
			},
			{
				Label: "Cricket Bat",
				Key:   "w_cricket_bat",
			},
			{
				Label: "Bat",
				Key:   "v_bat",
			},
			{
				Label: "Bat",
				Key:   "w_bat",
			},
			{
				Label: "Golfclub",
				Key:   "v_Golfclub",
			},
			{
				Label: "Golfclub",
				Key:   "w_Golfclub",
			},
			{
				Label: "Machete",
				Key:   "v_machete",
			},
			{
				Label: "Machete",
				Key:   "w_machete",
			},
			{
				Label: "Tonfa",
				Key:   "v_tonfa",
			},
			{
				Label: "Tonfa",
				Key:   "w_tonfa",
			},
			{
				Label: "Electric Guitar",
				Key:   "v_electric_guitar",
			},
			{
				Label: "Electric Guitar",
				Key:   "w_electric_guitar",
			},
			{
				Label: "Frying Pan",
				Key:   "v_frying_pan",
			},
			{
				Label: "Frying Pan",
				Key:   "w_frying_pan",
			},
			{
				Label: "Crowbar",
				Key:   "v_crowbar",
			},
			{
				Label: "Crowbar",
				Key:   "w_crowbar",
			},
			{
				Label: "Riotshield",
				Key:   "v_riotshield",
			},
			{
				Label: "Riotshield",
				Key:   "w_riotshield",
			},
		},
	},
	{
		Label: "Items",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "First aid kit",
				Key:   "v_first_aid_kit",
			},
			{
				Label: "First aid kit",
				Key:   "w_first_aid_kit",
			},
			{
				Label: "Defibrillator",
				Key:   "v_defibrillator",
			},
			{
				Label: "Defibrillator",
				Key:   "w_defibrillator",
			},
			{
				Label: "Adrenaline",
				Key:   "v_adrenaline",
			},
			{
				Label: "Adrenaline",
				Key:   "w_adrenaline",
			},
			{
				Label: "Pain pills",
				Key:   "v_pills",
			},
			{
				Label: "Pain pills",
				Key:   "w_pills",
			},
			{
				Label: "Molotov",
				Key:   "v_molotov",
			},
			{
				Label: "Molotov",
				Key:   "w_molotov",
			},
			{
				Label: "Pipebomb",
				Key:   "v_pipebomb",
			},
			{
				Label: "Pipebomb",
				Key:   "w_pipebomb",
			},
			{
				Label: "Vomit Jar",
				Key:   "v_bile_flask",
			},
			{
				Label: "Vomit Jar",
				Key:   "w_bile_flask",
			},
			{
				Label: "Gnome",
				Key:   "v_gnome",
			},
			{
				Label: "Gnome",
				Key:   "w_gnome",
			},
			{
				Label: "Cola Bottles",
				Key:   "v_cola",
			},
			{
				Label: "Cola Bottles",
				Key:   "w_cola",
			},
		},
	},
	{
		Label: "Scripts",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "Coach",
				Key:   "survivor_coach",
			},
			{
				Label: "Coach",
				Key:   "v_arms_coach",
			},
			{
				Label: "Nick",
				Key:   "survivor_mechanic",
			},
			{
				Label: "Nick",
				Key:   "v_arms-v_arms_mechanic",
			},
			{
				Label: "Ellis",
				Key:   "survivor_gambler",
			},
			{
				Label: "Ellis",
				Key:   "survivor_gambler",
			},
			{
				Label: "Rochelle",
				Key:   "survivor_producer",
			},
			{
				Label: "Rochelle",
				Key:   "v_arms_producer",
			},
			{
				Label: "Bill",
				Key:   "survivor_namvet",
			},
			{
				Label: "Bill",
				Key:   "v_arms_namvet",
			},
			{
				Label: "Francis",
				Key:   "survivor_biker",
			},
			{
				Label: "Francis",
				Key:   "v_arms_biker",
			},
			{
				Label: "Louis",
				Key:   "survivor_manager",
			},
			{
				Label: "Louis",
				Key:   "v_arms_manager",
			},
			{
				Label: "Zoey",
				Key:   "survivor_teenangst",
			},
			{
				Label: "Zoey",
				Key:   "v_arms_teenangst",
			},
		},
	},
	{
		Label: "Zombies",
		SubLabel: []schema.VPKSubLabel{
			{
				Label: "Boomer",
				Key:   "boomer",
			},
			{
				Label: "Boomer",
				Key:   "boomette",
			},
			{
				Label: "Charger",
				Key:   "charger",
			},
			{
				Label: "Jockey",
				Key:   "jockey",
			},
			{
				Label: "Smoker",
				Key:   "smoker",
			},
			{
				Label: "Spitter",
				Key:   "spitter",
			},
			{
				Label: "Tank",
				Key:   "hulk",
			},
			{
				Label: "Tank",
				Key:   "hulk_dlc3",
			},
			{
				Label: "Witch",
				Key:   "witch",
			},
			{
				Label: "Witch",
				Key:   "witch_bride",
			},
			{
				Label: "Zombie",
				Key:   "common_*",
			},
		},
	},
	{
		Label:    "Sounds",
		SubLabel: []schema.VPKSubLabel{},
	},
	{
		Label:    "Effects",
		SubLabel: []schema.VPKSubLabel{},
	},
}

func PrepareLabels() {

	key := ""

	for _, item := range LABELS {
		ModsCategoris = append(ModsCategoris, gui.SelectOption{
			Label: item.Label,
			Value: item.Label,
		})

		if len(item.SubLabel) > 0 {
			for _, sub := range item.SubLabel {
				if key != sub.Label {
					ModLables[item.Label] = append(ModLables[item.Label], gui.SelectOption{
						Label: sub.Label,
						Value: sub.Label,
					})
				}
				key = sub.Label

				MDL[sub.Key] = schema.VPKCategory{
					Category: item.Label,
					SubLabel: sub.Label,
				}
			}
		} else {
			ModLables[item.Label] = []gui.SelectOption{}
		}
	}
}

// var ModsCategoris = []gui.SelectOption{
// 	{
// 		Label: "All",
// 		Value: "All",
// 	},
// 	{
// 		Label: "Collections",
// 		Value: "Collections",
// 	},
// 	{
// 		Label: "Maps",
// 		Value: "Maps",
// 	},
// 	{
// 		Label: "Rifles",
// 		Value: "Rifles",
// 	},
// 	{
// 		Label: "Shotguns",
// 		Value: "Shotguns",
// 	},
// 	{
// 		Label: "Snipers",
// 		Value: "Snipers",
// 	},
// 	{
// 		Label: "SMG",
// 		Value: "SMG",
// 	},
// 	{
// 		Label: "Pistols",
// 		Value: "Pistols",
// 	},
// 	{
// 		Label: "GL",
// 		Value: "GL",
// 	},
// 	{
// 		Label: "Melee",
// 		Value: "Melee",
// 	},
// 	{
// 		Label: "Items",
// 		Value: "Items",
// 	},
// 	{
// 		Label: "Scripts",
// 		Value: "Scripts",
// 	},
// 	{
// 		Label: "Sounds",
// 		Value: "Sounds",
// 	},
// 	{
// 		Label: "Effects",
// 		Value: "Effects",
// 	},
// }

// var ModLables = map[string][]gui.SelectOption{
// 	"All":         {},
// 	"Maps":        {},
// 	"Collections": {},
// 	"Rifles": {
// 		{
// 			Label: "M16",
// 			Value: "M16",
// 		},
// 		{
// 			Label: "Scar",
// 			Value: "Scar",
// 		},
// 		{
// 			Label: "AK47",
// 			Value: "AK47",
// 		},
// 		{
// 			Label: "SG552",
// 			Value: "v",
// 		},
// 		{
// 			Label: "M60",
// 			Value: "M60",
// 		},
// 	},
// 	"Shotguns": {
// 		{
// 			Label: "Punmp",
// 			Value: "Punmp",
// 		},
// 		{
// 			Label: "Chrome",
// 			Value: "Chrome",
// 		},
// 		{
// 			Label: "Auto",
// 			Value: "Auto",
// 		},
// 		{
// 			Label: "Spas",
// 			Value: "Spas",
// 		},
// 		{
// 			Label: "",
// 			Value: "",
// 		},
// 	},
// 	"Snipers": {
// 		{
// 			Label: "Hunting",
// 			Value: "Hunting",
// 		},
// 		{
// 			Label: "c",
// 			Value: "Military",
// 		},
// 		{
// 			Label: "Scount",
// 			Value: "Scount",
// 		},
// 		{
// 			Label: "AWP",
// 			Value: "AWP",
// 		},
// 	},
// 	"SMG": {
// 		{
// 			Label: "SMG",
// 			Value: "SMG",
// 		},
// 		{
// 			Label: "Silenced",
// 			Value: "Silenced",
// 		},
// 		{
// 			Label: "MP5",
// 			Value: "MP5",
// 		},
// 		{
// 			Label: "",
// 			Value: "",
// 		},
// 	},
// 	"Pistols": {
// 		{
// 			Label: "Pistol",
// 			Value: "Pistol",
// 		},
// 		{
// 			Label: "Magnum",
// 			Value: "Magnum",
// 		},
// 	},
// 	"GL": {},
// 	"Melee": {
// 		{
// 			Label: "ChainSaw",
// 			Value: "ChainSaw",
// 		},
// 		{
// 			Label: "Fireaxe",
// 			Value: "Fireaxe",
// 		},
// 		{
// 			Label: "HuntingKnkife",
// 			Value: "HuntingKnkife",
// 		},
// 		{
// 			Label: "Katana",
// 			Value: "Katana",
// 		},
// 		{
// 			Label: "Cricket Bat",
// 			Value: "Cricket Bat",
// 		},
// 		{
// 			Label: "Baseball Bat",
// 			Value: "Baseball Bat",
// 		},
// 		{
// 			Label: "Golfclub",
// 			Value: "Golfclub",
// 		},
// 		{
// 			Label: "Machete",
// 			Value: "Machete",
// 		},
// 		{
// 			Label: "Tonfa",
// 			Value: "Tonfa",
// 		},
// 		{
// 			Label: "Electric Guitar",
// 			Value: "Electric Guitar",
// 		},
// 		{
// 			Label: "Frying Pan",
// 			Value: "Frying Pan",
// 		},
// 		{
// 			Label: "Crowbar",
// 			Value: "Crowbar",
// 		},
// 		{
// 			Label: "Riotshield",
// 			Value: "Riotshield",
// 		},
// 	},
// 	"Items": {
// 		{
// 			Label: "First aid kit",
// 			Value: "First aid kit",
// 		},
// 		{
// 			Label: "Defibrillator",
// 			Value: "Defibrillator",
// 		},
// 		{
// 			Label: "Addrenaline",
// 			Value: "Addrenaline",
// 		},
// 		{
// 			Label: "Pain pills",
// 			Value: "Pain pills",
// 		},
// 		{
// 			Label: "Molotov",
// 			Value: "Molotov",
// 		},
// 		{
// 			Label: "Pipebomb",
// 			Value: "Pipebomb",
// 		},
// 		{
// 			Label: "Vomit Jar",
// 			Value: "Vomit Jar",
// 		},
// 		{
// 			Label: "Gnome",
// 			Value: "Gnome",
// 		},
// 		{
// 			Label: "Cola Bottles",
// 			Value: "Cola Bottles",
// 		},
// 	},
// 	"Survivors": {
// 		{
// 			Label: "Coach",
// 			Value: "Coach",
// 		},
// 		{
// 			Label: "Ellis",
// 			Value: "Ellis",
// 		},
// 		{
// 			Label: "Nick",
// 			Value: "Nick",
// 		},
// 		{
// 			Label: "Rochelle",
// 			Value: "Rochelle",
// 		},
// 		{
// 			Label: "Bill",
// 			Value: "Bill",
// 		},
// 		{
// 			Label: "Francis",
// 			Value: "Francis",
// 		},
// 		{
// 			Label: "Louis",
// 			Value: "Louis",
// 		},
// 		{
// 			Label: "Zoey",
// 			Value: "Zoey",
// 		},
// 	},
// 	"Zombies": {
// 		{
// 			Label: "Boomer",
// 			Value: "Boomer",
// 		},
// 		{
// 			Label: "Charger",
// 			Value: "Charger",
// 		},
// 		{
// 			Label: "Hunter",
// 			Value: "Hunter",
// 		},
// 		{
// 			Label: "Jockey",
// 			Value: "Jockey",
// 		},
// 		{
// 			Label: "Smoker",
// 			Value: "Smoker",
// 		},
// 		{
// 			Label: "Spitter",
// 			Value: "Spitter",
// 		},
// 		{
// 			Label: "Tank",
// 			Value: "Tank",
// 		},
// 		{
// 			Label: "Witch",
// 			Value: "Witch",
// 		},
// 		{
// 			Label: "Zombie",
// 			Value: "Zombie",
// 		},
// 	},
// 	"Scripts": {},
// 	"Sounds":  {},
// 	"Effects": {},
// }

type SourceMode string

const (
	SourceDefault       SourceMode = "Default"
	SourceWorkshopFirst SourceMode = "Workshop First"
	SourceLocalFirst    SourceMode = "Local First"
	SourceEnableFirst   SourceMode = "Enable First"
	SourceDisableFirst  SourceMode = "Disable First"
	SourceWorkshopOnly  SourceMode = "Workshop Only"
	SourceLocalOnly     SourceMode = "Local Only"
	SourceEnableOnly    SourceMode = "Enable Only"
	SourceDisableOnly   SourceMode = "Disable Only"
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
		Label: string(SourceEnableFirst),
		Value: string(SourceEnableFirst),
	},
	{
		Label: string(SourceDisableFirst),
		Value: string(SourceDisableFirst),
	},
	{
		Label: string(SourceWorkshopOnly),
		Value: string(SourceWorkshopOnly),
	},
	{
		Label: string(SourceLocalOnly),
		Value: string(SourceLocalOnly),
	},
	{
		Label: string(SourceEnableOnly),
		Value: string(SourceEnableOnly),
	},
	{
		Label: string(SourceDisableOnly),
		Value: string(SourceDisableOnly),
	},
}

// var MDL = map[string]schema.VPKCategory{
// 	// rifles
// 	"v_rifle_m16a2": {
// 		Category: "Rifles",
// 		SubLabel: "M16",
// 	},
// 	"w_rifle_m16a2": {
// 		Category: "Rifles",
// 		SubLabel: "M16",
// 	},
// 	"v_desert_rifle": {
// 		Category: "Rifles",
// 		SubLabel: "Scar",
// 	},
// 	"w_desert_rifle": {
// 		Category: "Rifles",
// 		SubLabel: "Scar",
// 	},
// 	"v_rifle_ak47": {
// 		Category: "Rifles",
// 		SubLabel: "AK47",
// 	},
// 	"w_rifle_ak47": {
// 		Category: "Rifles",
// 		SubLabel: "AK47",
// 	},
// 	"v_rifle_sg552": {
// 		Category: "Rifles",
// 		SubLabel: "SG552",
// 	},
// 	"w_rifle_sg553": {
// 		Category: "Rifles",
// 		SubLabel: "SG552",
// 	},
// 	"v_60": {
// 		Category: "Rifles",
// 		SubLabel: "M60",
// 	},
// 	"w_60": {
// 		Category: "Rifles",
// 		SubLabel: "M60",
// 	},
// 	// shotgun
// 	"v_pumpshotgun": {
// 		Category: "Shotguns",
// 		SubLabel: "Pump",
// 	},
// 	"w_pumpshotgun": {
// 		Category: "Shotguns",
// 		SubLabel: "Pump",
// 	},
// 	"v_shotgun_chrome": {
// 		Category: "Shotguns",
// 		SubLabel: "Chrome",
// 	},
// 	"w_shotgun_chrome": {
// 		Category: "Shotguns",
// 		SubLabel: "Chrome",
// 	},
// 	"v_autoshotgun": {
// 		Category: "Shotguns",
// 		SubLabel: "Auto",
// 	},
// 	"w_autoshotgun": {
// 		Category: "Shotguns",
// 		SubLabel: "Auto",
// 	},
// 	"v_shotgun_spas": {
// 		Category: "Shotguns",
// 		SubLabel: "Auto",
// 	},
// 	"w_shotgun_spas": {
// 		Category: "Shotguns",
// 		SubLabel: "Auto",
// 	},

// 	// sniper
// 	"v_huntingrifle": {
// 		Category: "Snipers",
// 		SubLabel: "Hunting",
// 	},
// 	"w_huntingrifle": {
// 		Category: "Snipers",
// 		SubLabel: "Hunting",
// 	},
// 	"v_sniper_military": {
// 		Category: "Snipers",
// 		SubLabel: "Military",
// 	},
// 	"w_sniper_military": {
// 		Category: "Snipers",
// 		SubLabel: "Military",
// 	},
// 	"v_sniper_scout": {
// 		Category: "Snipers",
// 		SubLabel: "Scout",
// 	},
// 	"w_sniper_scout": {
// 		Category: "Snipers",
// 		SubLabel: "Scout",
// 	},
// 	"v_sniper_awp": {
// 		Category: "Snipers",
// 		SubLabel: "AWP",
// 	},
// 	"w_sniper_awp": {
// 		Category: "Snipers",
// 		SubLabel: "AWP",
// 	},

// 	// SMG
// 	"v_smg_uzi": {
// 		Category: "SMG",
// 		SubLabel: "SMG(Uzi)",
// 	},
// 	"w_smg_uzi": {
// 		Category: "SMG",
// 		SubLabel: "SMG(Uzi)",
// 	},
// 	"v_smg_a": {
// 		Category: "SMG",
// 		SubLabel: "Silenced (MAC-10)",
// 	},
// 	"w_smg_a": {
// 		Category: "SMG",
// 		SubLabel: "Silenced (MAC-10)",
// 	},
// 	"v_smg_mp5": {
// 		Category: "SMG",
// 		SubLabel: "MP5",
// 	},
// 	"w_smg_mp5": {
// 		Category: "SMG",
// 		SubLabel: "MP5",
// 	},

// 	// Pistols
// 	"v_pistola": {
// 		Category: "Pistols",
// 		SubLabel: "Pistol",
// 	},
// 	"w_pistol_a": {
// 		Category: "Pistols",
// 		SubLabel: "Pistol",
// 	},
// 	"v_dual_pistola": {
// 		Category: "Pistols",
// 		SubLabel: "Pistol",
// 	},
// 	"v_desert_eagle": {
// 		Category: "Pistols",
// 		SubLabel: "Magnum",
// 	},
// 	"w_desert_eagle": {
// 		Category: "Pistols",
// 		SubLabel: "Magnum",
// 	},

// 	// GL
// 	"v_grenadelaunche": {
// 		Category: "GL",
// 		SubLabel: "",
// 	},
// 	"w_grenadelaunche": {
// 		Category: "GL",
// 		SubLabel: "",
// 	},

// 	// Melee
// 	"v_chainsaw": {
// 		Category: "Melee",
// 		SubLabel: "ChainSaw",
// 	},
// 	"w_chainsaw": {
// 		Category: "Melee",
// 		SubLabel: "ChainSaw",
// 	},
// 	"v_fireaxe": {
// 		Category: "Melee",
// 		SubLabel: "Fireaxe",
// 	},
// 	"w_fireaxe": {
// 		Category: "Melee",
// 		SubLabel: "Fireaxe",
// 	},
// 	"v_hunting_knife": {
// 		Category: "Melee",
// 		SubLabel: "HuntingKnife",
// 	},
// 	"w_hunting_knife": {
// 		Category: "Melee",
// 		SubLabel: "HuntingKnife",
// 	},
// 	"v_katana": {
// 		Category: "Melee",
// 		SubLabel: "Katana",
// 	},
// 	"w_katana": {
// 		Category: "Melee",
// 		SubLabel: "Katana",
// 	},
// 	"v_cricket_bat": {
// 		Category: "Melee",
// 		SubLabel: "Cricket Bat",
// 	},
// 	"w_cricket_bat": {
// 		Category: "Melee",
// 		SubLabel: "Cricket Bat",
// 	},
// 	"v_bat": {
// 		Category: "Melee",
// 		SubLabel: "Bat",
// 	},
// 	"w_bat": {
// 		Category: "Melee",
// 		SubLabel: "Bat",
// 	},
// 	"v_golfclub": {
// 		Category: "Melee",
// 		SubLabel: "Golfclub",
// 	},
// 	"w_golfclub": {
// 		Category: "Melee",
// 		SubLabel: "Golfclub",
// 	},
// 	"v_machete": {
// 		Category: "Melee",
// 		SubLabel: "Machete",
// 	},
// 	"w_machete": {
// 		Category: "Melee",
// 		SubLabel: "Machete",
// 	},
// 	"v_tonfa": {
// 		Category: "Melee",
// 		SubLabel: "Tonfa",
// 	},
// 	"w_tonfa": {
// 		Category: "Melee",
// 		SubLabel: "Tonfa",
// 	},
// 	"v_electric_guitar": {
// 		Category: "Melee",
// 		SubLabel: "Electric Guitar",
// 	},
// 	"w_electric_guitar": {
// 		Category: "Melee",
// 		SubLabel: "Electric Guitar",
// 	},
// 	"v_frying_pan": {
// 		Category: "Melee",
// 		SubLabel: "Frying Pan",
// 	},
// 	"w_frying_pan": {
// 		Category: "Melee",
// 		SubLabel: "Frying Pan",
// 	},
// 	"v_crowbar": {
// 		Category: "Melee",
// 		SubLabel: "Crowbar",
// 	},
// 	"w_crowbar": {
// 		Category: "Melee",
// 		SubLabel: "Crowbar",
// 	},
// 	"v_riotshield": {
// 		Category: "Melee",
// 		SubLabel: "Riotshield",
// 	},
// 	"w_riotshield": {
// 		Category: "Melee",
// 		SubLabel: "Riotshield",
// 	},

// 	// Items
// 	"v_first_aid_kit": {
// 		Category: "Items",
// 		SubLabel: "First aid kit",
// 	},
// 	"w_first_aid_kit": {
// 		Category: "Items",
// 		SubLabel: "First aid kit",
// 	},
// 	"v_defibrillator": {
// 		Category: "Items",
// 		SubLabel: "Defibrillator",
// 	},
// 	"w_defibrillator": {
// 		Category: "Items",
// 		SubLabel: "Defibrillator",
// 	},
// 	"v_adrenaline": {
// 		Category: "Items",
// 		SubLabel: "Adrenaline",
// 	},
// 	"w_adrenaline": {
// 		Category: "Items",
// 		SubLabel: "Adrenaline",
// 	},
// 	"v_pills": {
// 		Category: "Items",
// 		SubLabel: "Pain pills",
// 	},
// 	"w_pills": {
// 		Category: "Items",
// 		SubLabel: "Pain pills",
// 	},
// 	"v_molotov": {
// 		Category: "Items",
// 		SubLabel: "Molotov",
// 	},
// 	"w_molotov": {
// 		Category: "Items",
// 		SubLabel: "Molotov",
// 	},
// 	"v_pipebomb": {
// 		Category: "Items",
// 		SubLabel: "Pipebomb",
// 	},
// 	"w_pipebomb": {
// 		Category: "Items",
// 		SubLabel: "Pipebomb",
// 	},
// 	"v_bile_flask": {
// 		Category: "Items",
// 		SubLabel: "Vomit Jar",
// 	},
// 	"w_bile_flask": {
// 		Category: "Items",
// 		SubLabel: "Vomit Jar",
// 	},
// 	"v_gnome": {
// 		Category: "Items",
// 		SubLabel: "Gnome",
// 	},
// 	"w_gnome": {
// 		Category: "Items",
// 		SubLabel: "Gnome",
// 	},
// 	"v_cola": {
// 		Category: "Items",
// 		SubLabel: "Cola Bottles",
// 	},
// 	"w_cola": {
// 		Category: "Items",
// 		SubLabel: "Cola Bottles",
// 	},

// 	// Survivors
// 	"survivor_coach": {
// 		Category: "Survivors",
// 		SubLabel: "Coach",
// 	},
// 	"v_arms_coach": {
// 		Category: "Survivors",
// 		SubLabel: "Coach",
// 	},
// 	"survivor_mechanic": {
// 		Category: "Survivors",
// 		SubLabel: "Nick",
// 	},
// 	"v_arms__mechanic": {
// 		Category: "Survivors",
// 		SubLabel: "Nick",
// 	},
// 	"survivor_gambler": {
// 		Category: "Survivors",
// 		SubLabel: "Ellis",
// 	},
// 	"v_arms_gambler": {
// 		Category: "Survivors",
// 		SubLabel: "Nick",
// 	},
// 	"survivor_producer": {
// 		Category: "Survivors",
// 		SubLabel: "Rochelle",
// 	},
// 	"v_arms_producer": {
// 		Category: "Survivors",
// 		SubLabel: "Rochelle",
// 	},
// 	"survivor_namvet": {
// 		Category: "Survivors",
// 		SubLabel: "Bill",
// 	},
// 	"v_arms_namvet": {
// 		Category: "Survivors",
// 		SubLabel: "Bill",
// 	},
// 	"survivor_biker": {
// 		Category: "Survivors",
// 		SubLabel: "Francis",
// 	},
// 	"v_arms_biker": {
// 		Category: "Survivors",
// 		SubLabel: "Francis",
// 	},
// 	"survivor_manager": {
// 		Category: "Survivors",
// 		SubLabel: "Louis",
// 	},
// 	"v_aaarms_manager": {
// 		Category: "Survivors",
// 		SubLabel: "Louis",
// 	},
// 	"survivor_teenangst": {
// 		Category: "Survivors",
// 		SubLabel: "Zoey",
// 	},
// 	"v_arms_teenangst": {
// 		Category: "Survivors",
// 		SubLabel: "Zoey",
// 	},

// 	// Zombies
// 	"boomer": {
// 		Category: "Zombies",
// 		SubLabel: "Boomer",
// 	},
// 	"boomette": {
// 		Category: "Zombies",
// 		SubLabel: "Boomer",
// 	},
// 	"charger": {
// 		Category: "Zombies",
// 		SubLabel: "Charger",
// 	},
// 	"hunter": {
// 		Category: "Zombies",
// 		SubLabel: "Hunter",
// 	},
// 	"jockey": {
// 		Category: "Zombies",
// 		SubLabel: "Jockey",
// 	},
// 	"smoker": {
// 		Category: "Zombies",
// 		SubLabel: "Smoker",
// 	},
// 	"spitter": {
// 		Category: "Zombies",
// 		SubLabel: "Spitter",
// 	},
// 	"hulk": {
// 		Category: "Zombies",
// 		SubLabel: "Tank",
// 	},
// 	"hulk_dlc3": {
// 		Category: "Zombies",
// 		SubLabel: "Tank",
// 	},
// 	"witch": {
// 		Category: "Zombies",
// 		SubLabel: "Witch",
// 	},
// 	"witch_bride": {
// 		Category: "Zombies",
// 		SubLabel: "Witch",
// 	},
// 	"common_*": {
// 		Category: "Zombies",
// 		SubLabel: "Zombie",
// 	},
// }
