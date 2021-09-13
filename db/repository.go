package db

import (
	"go-gorm-samples/conditions"
	"gorm.io/gorm"
)

type Mapper interface {
	Insert(value interface{}) (int64, error)
	Select(wrapper conditions.Wrapper, value interface{}) (int64, error)
	Find(value IModel) (int64, error)
	FindArray(value interface{}) (int64, error)
}
type BaseMapper struct {
	DB *gorm.DB
}

func (mapper *BaseMapper) Insert(value interface{}) (int64, error) {
	result := mapper.DB.Create(value)
	return result.RowsAffected, result.Error
}

func (mapper BaseMapper) Select(wrapper conditions.Wrapper, value interface{}) (int64, error) {
	expressions := wrapper.GetExpressions()
	for _, expression := range expressions {
		mapper.DB.Statement.AddClause(expression)
	}
	result := mapper.DB.Find(value)
	return result.RowsAffected, result.Error
}
func (mapper BaseMapper) Find(value IModel) (int64, error) {
	result := mapper.DB.Take(value)
	return result.RowsAffected, result.Error
}

func (mapper BaseMapper) FindArray(value interface{}) (int64, error) {
	result := mapper.DB.Find(value)
	return result.RowsAffected, result.Error
}
