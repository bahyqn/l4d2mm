package internal

import (
	"fmt"
	"l4d2mm/internal/schema"
	"path/filepath"
	"runtime"
	"strings"

	// "github.com/NublyBR/go-vpk"
	"github.com/bahyqn/vvdf"
	"github.com/bahyqn/vvpk"
)

const (
	addonlist        = "addonlist.txt"
	localModsPath    = "addons"
	workshopModsPath = "addons/workshop"
)

var maxWrokers = runtime.NumCPU() * 2

type TempVpk struct {
	ModName string
	Addons  []string
	VpkFile []string
}

type Vpk struct {
	AppDi       *App
	ppath       string
	Addonlist   []string
	AddonsDir   string
	WorkshopDir string
	Mods        []schema.Mod
	Collections []string
}

func NewVpk(di *App) *Vpk {
	return &Vpk{
		AppDi: di,
	}
}

func (vpk *Vpk) SetupPath(ppath string) {
	vpk.ppath = ppath
	// vpk.Addonlist = filepath.Join(vpk.ppath, "addonlist.txt")

	// ok := vpk.VerificationPath(vpk.Addonlist)

	// if !ok {
	// 	fmt.Println("Selected path is invalid.")

	// 	vpk.ppath = ""
	// 	// vpk.Addonlist = ""

	// 	return
	// }
	vpk.AddonsDir = filepath.Join(vpk.ppath, localModsPath)
	vpk.WorkshopDir = filepath.Join(vpk.ppath, workshopModsPath)

	// fmt.Println("addonlist.txt: ", vpk.Addonlist)
	// fmt.Println("addonsDir: ", vpk.AddonsDir)
	// fmt.Println("workshopDir: ", vpk.WorkshopDir)

	vpk.OpenAddonlist()
	vpk.ReadAllVpk(vpk.WorkshopDir)
}

func (vpk *Vpk) ReadVpkInfo(mod *schema.Mod) schema.VpkInfo {
	var absPath = ""

	if mod.IsFromWorkshop {
		absPath = filepath.Join(vpk.ppath, workshopModsPath, mod.Id+".vpk")
	} else {
		absPath = filepath.Join(vpk.ppath, localModsPath+".vpk")
	}

	// pak, err := vpk.OpenAny(absPath)
	fmap := vvpk.OpenVpk(absPath)

	vpkInfo := schema.VpkInfo{}

	// fmt.Printf("id: %s --> name: %s \n", mod.Id, mod.Name)
	if addoninfo, err := vvdf.StringToMap(fmap.Addoninfo); err == nil {
		vpkInfo.Addoninfo = addoninfo
	}

	for _, el := range fmap.Missions {
		if missions, err := vvdf.StringToMap(el); err == nil {
			vpkInfo.Missions = append(vpkInfo.Missions, missions)
		}
	}

	vpkInfo.Version = int(fmap.Version)

	return vpkInfo
}

func (vpk *Vpk) ReadAllVpk(ppath string) {
	pattern := filepath.Join(ppath, "*.vpk")
	tvpks, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for idx, p := range tvpks {
		// 1xxxxx.vpk
		// vpk.ReadVpkInfo()
		vpk.Mods = append(vpk.Mods, schema.Mod{
			Idx:            idx,
			Id:             strings.Split(filepath.Base(p), ".")[0],
			IsEnable:       true,
			IsFromWorkshop: true,
		})

		info := vpk.ReadVpkInfo(&vpk.Mods[idx])

		vpk.Mods[idx].Addoninfo = info.Addoninfo
		vpk.Mods[idx].Missions = info.Missions
		vpk.Mods[idx].Version = int(info.Version)

		if v, ok := GetValue[string](info.Addoninfo, "addontitle"); ok {
			vpk.Mods[idx].Name = v
		}
		if v, ok := GetValue[string](info.Addoninfo, "addonauthor"); ok {
			vpk.Mods[idx].Author = v

			if vpk.Mods[idx].Id == "2598614815" {
				fmt.Println(v)
			}
		}
	}
}

func (vpk *Vpk) OpenAddonlist() {
	tpath := filepath.Join(vpk.ppath, addonlist)

	if ok := StatPath(tpath); !ok {
		return
	}

	var err error
	if vpk.Addonlist, err = vvpk.OpenAddonlist(tpath); err != nil {
		return
	}
}

func (vpk *Vpk) DisableMod(mod *schema.Mod) {
	vpk.Mods[mod.Idx].IsEnable = !vpk.Mods[mod.Idx].IsEnable
	vvpk.UpdateModStatus(vpk.Addonlist, mod.Id)
}

func (vpk *Vpk) GetLocalMods(addonsDir string) {

}

func (vpk *Vpk) GetWorkshopMods(WorkshopDir string) {

}

func (vpk *Vpk) InsertCollection() {

}
