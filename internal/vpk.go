package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"l4d2mm/internal/schema"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/NublyBR/go-vpk"
)

const (
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
	Addonlist   string
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

func (vp *Vpk) SetupPath(ppath string) {
	vp.ppath = ppath
	vp.Addonlist = filepath.Join(vp.ppath, "addonlist.txt")

	ok := vp.VerificationPath(vp.Addonlist)

	if !ok {
		fmt.Println("Selected path is invalid.")

		vp.ppath = ""
		vp.Addonlist = ""

		return
	}
	vp.AddonsDir = filepath.Join(vp.ppath, localModsPath)
	vp.WorkshopDir = filepath.Join(vp.ppath, workshopModsPath)

	// fmt.Println("addonlist.txt: ", vp.Addonlist)
	// fmt.Println("addonsDir: ", vp.AddonsDir)
	// fmt.Println("workshopDir: ", vp.WorkshopDir)

	vp.ReadAllVpk(vp.WorkshopDir)
}

func (vp *Vpk) VerificationPath(ppath string) bool {
	_, err := os.Stat(ppath)

	// fmt.Println("ppath: ", ppath)

	if err == nil {
		return true
	}
	return false
}

func (vp *Vpk) ReadVpkInfo(mod *schema.Mod) {
	var absPath = ""

	if mod.IsFromWorkshop {
		absPath = filepath.Join(vp.ppath, workshopModsPath, mod.Id+".vpk")
	} else {
		absPath = filepath.Join(vp.ppath, localModsPath+".vpk")
	}

	pak, err := vpk.OpenAny(absPath)

	if err != nil {
		panic(err)
	}

	defer pak.Close()

	for idx, file := range pak.Entries() {
		fmt.Printf("%d: %s \n", idx, file.Filename())

		if file.Filename() == "addoninfo.txt" {
			reader, err := file.Open()

			if err != nil {
				fmt.Println("Failed to open addoninfo.txt")
				continue
			}

			content, err := io.ReadAll(reader)
			reader.Close()

			if err != nil {
				fmt.Println("err: ", err)
			}

			fmt.Println(string(content))
			fmt.Println()
		}
	}
}

func (vp *Vpk) DisableMod(mod *schema.Mod) {
	vp.Mods[mod.Idx].IsEnable = !vp.Mods[mod.Idx].IsEnable
}

func (vp *Vpk) ReadAllVpk(ppath string) {
	// _, err := os.Stat(ppath)

	// fmt.Println(ppath)

	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("the vpk path is invalid.")
	// 		panic(err)
	// 	}
	// }

	pattern := filepath.Join(ppath, "*.vpk")
	tvpks, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	for idx, p := range tvpks {
		// 1xxxxx.vpk
		vp.Mods = append(vp.Mods, schema.Mod{
			Idx:            idx,
			Id:             strings.Split(filepath.Base(p), ".")[0],
			IsEnable:       true,
			IsFromWorkshop: true,
		})
	}
	// tvk := []TempVpk{}

	// for _, el := range tvpks {
	// 	// fmt.Printf("%d \t %s\n", idx, el)

	// 	// if idx == 1 {
	// 	// 	return
	// 	// }

	// 	pak, err := vpk.OpenAny(el)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer pak.Close()

	// 	tv := TempVpk{}
	// 	for _, file := range pak.Entries() {
	// 		// fmt.Printf("% 8d %s %s %d\n", file.Length(), file.Filename(), file.Basename(), file.CRC())

	// 		tv.VpkFile = append(tv.VpkFile, file.Filename())

	// 		if file.Filename() == "addoninfo.txt" {
	// 			reader, err := file.Open()
	// 			if err != nil {
	// 				fmt.Println("Failed to open addoninfo.txt")
	// 				continue
	// 			}

	// 			content, err := io.ReadAll(reader)
	// 			reader.Close()

	// 			if err != nil {
	// 				fmt.Println("err: ", err)
	// 				continue
	// 			}

	// 			vp.ToKVLines(string(content), &tv)
	// 			fmt.Println(string(content))
	// 			// fmt.Println()
	// 		}
	// 	}
	// 	tvk = append(tvk, tv)
	// }

	// jsonData, err := json.MarshalIndent(tvk, "", "	")

	// if err != nil {
	// 	fmt.Println("Failed to serializer json")
	// 	return
	// }

	// err = os.WriteFile("mods.json", jsonData, 0644)
	// if err != nil {
	// 	fmt.Println("Failed to writing")
	// }
}

func (vpk *Vpk) trimToBraces(data []byte) []byte {
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	start := bytes.IndexByte(data, '{')

	if start == -1 {
		return data
	}

	end := bytes.LastIndexByte(data, '}')
	if end == -1 || end < start {
		return data
	}

	return data[start+1 : end]
}

func (vp *Vpk) parseKVLine(line string) (key, val string, ok bool) {
	// 1. 将 Unicode 特殊空格 (\u00a0) 替换为标准空格
	line = strings.ReplaceAll(line, "\u00a0", " ")
	line = strings.TrimSpace(line)

	// 2. 过滤空行、纯注释行或大括号
	if line == "" || strings.HasPrefix(line, "//") || line == "{" || line == "}" {
		return "", "", false
	}

	var tokens []string
	var current strings.Builder
	inQuotes := false
	escaped := false

	// 3. 逐字符扫描，处理双引号、转义符与注释
	for i := 0; i < len(line); i++ {
		ch := line[i]

		if escaped {
			current.WriteByte(ch)
			escaped = false
			continue
		}

		if ch == '\\' && inQuotes {
			escaped = true
			continue
		}

		if ch == '"' {
			inQuotes = !inQuotes
			continue
		}

		// 在双引号外部时的逻辑
		if !inQuotes {
			// 遇到 // 注释，说明后面全是注释，直接结束扫描
			if ch == '/' && i+1 < len(line) && line[i+1] == '/' {
				break
			}

			// 遇到空格/Tab，说明一个 Token 结束了
			if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
					if len(tokens) == 2 { // 已经抓取到了 Key 和 Value，无需继续
						break
					}
				}
				continue
			}
		}

		current.WriteByte(ch)
	}

	// 补全最后一个 token（如果不带双引号且在行尾）
	if current.Len() > 0 && len(tokens) < 2 {
		tokens = append(tokens, current.String())
	}

	if len(tokens) >= 2 {
		return tokens[0], tokens[1], true
	}

	return "", "", false
}

func (vp *Vpk) ToKVLines(content string, tvpk *TempVpk) {
	// 去除 UTF-8 BOM
	content = strings.TrimPrefix(content, "\xef\xbb\xbf")

	scanner := bufio.NewScanner(strings.NewReader(content))
	inBlock := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(strings.ReplaceAll(line, "\u00a0", " "))

		if trimmed == "" || strings.HasPrefix(trimmed, "//") {
			continue
		}

		if strings.Contains(trimmed, "{") {
			inBlock = true
			continue
		}

		if strings.Contains(trimmed, "}") {
			inBlock = false
			break
		}

		if inBlock {
			// 使用精准的状态机解析行
			key, val, ok := vp.parseKVLine(line)
			if ok {
				tvpk.Addons = append(tvpk.Addons, key+" "+val)
			}
		}
	}
}

func (vp *Vpk) cleanLIne(raw string) string {
	if idx := strings.Index(raw, "//"); idx != -1 {
		raw = raw[:idx]
	}

	clean := strings.TrimSpace(raw)

	fields := strings.Fields(clean)

	key := ""
	value := ""

	if len(fields) > 2 {
		key = strings.Trim(fields[0], `""`)
		value = strings.Trim(fields[1], `""`)
	}
	return key + " " + value
}

func (vp *Vpk) GetLocalMods(addonsDir string) {

}

func (vp *Vpk) GetWorkshopMods(WorkshopDir string) {

}

func (vp *Vpk) InsertCollection() {

}
