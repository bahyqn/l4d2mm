package internal

import (
	"fmt"
	"l4d2mm/internal/schema"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

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

	vpk.AppDi.DynamicMods()
}

func (vpk *Vpk) GetModAbsPath(mod *schema.Mod) string {
	if mod.IsFromWorkshop {
		return filepath.Join(vpk.ppath, workshopModsPath, mod.Id+".vpk")
	}
	return filepath.Join(vpk.ppath, localModsPath+".vpk")
}

func (vpk *Vpk) ReadVpkInfo(mod *schema.Mod) schema.VpkInfo {
	var absPath = ""

	absPath = vpk.GetModAbsPath(mod)

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

	localMods := make([]schema.Mod, len(tvpks))

	var wg sync.WaitGroup
	ch := make(chan struct{}, maxWrokers)

	for i, p := range tvpks {
		wg.Add(1)
		ch <- struct{}{}

		go func(idx int, path string) {
			defer wg.Done()
			defer func() { <-ch }()
			// 1xxxxx.vpk
			// vpk.ReadVpkInfo()
			mod := schema.Mod{
				Idx:            idx,
				Id:             strings.Split(filepath.Base(path), ".")[0],
				IsEnable:       true,
				IsFromWorkshop: true,
			}

			info := vpk.ReadVpkInfo(&mod)

			mod.Addoninfo = info.Addoninfo
			mod.Missions = info.Missions
			mod.Version = int(info.Version)

			if v, ok := GetValue[string](info.Addoninfo, "addontitle"); ok {
				mod.Name = v
			}
			if v, ok := GetValue[string](info.Addoninfo, "addonauthor"); ok {
				mod.Author = v
			}

			// if v, ok := GetValue[string](info.Addoninfo, "title"); ok {
			// 	vpk.Mods[idx].Name = v
			// }
			// if v, ok := GetValue[string](info.Addoninfo, "author"); ok {
			// 	vpk.Mods[idx].Author = v
			// }
			localMods[idx] = mod
		}(i, p)

	}

	wg.Wait()
	close(ch)

	vpk.Mods = append(vpk.Mods, localMods...)
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

func (vpk *Vpk) SaveVpkInfo(mod *schema.Mod) {
	// archive, err := vvpk.OpenVpkDev(vpk.GetModAbsPath(mod))

	fmt.Printf("%+v\n\n", mod.Addoninfo)

	// if err != nil {
	// 	fmt.Println("Error: OpenVpkDev")
	// 	return
	// }

	// for idx, el := range archive.Entries {
	// 	fmt.Printf("%d:  %s ---> %s ---> %s\n", idx, el.Extension, el.Path, el.Filename)
	// }
}

func (vpk *Vpk) GetLocalMods(addonsDir string) {

}

func (vpk *Vpk) GetWorkshopMods(WorkshopDir string) {

}

func (vpk *Vpk) InsertCollection() {

}
