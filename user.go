package main

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

func (u User) CreateTable() {

	// 迁移 schema
	err := db.AutoMigrate(&u)
	if err != nil {
		panic("auto migrate failed:%v")
	}
}

func (u User) Insert() {
	user := User{Name: "test", Age: 18}
	result := db.Create(&user)
	fmt.Printf("受影响的行数: %d\n", result.RowsAffected)
	fmt.Printf("%v\n", user)
}

// InsertBatch 批量插入
func (u User) InsertBatch() {
	names := []string{"小红", "小黑", "小南", "晓楠", "小白", "晶晶", "晓晓", "沐沐", "白小仙"}
	var users []User
	for i := 0; i < 20; i++ {

		users = append(users, User{Name: names[rand.Intn(len(names))], Age: uint(rand.Intn(200))})
	}
	db.CreateInBatches(users, len(users))
}

// FindByName 通过name查询
func (u User) FindByName(name string) []User {
	var user []User
	db.Where(`name like ?`, name+"%").Find(&user)
	return user
}

// Page 根据当前页与大小查询
func (u User) Page(currentPage, countPage int) []User {
	var user []User
	db.Offset(currentPage).Limit(countPage).Find(&user)
	return user
}

// PageCount 统计
func (u User) PageCount() int64 {
	var count int64
	db.Model(&u).Count(&count)
	return count
}

// TxRollback 嵌套事务测试回滚
func (u User) TxRollback(id uint) bool {
	err := db.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", id).Delete(&u)
		return errors.New("test transaction rollback")
	})
	if err != nil {
		return true
	}
	return false

}
func (u User) TxRollback2() {
	err := db.Transaction(func(tx *gorm.DB) error {
		u.Insert1()
		u.Insert2()
		return errors.New("测试事务")
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
}
func (u User) Insert1() {
	user := User{Name: "事务1", Age: 110}
	db.Create(&user)

}
func (u User) Insert2() {
	user := User{Name: "事务2", Age: 110}
	db.Create(&user)
}

type User struct {
	ID   uint   `gorm:"column:id;not null;type:uint;primaryKey;autoIncrement;comment:'id'"`
	Name string `gorm:"column:name;not null;type:varchar(32);comment:'名称';index"`
	Age  uint   `gorm:"column:age;type:uint;comment:'年龄'"`
}

//func (u User) TableName() string {
//
//}
