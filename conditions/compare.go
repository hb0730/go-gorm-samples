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

// Eq 等于 =
func (c Compare) Eq(column string, value interface{}) Compare {
	return c.addCondition(column, EQ, value)
}

// Ne 不等于 <>
func (c Compare) Ne(column string, value interface{}) Compare {
	return c.addCondition(column, NE, value)
}

// Gt 大于 >
func (c Compare) Gt(column string, value interface{}) Compare {
	return c.addCondition(column, GT, value)
}

// Ge 大于等于 >=
func (c Compare) Ge(column string, value interface{}) Compare {
	return c.addCondition(column, GE, value)
}

// Lt 小于 <
func (c Compare) Lt(column string, value interface{}) Compare {
	return c.addCondition(column, LT, value)
}

// Le 小于等于 <=
func (c Compare) Le(column string, value interface{}) Compare {
	return c.addCondition(column, LE, value)
}

// Between 值1 AND 值2
func (c Compare) Between(column string, val1 interface{}, val2 interface{}) Compare {
	return c.addConditionSql(fmt.Sprintf("%s %s ? %s ?", column, BETWEEN, AND), val1, val2)
}

// NotBetween NOT BETWEEN 值1 AND 值2
func (c Compare) NotBetween(column string, val1 interface{}, val2 interface{}) Compare {
	return c.addConditionSql(fmt.Sprintf("%s %s ? %s ?", column, NOT_BETWEEN, AND), val1, val2)
}

// Like like 值
func (c Compare) Like(column string, value interface{}) Compare {
	return c.addCondition(column, LIKE, value)
}

// NotLike not like 值
func (c Compare) NotLike(column string, value interface{}) Compare {
	return c.addCondition(column, NOT_LIKE, value)
}

func (c Compare) addCondition(column string, keyword Keyword, value interface{}) Compare {
	return c.addConditionSql(fmt.Sprintf("%s %s ?", column, keyword), value)
}

func (c Compare) addConditionSql(sql string, value ...interface{}) Compare {
	where := clause.Where{}
	where.Exprs = []clause.Expression{clause.Expr{SQL: sql, Vars: value}}
	c.expressions = append(c.expressions, where)
	return c
}

func (c Compare) GetExpressions() []clause.Interface {
	return c.expressions
}
