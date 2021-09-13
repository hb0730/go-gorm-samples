package db

import (
	"fmt"
	"testing"
)

func TestBaseMapper_Insert(t *testing.T) {
	value := User{Name: "insert_test", Age: 18}
	mapper := NewBaseMapper(DB)
	_, err := mapper.Insert(&value)
	if err != nil {
		fmt.Printf("Inser error:%v", err)
		return
	}
	fmt.Printf("%v", value)
}

func TestBaseMapper_InsertBatch(t *testing.T) {
	var values []User
	for i := 0; i < 3; i++ {
		values = append(values, User{Name: fmt.Sprintf("Insert_Batch_%d", i), Age: 18 + i})
	}
	mapper := NewBaseMapper(DB)
	_, err := mapper.InsertBatch(values, len(values))
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", values)
}

func TestBaseMapper_DeleteById(t *testing.T) {
	value := User{}
	mapper := NewBaseMapper(DB)
	_, err := mapper.DeleteById(1, &value)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", value)
}

func TestBaseMapper_DeleteBatchIds(t *testing.T) {
	value := User{}
	mapper := NewBaseMapper(DB)
	_, err := mapper.DeleteBatchIds(&value, 2, 3)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", value)
}

func TestBaseMapper_DeleteByMap(t *testing.T) {
	params := map[string]interface{}{"id": 4, "age": 20}
	value := User{}
	mapper := NewBaseMapper(DB)
	_, err := mapper.DeleteByMap(params, &value)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", value)
}

func TestBaseMapper_UpdateById(t *testing.T) {
	user := User{Name: "update_id", Age: 19}
	mapper := NewBaseMapper(DB)
	_, err := mapper.UpdateById(1, &user)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", user)
}

func TestBaseMapper_Update(t *testing.T) {
	user := User{Name: "update_id", Age: 40}
	mapper := NewBaseMapper(DB)
	_, err := mapper.Update(user, map[string]interface{}{"id": 1, "age": 19})
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", user)
}

func TestBaseMapper_SelectById(t *testing.T) {
	value := User{}
	mapper := NewBaseMapper(DB)
	_, err := mapper.SelectById(1, &value)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", value)
}

func TestBaseMapper_SelectBatchIds(t *testing.T) {
	var values []User
	//var values = make([]User, 0)
	mapper := NewBaseMapper(DB)
	_, err := mapper.SelectBatchIds(&values, User{}.IdColumnName(), 1, 2)
	if err != nil {
		fmt.Printf("Insert Batch error:%v", err)
		return
	}
	fmt.Printf("%#v", values)
}
