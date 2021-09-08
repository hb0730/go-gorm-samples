package db

import "gorm.io/gorm"

type Mapper interface {
	Insert(value interface{}, db *gorm.DB) (int64, error)
	InsertBatch(values []interface{}, db *gorm.DB) (int64, error)
	DeleteById(id int, model interface{}, db *gorm.DB) (int64, error)
}
type BaseMapper struct {
}

func (mapper *BaseMapper) Insert(value interface{}, db *gorm.DB) (int64, error) {
	result := db.Create(value)
	return result.RowsAffected, result.Error
}

func (mapper *BaseMapper) InsertBatch(values []interface{}, db *gorm.DB) (int64, error) {
	result := db.CreateInBatches(values, len(values))
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) DeleteById(id int, model interface{}, db *gorm.DB) (int64, error) {
	result := db.Where(`id = ?`, id).Delete(model)
	return result.RowsAffected, result.Error
}
