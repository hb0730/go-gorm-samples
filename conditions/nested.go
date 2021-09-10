package conditions

import "gorm.io/gorm/clause"

type Nested struct {
	*Condition
}

func NewNested() *Nested {
	return &Nested{&Condition{expressions: []clause.Interface{}}}
}

func (n *Nested) Or(ef ...clause.Expression) *Nested {
	n.expressions = append(n.expressions,
		clause.Where{Exprs: []clause.Expression{clause.Or(clause.And(ef...))}})
	return n
}

func (n *Nested) GetExpressions() []clause.Interface {
	return n.expressions
}
