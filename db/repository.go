package db

import (
	"gorm.io/gorm"
)

// BaseMapper 实现 Mapper
type BaseMapper struct {
	DB *gorm.DB
}

func NewBaseMapper(db *gorm.DB) *BaseMapper {
	return &BaseMapper{DB: db}
}

func (mapper *BaseMapper) Insert(value IModel) (int64, error) {
	result := mapper.DB.Create(value)
	return result.RowsAffected, result.Error
}

func (mapper *BaseMapper) InsertBatch(values interface{}, len int) (int64, error) {
	result := mapper.DB.CreateInBatches(values, len)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) DeleteById(id interface{}, model IModel) (int64, error) {
	result := mapper.DB.Where(model.IdColumnName()+"= ?", id).Delete(model)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) DeleteBatchIds(model IModel, id ...interface{}) (int64, error) {
	result := mapper.DB.Where(model.IdColumnName()+" in (?)", id).Delete(model)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) DeleteByMap(params map[string]interface{}, model IModel) (int64, error) {
	var db *gorm.DB = mapper.DB
	for k, v := range params {
		db = db.Where(k+" = ?", v)
	}
	result := db.Delete(model)
	return result.RowsAffected, result.Error
}

func (mapper *BaseMapper) UpdateById(id interface{}, value IModel) (int64, error) {
	result := mapper.DB.Where(value.IdColumnName()+" = ?", id).Updates(value)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) Update(value IModel, params map[string]interface{}) (int64, error) {
	db := mapper.DB
	for k, v := range params {
		db = db.Where(k+"= ?", v)
	}
	result := db.Updates(value)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) SelectById(id interface{}, model IModel) (int64, error) {
	result := mapper.DB.Where(model.IdColumnName()+" = ?", id).Find(model)
	return result.RowsAffected, result.Error
}
func (mapper *BaseMapper) SelectBatchIds(model interface{}, idName string, id ...interface{}) (int64, error) {
	result := mapper.DB.Where(idName+" in (?)", id).Find(model)
	return result.RowsAffected, result.Error
}

func (mapper *BaseMapper) SelectByMap(model interface{}, params map[string]interface{}) (int64, error) {
	var db = mapper.DB
	for k, v := range params {
		db = db.Where(k+" = ?", v)
	}
	result := db.Find(model)
	return result.RowsAffected, result.Error
}
