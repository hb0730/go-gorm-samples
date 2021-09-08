package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	var (
		err error
	)
	DB, err = gorm.Open(sqlite.Open("./test.db"))
	if err != nil {
		panic("failed to connect database")
	}
	DB.SkipDefaultTransaction = false
	DB.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "sys_",
		SingularTable: true,
		NoLowerCase:   false,
	}
	//DB.Logger=logger.Default
	DB.DisableNestedTransaction = false
	DB.DisableForeignKeyConstraintWhenMigrating = true
	DB = DB.Debug()

	AutoTable()
}

func AutoTable() {
	DB.AutoMigrate(&User{})
}
