package conditions

import (
	"fmt"
	"gorm.io/gorm/clause"
)

type Compare struct {
	Condition
}

func NewCompare() Compare {
	return Compare{Condition{expressions: []clause.Interface{}}}
}

func (c Compare) Eq(column string, value interface{}) Compare {
	return c.addCondition(column, EQ, value)
}

func (c Compare) Ne(column string, value interface{}) Compare {
	return c.addCondition(column, NE, value)
}
func (c Compare) Gt(column string, value interface{}) Compare {
	return c.addCondition(column, GT, value)
}
func (c Compare) Ge(column string, value interface{}) Compare {
	return c.addCondition(column, GE, value)
}
func (c Compare) Lt(column string, value interface{}) Compare {
	return c.addCondition(column, LT, value)
}
func (c Compare) Le(column string, value interface{}) Compare {
	return c.addCondition(column, LE, value)
}
func (c Compare) Between(column string, val1 interface{}, val2 interface{}) Compare {

	return c
}

func (c Compare) addCondition(column string, keyword Keyword, value interface{}) Compare {
	where := clause.Where{}
	where.Exprs = []clause.Expression{clause.Expr{SQL: fmt.Sprintf("%s %s ?", column, keyword), Vars: []interface{}{value}}}
	c.expressions = append(c.expressions, where)
	return c
}

func (c Compare) GetExpressions() []clause.Interface {
	return c.expressions
}
