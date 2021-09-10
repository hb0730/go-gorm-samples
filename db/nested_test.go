package db

import (
	"fmt"
	"go-gorm-samples/conditions"
	"testing"
)

func TestNested_Or(t *testing.T) {
	mapper := BaseMapper{DB: DB}
	wrapper := conditions.NewWrapper()
	wrapper.Eq("id", 2)
	wrapper.Or(conditions.NewExpression().Eq("age", 19))
	var user []User
	mapper.Select(wrapper, &user)
	fmt.Printf("%v", user)
}
