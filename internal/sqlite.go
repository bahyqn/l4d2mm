package internal

import (
	"database/sql"
	"fmt"
	"l4d2mm/internal/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sqlite3 struct {
	AppDi *App
	SDB   *gorm.DB
	Conn  *sql.DB
}

func NewSqlite3(di *App) *Sqlite3 {
	return &Sqlite3{
		AppDi: di,
	}
}


func (sql3 *Sqlite3) Connecte(dir string) {
	var err error

	sql3.SDB, err = gorm.Open(sqlite.Open(dir+"/data.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		fmt.Println("err: ", err)
	}

	sql3.Conn, err = sql3.SDB.DB()

	if err != nil {
		fmt.Println("Failed to connect sqlite3: ", err)
	}

	//sql3.FetchTables()
}

func (sql3 *Sqlite3) FetchTables() {
	sql3.SDB.AutoMigrate(&schema.AppConfig{})
	sql3.SDB.AutoMigrate(&schema.Mod{})
	sql3.SDB.AutoMigrate(&schema.Collection{})

	var tcfg schema.AppConfig
	err := sql3.SDB.Model(&schema.AppConfig{}).Where("id=?", 1).Scan(&tcfg).Error

	if err == nil && tcfg.Id == 0 {
		sql3.SDB.Model(&schema.AppConfig{}).Create(DefaultAppConfig())
	}


	sql3.SDB.Model(&schema.AppConfig{}).Where("id=?", 1).Scan(&sql3.AppDi.AppConfig)
}