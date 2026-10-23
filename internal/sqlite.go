package internal

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"l4d2mm/internal/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Sqlite3 struct {
	AppDi *App
	SDB   *gorm.DB
	Conn  *sql.DB
}

func NewSqlite3(di *App) *Sqlite3 {
	sql3 := Sqlite3{
		AppDi: di,
	}

	sql3.Connect()
	return &sql3
}

func (sql3 *Sqlite3) Connect() {
	var err error
	dir, _ := os.Getwd()
	rootPath := filepath.Clean(dir)

	dsn := "file:" + rootPath + "/data.db" + "?_journal_mode=WAL&_busy_timeout=2000"
	sql3.SDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		fmt.Println("err: ", err)
	}

	sql3.Conn, err = sql3.SDB.DB()
	sql3.Conn.SetMaxOpenConns(1)

	if err != nil {
		fmt.Println("Failed to connect sqlite3: ", err)
	}

	sql3.FetchTables()
}

func (sql3 *Sqlite3) FetchTables() {
	sql3.SDB.AutoMigrate(&schema.AppConfig{})
	sql3.SDB.AutoMigrate(&schema.Mod{})
	sql3.SDB.AutoMigrate(&schema.Collection{})

	var tcfg schema.AppConfig
	// err := sql3.SDB.Model(&schema.AppConfig{}).Where("id=?", 1).Scan(&tcfg).Error
	err := sql3.SDB.FirstOrCreate(&tcfg, schema.AppConfig{
		ID:    1,
		Theme: "gnome",
	}).Error

	if err == nil && tcfg.ID == 0 {
		config := DefaultAppConfig()
		sql3.SDB.Model(&schema.AppConfig{}).Create(&config)
	}

	sql3.AppDi.AppConfig.ID = tcfg.ID
	sql3.AppDi.AppConfig.Theme = tcfg.Theme
	sql3.AppDi.AppConfig.GamePath = tcfg.GamePath
	sql3.AppDi.AppConfig.ActiveCollectionID = tcfg.ActiveCollectionID
}

func (sql3 *Sqlite3) UpdateGamePath(path string) error {
	err := sql3.SDB.Model(&schema.AppConfig{}).
		Where("id = ?", 1).
		Update("game_path", path).Error
	if err != nil {
		fmt.Println("Failed to update game path: ", err)
		return err
	}

	sql3.AppDi.AppConfig.GamePath = path
	return nil
}

func (sql3 *Sqlite3) InsertMods(batch *[]schema.Mod) error {
	err := sql3.SDB.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(batch, 1000).Error
	})
	if err != nil {
		return err
	}
	return nil
}

func (sql3 *Sqlite3) QueryAllMods() ([]schema.Mod, error) {
	var mods []schema.Mod

	err := sql3.SDB.Model(&schema.Mod{}).Find(&mods).Error
	if err != nil {
		return nil, err
	}

	return mods, nil
}

func (sql3 *Sqlite3) DiffMods(batch []string) ([]string, error) {
	var batchExistIDs []string

	err := sql3.SDB.Model(&schema.Mod{}).Where("id in ?", batch).Pluck("id", &batchExistIDs).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (sql3 *Sqlite3) QueryCollections() ([]schema.Collection, error) {
	colls := []schema.Collection{}

	err := sql3.SDB.Model(&schema.Collection{}).Scan(&colls).Error
	if err != nil {
		fmt.Println("line 129")
		return nil, err
	}

	return colls, nil
}

func (sql3 *Sqlite3) AddCollection() (string, bool) {
	colls, err := sql3.QueryCollections()
	if err != nil {
		return "", false
	}

	isSame := false
	for _, coll := range colls {
		if len(coll.Mods) != len(GLOBALAPP.DI.Task.Selected) {
			continue
		}

		for idx, el := range coll.Mods {
			_, exists := GLOBALAPP.DI.Task.Selected[el]
			if !exists {
				break
			}

			if idx+1 == len(coll.Mods) {
				isSame = true
			}
		}

		if isSame {
			break
		}
	}

	if isSame {
		return "Mods you picked has the same in the other collections.", false
	}

	selectedSlice := []string{}

	for k := range GLOBALAPP.DI.Task.Selected {
		selectedSlice = append(selectedSlice, k)
	}

	collectionID := strconv.FormatInt(time.Now().UnixMilli(), 10)

	err = sql3.SDB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&schema.Collection{}).Create(&schema.Collection{
			ID:   collectionID,
			Name: GLOBALAPP.ComponentStatus.PageTools.CollectionName,
			Mods: selectedSlice,
		}).Error; err != nil {
			return err
		}

		if GLOBALAPP.ComponentStatus.PageTools.ApplyCollection {
			if err = tx.Model(&schema.AppConfig{}).
				Where("id = ?", 1).
				Update("active_collection_id", collectionID).Error; err != nil {
				return err
			}

			GLOBALAPP.AppConfig.ActiveCollectionID = collectionID
		}

		return nil
	})
	if err != nil {
		ToastFailed("", "Create collection failed.")
		return "", false
	}

	GLOBALAPP.ComponentStatus.PageTools.ApplyCollection = false
	GLOBALAPP.ComponentStatus.PageTools.CollectionName = ""
	clear(GLOBALAPP.DI.Task.Selected)
	return "Create collection Successful.", true
}

func (sql3 *Sqlite3) DeleteCollections() {
}
