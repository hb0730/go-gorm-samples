package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func main() {
	userTest()

}
func userTest() {
	user := User{}
	user.TxRollback2()
	//user.Insert()
	//user.CreateTable()
	//user.InsertBatch()
	//users := user.FindByName("晓")
	//fmt.Printf("%#v", users)
	//users :=user.Page(1,10)
	//fmt.Printf("%#v", users)
	//count := user.PageCount()
	//fmt.Println(count)
	//user.TxRollback(1)
}
func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//NamingStrategy 命名规则
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
			NoLowerCase:   false,
		},
		//DisableForeignKeyConstraintWhenMigrating 禁用创建外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
}
