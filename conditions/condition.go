package conditions

import "gorm.io/gorm/clause"

type ICondition interface {
	GetExpressions() []clause.Interface
}
type Condition struct {
	expressions []clause.Interface
}
