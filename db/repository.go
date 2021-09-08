package db

import (
	"errors"
	"gorm.io/gorm"
)

type Mapper interface {
	Insert(value interface{}) (int64, error)
	InsertBatch(values []interface{}) (int64, error)
	DeleteById(id int, model interface{}) (int64, error)
}
type BaseMapper struct {
	DB *gorm.DB
}

func (mapper *BaseMapper) Insert(value interface{}) (int64, error) {
	result := mapper.DB.Create(value)
	return result.RowsAffected, result.Error
}

func (mapper *BaseMapper) InsertBatch(values []interface{}) (int64, error) {
	result := mapper.DB.CreateInBatches(values, len(values))
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) DeleteById(id int, model interface{}) (int64, error) {
	value, ok := model.(IModel)
	if !ok {
		return 0, errors.New("unsupported type")
	}
	result := mapper.DB.Where(value.IdColumnName()+` = ?`, id).Delete(model)
	return result.RowsAffected, result.Error
}
