package internal

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"l4d2mm/internal/schema"

	// "github.com/NublyBR/go-vpk"
	"github.com/bahyqn/vvdf"
	"github.com/bahyqn/vvpk"
	"github.com/go-gui-org/go-gui/gui"
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

	// steam/steamapps/common/Left 4 Dead 2/left4dead2/addons
	AddonsDir string
	// steam/steamapps/common/Left 4 Dead 2/left4dead2/addons/workshop
	WorkshopDir string
	Mods        []schema.Mod
	Collections []schema.Collection

	// addonlist.txt
	AddonlistRawStr []string
	// key 11111	value: index (index 111 in addonlistRawStr)
	AddonlistIdxMap map[string]vvpk.IdxMap
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

	// fmt.Println("addonsDir: ", vpk.AddonsDir)
	// fmt.Println("workshopDir: ", vpk.WorkshopDir)

	vpk.OpenAddonlist()
	// vpk.ReadAllVpk(vpk.WorkshopDir)
}

func (vpk *Vpk) ModPath(mod *schema.Mod) string {
	path := ""

	if mod.Location == schema.LocationGame {
		if mod.IsFromWorkshop {
			path = vpk.WorkshopDir
		} else {
			path = vpk.AddonsDir
		}
		path = filepath.Join(path, mod.Filename)
	} else {
		p, err := os.Getwd()

		if err == nil {
			path = p
		}
	}

	return path
}

func (vpk *Vpk) SyncVPKs(path string) {
	if path == "" {
		return
	}

	vpk.SetupPath(path)
	// sqlite3 := vpk.AppDi.DI.SQL
	dbMods, err := vpk.AppDi.DI.SQL.QueryAllMods()
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

	vpk.syncWorkshopVPKs(dbMods, twvpks, wmods, schema.LocationGame)

	// steam/steamapps/common/Left 4 Dead 2/left4dead2/addons
	tlvpks, lmods, ok := vpk.GetMods(vpk.AddonsDir)
	if !ok {
		panic("Failed to query mods from adddons")
	}
	vpk.syncAddonsVPKs(dbMods, tlvpks, lmods, schema.LocationGame)

	err = GLOBALAPP.DI.SQL.InsertMods(&vpk.Mods)
	if err != nil {
		panic("inserts mods to sqlite failed.")
	}

	// Get all of the vpks tha there's new vpks
	clear(dbMods)
	dbMods = dbMods[:0]
	dbMods, err = vpk.AppDi.DI.SQL.QueryAllMods()
	vpk.Mods = append(vpk.Mods, dbMods...)
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, dbMods...)
	GLOBALAPP.DynamicPageSelect()
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

func (vpk *Vpk) syncWorkshopVPKs(dbmods []schema.Mod, twvpks []string, wmods map[string]string, location schema.ModLocation) {
	// sqlite has no any mods, but have l4d2 path
	if len(dbmods) == 0 {
		vpk.parseAndSaveNewVPKs(twvpks, true, location)
		return
	}

	// sqlite has mods
	var newMods []string
	cloneWmods := maps.Clone(wmods)

	for _, el := range dbmods {
		_, ok := cloneWmods[el.ID]
		if ok && el.IsFromWorkshop {
			delete(cloneWmods, el.ID)
		}
	}

	if len(cloneWmods) == 0 {
		return
	}

	for _, v := range cloneWmods {
		newMods = append(newMods, v)
	}

	vpk.parseAndSaveNewVPKs(newMods, true, location)
}

func (vpk *Vpk) syncAddonsVPKs(dbmods []schema.Mod, tlmods []string, lmods map[string]string, location schema.ModLocation) {
	if len(dbmods) == 0 {
		vpk.parseAndSaveNewVPKs(tlmods, false, location)
		return
	}

	// var newMods []string
}

func (vpk *Vpk) GetModAbsPath(mod *schema.Mod) string {
	if mod.IsFromWorkshop {
		return filepath.Join(vpk.ppath, workshopModsPath, mod.ID+".vpk")
	}
	return filepath.Join(vpk.ppath, localModsPath, mod.ID+".vpk")
}

func (vpk *Vpk) ReadVpkInfo(path string, mod *schema.Mod) schema.VpkInfo {
	// var absPath = ""

	// absPath = vpk.GetModAbsPath(mod)

	// pak, err := vpk.OpenAny(absPath)
	vpkSummary := vvpk.OpenVpk(path)

	vpkInfo := schema.VpkInfo{}

	// fmt.Printf("id: %s --> name: %s \n", mod.Id, mod.Name)
	if addoninfo, err := vvdf.StringToMap(vpkSummary.Addoninfo); err == nil {
		vpkInfo.Addoninfo = addoninfo
	}

	for _, el := range vpkSummary.Missions {
		if missions, err := vvdf.StringToMap(el); err == nil {
			vpkInfo.Missions = append(vpkInfo.Missions, missions)
		}
	}

	vpkInfo.VpkVersion = vpkSummary.VpkVersion
	vpkInfo.Treesize = vpkSummary.TreeSize

	return vpkInfo
}

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

	batchSize := 500
	var missingIDs []string
	existngMap := make(map[string]bool, idsLength)

	for i := 0; i < idsLength; i += batchSize {
		end := i + batchSize
		if end > idsLength {
			end = idsLength
		}

		batch := ids[i:end]

		var batchExistIDs []string
		err := vpk.AppDi.DI.SQL.SDB.Model(&schema.Mod{}).Where("id in ?", batch).Pluck("id", &batchExistIDs).Error
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

func (vpk *Vpk) parseAndSaveNewVPKs(tvpks []string, IsFromWorkshop bool, location schema.ModLocation) {
	localMods := make([]schema.Mod, len(tvpks))

	var wg sync.WaitGroup
	ch := make(chan struct{}, maxWrokers)

	for i, p := range tvpks {
		wg.Add(1)
		ch <- struct{}{}

		go func(idx int, path string, loc schema.ModLocation) {
			defer wg.Done()
			defer func() { <-ch }()
			// 1xxxxx.vpk
			// vpk.ReadVpkInfo()
			mod := schema.Mod{
				Idx:            idx,
				Filename:       strings.Split(filepath.Base(path), ".")[0],
				ID:             strings.Split(filepath.Base(path), ".")[0],
				IsEnable:       true,
				Location:       loc,
				IsFromWorkshop: IsFromWorkshop,
			}

			info := vpk.ReadVpkInfo(path, &mod)

			mod.Addoninfo = info.Addoninfo
			mod.Missions = info.Missions
			mod.VpkVersion = info.VpkVersion
			mod.Treesize = info.Treesize

			if !IsFromWorkshop {
				mod.ID = vpk.newAddonsVpkId(path, mod.Treesize)
			}

			if v, ok := GetValue[string](info.Addoninfo, "addontitle"); ok {
				mod.ModName = v
			}
			if v, ok := GetValue[string](info.Addoninfo, "addonauthor"); ok {
				mod.Author = v
			}

			localMods[idx] = mod
		}(i, p, location)

	}

	wg.Wait()
	close(ch)

	if err := GLOBALAPP.DI.SQL.InsertMods(&localMods); err != nil {
		fmt.Println("insert vpks to sqlite failed.")
	}
	// vpk.Mods = append(vpk.Mods, localMods...)
	// GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, localMods...)
	// GLOBALAPP.DynamicPageSelect()
}

func (vpk *Vpk) newAddonsVpkId(path string, treesize uint32) string {
	if size, err := GetFileSize(path); err == nil {
		newId := uint64(treesize) + size
		return strconv.FormatUint(newId, 10)
	}
	return ""
}

func (vpk *Vpk) FecthVpks() {
}

func (vpk *Vpk) OpenAddonlist() {
	vpk.AddonlistRawStr = vpk.AddonlistRawStr[:0]
	clear(vpk.AddonlistIdxMap)

	tpath := filepath.Join(vpk.ppath, addonlist)
	if ok := StatPath(tpath); !ok {
		panic("addonlist path is invalid")
		// return
	}

	addonlistRawStr, idxMap, err := vvpk.OpenAddonlist(tpath)
	if err == nil {
		vpk.AddonlistRawStr = addonlistRawStr
		vpk.AddonlistIdxMap = idxMap

		return
	}
}

func (vpk *Vpk) DisableMod(mod *schema.Mod) {
	vpk.Mods[mod.Idx].IsEnable = !vpk.Mods[mod.Idx].IsEnable
	vvpk.UpdateModStatus(vpk.AddonlistRawStr, mod.ID, "0")
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

func (vpk *Vpk) clearAddonlistInfo() {
	clear(vpk.AddonlistRawStr)
	vpk.AddonlistRawStr = vpk.AddonlistRawStr[:0]
	clear(vpk.AddonlistIdxMap)
}

func (vpk *Vpk) SyncAddonlistToDB() {
	vpk.clearAddonlistInfo()

	vpk.OpenAddonlist()

	for idx, el := range vpk.Mods {
		id := el.ID
		if vmap, ok := vpk.AddonlistIdxMap[id]; ok {
			switch vmap.Value {
			case "0":
				vpk.Mods[idx].IsEnable = false
			case "1":
				vpk.Mods[idx].IsEnable = true
			}
		}
	}

	clear(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx)
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[:0]
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[:0]
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, vpk.Mods...)

	ToastSuccess("Synced", "Sync addonlist to sqlite successful")
}

func (vpk *Vpk) FetchCollection() {
	dc := &GLOBALAPP.ComponentStatus.PageTools.DeleteCollection

	if dc.Toggled {
		return
	}

	colls, err := vpk.AppDi.DI.SQL.QueryCollections()
	if err != nil {
		return
	}

	vpk.Collections = append(vpk.Collections, colls...)

	dc.Header = append(dc.Header, schema.TableField{
		Text:     "",
		MinWidth: 100,
	})
	dc.Header = append(dc.Header, schema.TableField{
		Text:     "id",
		MinWidth: 200,
	})
	dc.Header = append(dc.Header, schema.TableField{
		Text:     "mod-name",
		MinWidth: 300,
	})
	dc.Header = append(dc.Header, schema.TableField{
		Text: "mods-amount",
	})

	tempSelectOption := []gui.SelectOption{}

	for _, el := range vpk.Collections {
		tf := []schema.TableField{}

		tf = append(tf, schema.TableField{
			Text:     el.ID,
			MinWidth: 200,
		})
		tf = append(tf, schema.TableField{
			Text:     el.Name,
			MinWidth: 300,
		})
		tf = append(tf, schema.TableField{
			Text:     strconv.Itoa(len(el.Mods)),
			MinWidth: 200,
		})

		dc.RowData = append(dc.RowData, tf)
		dc.Selected = append(dc.Selected, false)

		tempSelectOption = append(tempSelectOption, gui.SelectOption{
			Label: el.Name,
			Value: el.Name,
		})
	}

	GenerateSelectOptions("Collections", tempSelectOption)
}
