package conditions

import (
	"fmt"
	"gorm.io/gorm/clause"
)

type Expression struct {
}

func NewExpression() Expression {
	return Expression{}
}

func (e Expression) Eq(column string, value interface{}) clause.Expression {
	return e.addCondition(column, EQ, value)
}

func (e *Expression) addCondition(column string, keyword Keyword, value ...interface{}) clause.Expression {
	return e.addConditionSql(fmt.Sprintf("%s %s ?", column, keyword), value)
}
func (e Expression) addConditionSql(sql string, value []interface{}) clause.Expression {
	return clause.Expression(clause.Expr{SQL: sql, Vars: value, WithoutParentheses: true})
}
