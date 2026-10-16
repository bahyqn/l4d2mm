package internal

import (
	"database/sql"
	"fmt"
	"l4d2mm/internal/schema"
	"os"
	"path/filepath"

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
	err := sql3.SDB.Model(&schema.AppConfig{}).Where("id=?", 1).Scan(&tcfg).Error

	if err == nil && tcfg.Id == 0 {
		config := DefaultAppConfig()
		sql3.SDB.Model(&schema.AppConfig{}).Create(&config)
	}

	sql3.SDB.Model(&schema.AppConfig{}).Where("id=?", 1).Scan(&sql3.AppDi.AppConfig)
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
