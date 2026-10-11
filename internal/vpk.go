package internal

import (
	"fmt"
	"l4d2mm/internal/schema"
	"maps"
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
	AppDi *App
	ppath string
	// addonlist.txt
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

func (vpk *Vpk) SetupPath(path string) {
	vpk.ppath = path

	vpk.AddonsDir = filepath.Join(vpk.ppath, localModsPath)
	vpk.WorkshopDir = filepath.Join(vpk.ppath, workshopModsPath)

	// fmt.Println("addonlist.txt: ", vpk.Addonlist)
	// fmt.Println("addonsDir: ", vpk.AddonsDir)
	// fmt.Println("workshopDir: ", vpk.WorkshopDir)

	vpk.OpenAddonlist()
	// vpk.ReadAllVpk(vpk.WorkshopDir)
}

func (vpk *Vpk) SyncVPKs(path string) {
	if path == "" {
		return
	}

	vpk.SetupPath(path)
	// sqlite3 := vpk.AppDi.DI.Sql
	mods, err := vpk.AppDi.DI.Sql.QueryAllMods()
	if err != nil {
		panic("Failed query table mods from sqlite")
	}

	// .../steam/steamapps/common/Left 4 Dead 2/left4dead2/addons/workshop
	// twvpks: sort by disk
	// wmods: map{ "1111": ".../steam/steamapps/common/Left 4 Dead 2/left4dead2/addons/workshop/1111.vpk"}
	twvpks, wmods, ok := vpk.GetMods(vpk.WorkshopDir)
	if !ok {
		panic("Failed to query mods from workshop")
	}
	vpk.syncWorkshopVPKs(mods, twvpks, wmods)

	// steam/steamapps/common/Left 4 Dead 2/left4dead2/addons
	// vpk.findNewVPKs(vpk.AddonsDir, mods)
}

func (vpk *Vpk) findNewVPKs(path string, modsFromDB []schema.Mod) []string {
	var newVPKs []string

	// var wg sync.WaitGroup
	// ch := make(chan struct{}, maxWrokers)

	// if localMods, ok := vpk.GetMods(path); ok {
	// 	for idx, el := range localMods {

	// 	}
	// }

	return newVPKs
}

func (vpk *Vpk) syncWorkshopVPKs(dbmods []schema.Mod, twvpks []string, wmods map[string]string) {
	// sqlite has no any mods, but have l4d2 path
	if len(dbmods) == 0 {
		vpk.parseAndSaveNewVPKs(twvpks)
		return
	}

	// sqlite has mods
	var newMods []string
	cloneWmods := maps.Clone(wmods)

	for _, el := range dbmods {
		_, ok := cloneWmods[el.Id]
		if ok && el.IsFromWorkshop {
			delete(cloneWmods, el.Id)
		}

		newMods = append(newMods, el.Id)
	}

	if len(cloneWmods) == 0 {
		return
	}

	vpk.parseAndSaveNewVPKs(newMods)
}

func (vpk *Vpk) scanAddonsVPKs() {

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

	vpkInfo.VpkVersion = int(fmap.Version)

	return vpkInfo
}

// func (vpk *Vpk) FecthAllMods(path string) bool {
// 	tvpks := []string{}

// 	wsvpks, ok := vpk.GetMods(vpk.WorkshopDir)
// 	if ok {
// 		tvpks = append(tvpks, wsvpks...)
// 	}

// 	localaVpks, ok := vpk.GetMods(vpk.AddonsDir)
// 	if ok {
// 		tvpks = append(tvpks, localaVpks...)
// 	}

// 	if len(tvpks) == 0 {
// 		return false
// 	}
// 	return true
// }

func (vpk *Vpk) GetModsSlice(path string) ([]string, bool) {
	pattern := filepath.Join(path, "*.vpk")
	tvpks, err := filepath.Glob(pattern)

	if err != nil {
		return nil, false
	}
	return tvpks, true
}

func (vpk *Vpk) GetMods(path string) ([]string, map[string]string, bool) {
	pattern := filepath.Join(path, "*.vpk")
	tvpks, err := filepath.Glob(pattern)
	if err != nil {
		return nil, nil, false
	}

	filenames := make(map[string]string, len(tvpks))
	for _, fullPath := range tvpks {
		key := strings.Split(filepath.Base(fullPath), ".")[0]
		filenames[key] = fullPath
	}

	return tvpks, filenames, true
}

func (vpk *Vpk) getMissingModIDs(ids []string) ([]string, error) {
	idsLength := len(ids)

	if idsLength == 0 {
		return nil, nil
	}

	var batchSize = 500
	var missingIDs []string
	existngMap := make(map[string]bool, idsLength)

	for i := 0; i < idsLength; i += batchSize {
		end := i + batchSize
		if end > idsLength {
			end = idsLength
		}

		batch := ids[i:end]

		var batchExistIDs []string
		err := vpk.AppDi.DI.Sql.SDB.Model(&schema.Mod{}).Where("id in ?", batch).Pluck("id", &batchExistIDs).Error

		if err != nil {
			return nil, err
		}

		for _, id := range batchExistIDs {
			existngMap[id] = true
		}

		for _, id := range ids {
			if !existngMap[id] {
				missingIDs = append(missingIDs, id)
			}
		}
	}
	return missingIDs, nil
}

func (vpk *Vpk) parseAndSaveNewVPKs(tvpks []string) {

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

			// local addons
			// if idx%2 == 0 {
			// 	mod.IsFromWorkshop = false
			// }

			info := vpk.ReadVpkInfo(&mod)

			mod.Addoninfo = info.Addoninfo
			mod.Missions = info.Missions
			mod.VpkVersion = int(info.VpkVersion)

			if v, ok := GetValue[string](info.Addoninfo, "addontitle"); ok {
				mod.Name = v
			}
			if v, ok := GetValue[string](info.Addoninfo, "addonauthor"); ok {
				mod.Author = v
			}
			// if v, ok := GetValue[string](info.Addoninfo, "addonurl0"); ok {
			// 	mod.Url = v
			// }

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
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, localMods...)
	GLOBALAPP.DynamicPageSelect()
}

func (vpk *Vpk) FecthVpks() {

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
