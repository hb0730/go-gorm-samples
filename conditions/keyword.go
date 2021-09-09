package conditions

type Keyword string

const (
	AND         Keyword = "AND"
	OR          Keyword = "OR"
	NOT         Keyword = "NOT"
	IN          Keyword = "IN"
	NOT_IN      Keyword = "NOT IN"
	LIKE        Keyword = "LIKE"
	NOT_LIKE    Keyword = "NOT LIKE"
	EQ          Keyword = "="
	NE          Keyword = "<>"
	GT          Keyword = ">"
	GE          Keyword = ">="
	LT          Keyword = "<"
	LE          Keyword = "<="
	IS_NULL     Keyword = "IS NULL"
	IS_NOT_NULL Keyword = "IS NOT NULL"
	GROUP_BY    Keyword = "GROUP BY"
	HAVING      Keyword = "HAVING"
	ORDER_BY    Keyword = "ORDER BY"
	EXISTS      Keyword = "EXISTS"
	NOT_EXISTS  Keyword = "NOT EXISTS"
	BETWEEN     Keyword = "BETWEEN"
	NOT_BETWEEN Keyword = "NOT BETWEEN"
	ASC         Keyword = "ASC"
	DESC        Keyword = "DESC"
)
