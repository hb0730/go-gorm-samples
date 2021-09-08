package db

type IModel interface {
	GetId() int
	IdColumnName() string
}
type Model struct {
	Id int `gorm:"column:id;type:int;primaryKey;not null;autoIncrement;comment:'id'"`
}

func (model Model) GetId() int {
	return model.Id
}
func (model Model) IdColumnName() string {
	return "id"
}

type User struct {
	Model
	Name string `gorm:"column:name;type:varchar(32);not null;comment:'名称'"`
	Age  int    `gorm:"column:age;type:int;not null;comment:'年龄';default:0"`
}
