package db

import (
	"fmt"
	"testing"
)

func TestUser_Insert(t *testing.T) {
	u := &User{}
	inserU := &User{Name: "测试", Age: 18}
	_, err := u.Insert(inserU)
	if err != nil {
		fmt.Printf("inser failed:%v", err)
	}
	fmt.Printf("insert success %v", inserU)
}
