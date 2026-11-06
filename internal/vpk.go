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

var (
	maxWrokers = runtime.NumCPU() * 2
	isClassify = false
)

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
	for i := range dbMods {
		dbMods[i].Idx = i
		vpk.Mods = append(vpk.Mods, dbMods[i])
	}
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

func (vpk *Vpk) ListAllFiles(mod *schema.Mod) {
	if arch, err := vvpk.OpenVpkDev(vpk.GetModAbsPath(mod)); err == nil {
		for i, item := range arch.Entries {
			fmt.Printf("%d - %s -> %s -> %s\n", i, item.Extension, item.Path, item.Filename)
		}
	}
}

func (vpk *Vpk) Classify() {
	// filtering vpks that no category
	unCategoried := []*schema.Mod{}

	for i := range vpk.Mods {
		if vpk.Mods[i].Category == "" {
			unCategoried = append(unCategoried, &vpk.Mods[i])
		}
	}

	if len(unCategoried) == 0 {
		return
	}
	// 1.
	// double bucket
	// bucket one: confirmed maps
	// bucket two: unknown
	confirmedBucket := map[string][]*schema.Mod{}
	unknownBucket := map[string][]*schema.Mod{}

	tname := ""

	for i := range unCategoried {
		if unCategoried[i].ModName == "" {
			tname = "xxxxxxxxxxxxxxxxxxxxx"
			// continue
		}

		tempName := strings.Split(unCategoried[i].ModName, " ")
		if len(tempName) == 1 {
			if len(tempName[0])-3 > 0 {
				tname = tempName[0][:len(tempName[0])-3]
			}
		} else if len(tempName) > 1 {
			tname = tempName[0]
		}

		if len(unCategoried[i].Missions) > 0 {
			unCategoried[i].Category = "Maps"
			confirmedBucket[tname] = append(confirmedBucket[tname], unCategoried[i])
		} else {
			// if tname == "" {
			// 	tname = "xxxxxxxxxxxxxxxxxxxxx"
			// }
			unknownBucket[tname] = append(unknownBucket[tname], unCategoried[i])
		}
	}
	tname = ""

	// 2.
	// map1: storage similar vpks
	// slice2: maybe its all the scripts, hud
	// to get maps vpks from unknownBucket that existed inside the confirmedBucket
	slice2 := []*schema.Mod{}
	//  storage length similar vpks, if v length > 1,maybe that's the maps
	sameMap := map[int][]*schema.Mod{}
	for key, value := range confirmedBucket {
		vpks, ok := unknownBucket[key]
		if !ok {
			continue
		}

		// clean cache
		clear(sameMap)
		clear(slice2)
		slice2 = slice2[:0]

		for i := range vpks {
			match := IsRelatedPart(value[0].ModName, vpks[i].ModName)

			// fmt.Printf("%t %d %s \n", match, vpks[i].Idx, vpks[i].ModName)
			if match {
				vpk.Mods[vpks[i].Idx].Category = "Maps"
				delete(unknownBucket, key)
				continue
			}

			tlen := len(vpks[i].ModName) / 4
			sameMap[tlen] = append(sameMap[tlen], vpks[i])
		}

		// last check if there'are mpas
		if len(sameMap) > 0 {
			for _, v := range sameMap {
				// there's a problem
				if len(v) > 1 {
					for i := range v {
						vpk.Mods[v[i].Idx].Category = "Maps"
					}
					continue
				}
				// not map
				slice2 = append(slice2, v...)
			}
		}

		if len(slice2) == 0 {
			delete(unknownBucket, key)
		}

		// put back not map vpks
		if len(slice2) > 0 {
			clear(unknownBucket[key])
			unknownBucket[key] = unknownBucket[key][:0]
			unknownBucket[key] = append(unknownBucket[key], slice2...)
		}
	}

	clear(slice2)
	// cache not recgnized vpks
	slice2 = slice2[:0]
	// cache opened vpk's file
	filesMap := map[int][]vvpk.Metadata{}
	matched := false
	// 3. recognize the model
	for _, v := range unknownBucket {
		for _, item := range v {
			if arch, err := vvpk.OpenVpkDev(vpk.GetModAbsPath(item)); err == nil {
				matched = false
				for _, file := range arch.Entries {
					if category, ok := MDL[file.Filename]; ok {
						matched = true
						vpk.Mods[item.Idx].Category = category.Category
						vpk.Mods[item.Idx].SubType = category.SubLabel
						break
					}
				}

				if !matched {
					slice2 = append(slice2, item)
					// filesMap[item.Idx] = append(filesMap[item.Idx], arch.Entries...)
					filesMap[item.Idx] = arch.Entries
				}
			}
		}
	}

	exCounts := make(map[string]int)
	// var survivingSlice []SliceItem
	var isScript bool

	fmt.Println("slice2 -- > ", len(slice2))
	// 4. maybe is script, hud, sound, effect, other
	for _, item := range slice2 {
		clear(exCounts)
		isScript = false
		for _, row := range filesMap[item.Idx] {
			// extension	path	filename
			// xxx          scripts xxxxxxxxx
			if row.Path == "scripts" {
				isScript = true
				vpk.Mods[item.Idx].Category = "Scripts"
				break
			}

			// exclus: addoninfo.txt missions/*.txt
			if row.Extension == "txt" && (row.Path == "" && strings.ToLower(row.Path) == "missions") {
				continue
			}
			exCounts[row.Extension] += 1
			// extension	path			filename
			// xxx          scripts/xxxx	xxxxxxxxx
			// maybe extension is better
		}

		if !isScript {
			fmt.Printf("%d ---> %s \n", item.Idx, item.ModName)
			switch getDominantExtension(exCounts, len(filesMap[item.Idx])) {
			case "wav":
				vpk.Mods[item.Idx].Category = "Sounds"
			case "vmt":
				vpk.Mods[item.Idx].Category = "Misc"
				vpk.Mods[item.Idx].SubType = "Materials"
			case "vtf":
				vpk.Mods[item.Idx].Category = "Misc"
				vpk.Mods[item.Idx].SubType = "Materials"
			case "nut":
				vpk.Mods[item.Idx].Category = "Script"
			case "txt":
				vpk.Mods[item.Idx].Category = "Script"
			case "pcf":
				vpk.Mods[item.Idx].Category = "Misc"
				vpk.Mods[item.Idx].SubType = "Particles"
			default:
				vpk.Mods[item.Idx].Category = "Misc"
				vpk.Mods[item.Idx].SubType = "Other"
			}
		}
	}

	for i, item := range vpk.Mods {
		fmt.Printf("%d %s --> %s - %s \n", i, item.ModName, vpk.Mods[i].Category, vpk.Mods[i].SubType)
	}

	clear(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx)
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx[:0]
	GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx = append(GLOBALAPP.ComponentStatus.PageMods.ModsBySelectIdx, vpk.Mods...)
}

func getDominantExtension(exCounts map[string]int, metadataLength int) string {
	ex := ""
	exf := 0.0
	for k, v := range exCounts {
		// if k == "txt" {
		// 	continue
		// }
		if float64(v)/float64(metadataLength) > exf {
			exf = float64(v) / float64(metadataLength)
			ex = k
		}

		fmt.Printf("extension: %s, amount: %d \n", k, v)
	}
	fmt.Println(ex)
	fmt.Println()
	return ex
}
