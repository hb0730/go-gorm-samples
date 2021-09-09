package conditions

import (
	"gorm.io/gorm/clause"
	"strings"
)

type Func struct {
	Condition
}

func NewFunc() Func {
	return Func{Condition{expressions: []clause.Interface{}}}
}

func (f Func) OrderByAsc(column ...string) Func {
	return f.Order(true, column...)
}

func (f Func) OrderByDesc(column ...string) Func {
	return f.Order(false, column...)
}

func (f Func) Order(isAsc bool, column ...string) Func {
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
func (f Func) GetExpressions() []clause.Interface {
	return f.expressions
}
