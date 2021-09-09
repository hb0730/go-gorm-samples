package db

import (
	"fmt"
	"go-gorm-samples/conditions"
	"testing"
)

func TestBaseMapper_Insert(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	mapper.Insert(&User{Name: "测试1", Age: 19})
	mapper.Insert(&User{Name: "测试2", Age: 20})
	mapper.Insert(&User{Name: "测试3", Age: 30})
}
func TestBaseMapper_Select(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Func = wrapper.OrderByDesc("age", "name")
	wrapper.Compare = wrapper.Eq("id", 1).Ne("name", "cc")
	users := []User{}
	mapper.Select(wrapper, &users)
	fmt.Printf("%v", users)

}

func TestBaseMapper_Select_between(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Between("id", 1, 3)
	users := []User{}
	mapper.Select(wrapper, &users)
	fmt.Printf("%v", users)
}
