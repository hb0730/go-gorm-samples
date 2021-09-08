package db

import (
	"fmt"
	"testing"
)

func TestUserService_Insert(t *testing.T) {
	service := NewService()
	user := &User{Name: "ceshi", Age: 20}
	service.Insert(user)
	fmt.Printf("%v", user)
}
