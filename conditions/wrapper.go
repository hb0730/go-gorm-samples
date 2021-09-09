package conditions

import "gorm.io/gorm/clause"

type Wrapper struct {
	Func
	Compare
}

func NewWrapper() Wrapper {
	w := Wrapper{
		Func:    NewFunc(),
		Compare: NewCompare(),
	}
	return w
}
func (w Wrapper) GetExpressions() []clause.Interface {
	expressions := w.Func.GetExpressions()
	expressions = append(expressions, w.Compare.expressions...)
	return expressions
}
