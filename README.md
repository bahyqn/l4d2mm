
## About l4d2mm
l4d2mm is a modern, native mod manager designed specifically for Left 4 Dead 2 on Linux.

For a long time, the Linux community has lacked a polished, dedicated mod management tool for L4D2. This project aims to bridge that gap, providing Linux gamers with a smooth, native, and reliable way to manage their workshop items, add-ons, and game configurations.

### 323165424.vpk
- BaseName(): flashlight001.vtf
- Filename(): materials/effects/flashlight001.vtf
- 
- BaseName(): addoninfo.txt
- Filename(): addoninfo.txt

#### addoninfo.txt
1. HD Luminous ladder
```
"AddonInfo"
{
    addontitle "HD Luminous ladder" 
    addonversion 1
    addontagline "This mod replaces the Ladders.
    It shines in the dark and is easier to spot.
    This is a gift for those cute novice blind gamer    :P
    " 
    addonauthor "徒手开根号genhao"
    addonauthorSteamID "http://steamcommunity.com/id/911110/"
    addonDescription "This mod replaces the Ladders.
    It shines in the dark and is easier to spot"

    addonContent_Script 0
    addonContent_Music 0
    addonContent_Sound 0
    addonContent_prop 1 
    addonContent_Prefab 0 

    addonContent_BackgroundMovie 0 
    addonContent_Survivor 0 
    addonContent_BossInfected 0
    addonContent_CommonInfected 0 
    addonContent_WeaponModel 0
    addonContent_weapon 0 
    addonContent_Skin 1 
    addonContent_Spray 0 
    addonContent_Map 0
}
```

2. Dark wood(extended)
```
"AddonInfo"
{
	"addonSteamAppID"                  "550"
	"addonTitle"                       "Dark Wood (Extended)"
	"addonVersion"                     "1.9"
	"addonTagline"                     "Going to the woods at night wasn't such a good idea."
	"addonAuthor"                      "Phaeton"
	"addonURL0"                        "http://steamcommunity.com/workshop/filedetails/?id=575682109"
 
	"addonContent_Campaign"			"1"
	"addonContent_Survival"			"1"
	"addonContent_Sound"			"1"
	"addonContent_Prop"				"1"
	"addonContent_Script"			"1"
 
	"addonDescription"              "The survivors find themselves trapped on a small stretch of road at the middle of the night. Fleeing from the infected, they have no choice but to plunge into the darkness of the adjacent forest. They will have a long and dangerous journey through the dark woods, creepy caves, abandoned industrial zones, secret laboratory (which must be destroyed), 'Silent Hill' style deep underground prison, ancient catacombs, and finally face their last battle in the church on a flooded cemetery."
}
```

3. Urban Flight
```
"AddonInfo"
{
     addonSteamAppID 			550			// 550 is the app ID for Left 4 Dead 2
     addontitle 			"Urban Flight"
     addonversion 			"12"
     addontagline 			"No plane, no gain"	// short description
     addonauthor			"The Rabbit"
     addonSteamGroupName		"group"
     addonauthorSteamID			"The_Rabbit42"
     addonContent_Campaign		1			// campaign mode included
     addonContent_Survival		1			// survival mode included
     addonContent_Scavenge		1			// scavenge mode included
     addonContent_Versus		1			// versus   mode included
     addonURL0				"http://steamcommunity.com/sharedfiles/filedetails/?id=121086524"
								// where people can download your VPK

     addonDescription			"The city is burning. As ash falls, the survivors attempt to flee to a small military airfield on the other side of the river. It's a straight path down the main boulevard, but nothing is ever simple. Crashed cars, blazing fires, police barricades, and the city itself all stand in their way." 

     addonContent_Script		1			// Has Scripts 
     addonContent_Music			0			// Has Custom Music 
     addonContent_Sound			1			// Has Custom Sound 
     addonContent_prop			1			// This Add-on provides new props, 
     addonContent_Prefab		0			// Provides new prefabs 
     addonContent_BackgroundMovie	0			// Provides a replacement for the background movie.
     addonContent_Survivor 		0			// Provides a new survivor model. 0=false, 1=true, String in quotes if replaces specific single character, i.e. "Coach"
     								// eg addonContent_Survivor "rochelle" works fine, no number.

     addonContent_BossInfected		1			// Provides a new boss infected model. Break these out?
     addonContent_CommonInfected	0			// Provides a new common infected model
     Content_WeaponModel		0			// Provides a new appearance to existing weapons, but does not change their function
     Content_weapon			0			// provides new weapons or new zombie killing functionality, i.e. guns, explosives, booby traps, hot tar,
     addonContent_Skin			1			// 0 if no new skin textures for existing models. 1 if multiple skin pack. String in quotes if specific single skin
     addonContent_Spray			0			// Provides new sprays.
     addonContent_Map			0			// Add-on provides a standalone map
}
```