package conditions

import (
	"fmt"
	"gorm.io/gorm/clause"
	"strings"
)

type Func struct {
	*Condition
}

func NewFunc() *Func {
	return &Func{&Condition{expressions: []clause.Interface{}}}
}

func (f *Func) IsNull(column string) *Func {
	f.addConditionSql(fmt.Sprintf("%s %s", column, IS_NULL))
	return f
}

func (f *Func) IsNotNull(column string) *Func {
	f.addConditionSql(fmt.Sprintf("%s %s", column, IS_NOT_NULL))
	return f
}
func (f *Func) In(column string, value ...interface{}) *Func {
	f.addConditionSql(fmt.Sprintf("%s %s (?)", column, IN), value)
	return f
}

func (f *Func) NotIn(column string, value ...interface{}) *Func {
	f.addConditionSql(fmt.Sprintf("%s %s (?)", column, NOT_IN), value)
	return f
}

func (f *Func) GroupBy(column ...string) *Func {
	columns := strings.Join(column, ",")
	groupBy := clause.GroupBy{}
	groupBy.Columns = []clause.Column{
		{
			Name: columns,
			Raw:  true,
		},
	}
	f.expressions = append(f.expressions, groupBy)
	return f
}

func (f *Func) OrderByAsc(column ...string) *Func {
	return f.OrderBy(true, column...)
}

func (f *Func) OrderByDesc(column ...string) *Func {
	return f.OrderBy(false, column...)
}

func (f *Func) OrderBy(isAsc bool, column ...string) *Func {
	columns := strings.Join(column, ",")
	orderBy := clause.OrderBy{
		Columns: []clause.OrderByColumn{{
			Column: clause.Column{Name: columns, Raw: true},
			Desc:   !isAsc,
		}},
	}
	f.expressions = append(f.expressions, orderBy)
	return f
}

func (f *Func) Having(sql string, val ...interface{}) *Func {
	groupBy := clause.GroupBy{}
	groupBy.Having = []clause.Expression{clause.Expr{SQL: sql, Vars: val}}
	f.expressions = append(f.expressions, groupBy)
	return f
}

func (f *Func) GetExpressions() []clause.Interface {
	return f.expressions
}
