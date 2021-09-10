package conditions

import "gorm.io/gorm/clause"

type ICondition interface {
	GetExpressions() []clause.Interface
}
type Condition struct {
	expressions []clause.Interface
}

func (c *Condition) addConditionSql(sql string, value ...interface{}) {
	where := clause.Where{}
	where.Exprs = c.addExpression(sql, value)
	c.expressions = append(c.expressions, where)
}
func (c *Condition) addExpression(sql string, value ...interface{}) []clause.Expression {
	return []clause.Expression{clause.Expr{SQL: sql, Vars: value, WithoutParentheses: true}}
}
