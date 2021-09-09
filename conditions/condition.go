package conditions

import "gorm.io/gorm/clause"

type Condition struct {
	expressions []clause.Interface
}
type ICondition interface {
	GetExpressions() []clause.Interface
}

func (wrapper Wrapper) GetExpressions() []clause.Interface {
	return wrapper.expressions
}
