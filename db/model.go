package db

//
//type Mapper interface {
//	Insert() (int64, error)
//}
//
//type Model struct {
//	Id int `gorm:"column:id;type:int;primaryKey;not null;autoIncrement;comment:'id'"`
//}
//
//func (model *Model) Insert() (int64, error) {
//	result := DB.Create(model)
//	return result.RowsAffected, result.Error
//}
//
//type User struct {
//	Model
//	Name string `gorm:"column:name;type:varchar(32);not null;comment:'名称'"`
//	Age  int    `gorm:"column:age;type:int;not null;comment:'年龄';default:0"`
//}
//
//func init() {
//	err := DB.AutoMigrate(&User{})
//	if err != nil {
//		panic("auto migrate failed")
//	}
//}

type Mapper interface {
	Insert(value interface{}) (int64, error)
}
type Model struct {
	Id int `gorm:"column:id;type:int;primaryKey;not null;autoIncrement;comment:'id'"`
}

func (model *Model) Insert(value interface{}) (int64, error) {
	result := DB.Create(value)
	return result.RowsAffected, result.Error
}

type User struct {
	Model
	Name string `gorm:"column:name;type:varchar(32);not null;comment:'名称'"`
	Age  int    `gorm:"column:age;type:int;not null;comment:'年龄';default:0"`
}
