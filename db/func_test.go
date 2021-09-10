package db

import (
	"fmt"
	"go-gorm-samples/conditions"
	"testing"
)

func TestFunc_IsNull(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().IsNull("id")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestFunc_IsNotNull(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().IsNotNull("id")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestFunc_In(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().In("id", 1, 2)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestFunc_NotIn(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().NotIn("id", 1, 2)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestFunc_GroupBy(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().GroupBy("name")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestFunc_OrderByAsc(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().OrderByAsc("id")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestFunc_OrderByDesc(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = conditions.NewFunc().OrderByDesc("id")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestFunc_OrderBy(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.OrderBy(false, "name", "id")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
