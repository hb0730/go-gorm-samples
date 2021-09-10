package conditions

import "gorm.io/gorm/clause"

type Wrapper struct {
	*Func
	*Compare
	*Nested
}

func NewWrapper() Wrapper {
	w := Wrapper{
		Func:    NewFunc(),
		Compare: NewCompare(),
		Nested:  NewNested(),
	}
	return w
}
func (w Wrapper) GetExpressions() []clause.Interface {
	var expressions []clause.Interface
	expressions = append(expressions, w.Func.GetExpressions()...)
	expressions = append(expressions, w.Compare.GetExpressions()...)
	expressions = append(expressions, w.Nested.GetExpressions()...)
	return expressions
}
