package db

import (
	"fmt"
	"go-gorm-samples/conditions"
	"testing"
)

func TestCompare_Eq(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Eq("id", 1)
	var user User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestCompare_Ne(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Ne("id", 1)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestCompare_Gt(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Gt("id", 1)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestCompare_Ge(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Ge("id", 3)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestCompare_Lt(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Lt("id", 3)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestCompare_Le(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Le("id", 2)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestCompare_Between(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Between("id", 1, 2)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
func TestCompare_NotBetween(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.NotBetween("id", 1, 2)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestCompare_Like(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.Like("name", "%测%")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestCompare_NotLike(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.NotLike("name", "%2")
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}

func TestCompare_Where(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Compare = wrapper.NotLike("name", "%2").Gt("age", 20)
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
