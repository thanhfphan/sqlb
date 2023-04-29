package sqlb

import (
	"strconv"
	"strings"
)

type SelectBuilder struct {
	Cond

	distinct bool

	selectCols []string
	tables     []string
	joinTables []string
	joinExprs  [][]string

	whereExprs  []string
	havingExprs []string
	groupByCols []string
	orderByCols []string

	order  string
	limit  int
	offset int
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		limit:  -1,
		offset: -1,
	}
}

func Select(cols ...string) *SelectBuilder {
	return NewSelectBuilder().Select(cols...)
}

func (sb *SelectBuilder) Select(cols ...string) *SelectBuilder {
	sb.selectCols = cols
	return sb
}

func (sb *SelectBuilder) Distinct() *SelectBuilder {
	sb.distinct = true
	return sb
}

func (sb *SelectBuilder) From(tables ...string) *SelectBuilder {
	sb.tables = tables
	return sb
}

func (sb *SelectBuilder) Join(table string, onExpr ...string) *SelectBuilder {
	sb.joinTables = append(sb.joinTables, table)
	sb.joinExprs = append(sb.joinExprs, onExpr)
	return sb
}

func (sb *SelectBuilder) Where(andExprs ...string) *SelectBuilder {
	sb.whereExprs = append(sb.whereExprs, andExprs...)
	return sb
}

func (sb *SelectBuilder) GroupBy(cols ...string) *SelectBuilder {
	sb.groupByCols = append(sb.groupByCols, cols...)
	return sb
}

func (sb *SelectBuilder) OrderBy(cols ...string) *SelectBuilder {
	sb.orderByCols = append(sb.orderByCols, cols...)
	return sb
}

func (sb *SelectBuilder) Having(exprs ...string) *SelectBuilder {
	sb.havingExprs = append(sb.havingExprs, exprs...)
	return sb
}

func (sb *SelectBuilder) Asc() *SelectBuilder {
	sb.order = "ASC"
	return sb
}

func (sb *SelectBuilder) Desc() *SelectBuilder {
	sb.order = "DESC"
	return sb
}

func (sb *SelectBuilder) Limit(limit int) *SelectBuilder {
	sb.limit = limit
	return sb
}

func (sb *SelectBuilder) Offset(offset int) *SelectBuilder {
	sb.offset = offset
	return sb
}

func (sb *SelectBuilder) String() string {
	s, _ := sb.Build()
	return s
}

func (sb *SelectBuilder) Build() (string, []interface{}) {
	buf := &strings.Builder{}

	buf.WriteString("SELECT ")
	if sb.distinct {
		buf.WriteString(" DISTINCT ")
	}

	buf.WriteString(strings.Join(sb.selectCols, ", "))
	buf.WriteString(" FROM ")
	buf.WriteString(strings.Join(sb.tables, ", "))

	for i, item := range sb.joinTables {
		buf.WriteString(" JOIN ")
		buf.WriteString(item)

		exprs := sb.joinExprs[i]
		if len(exprs) > 0 {
			buf.WriteString(" ON ")
			buf.WriteString(strings.Join(sb.joinExprs[i], " AND "))
		}
	}

	if len(sb.whereExprs) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(sb.whereExprs, " AND "))
	}

	if len(sb.groupByCols) > 0 {
		buf.WriteString(" GROUP BY ")
		buf.WriteString(strings.Join(sb.groupByCols, ", "))
		if len(sb.havingExprs) > 0 {
			buf.WriteString(" HAVING ")
			buf.WriteString(strings.Join(sb.havingExprs, " AND "))
		}
	}

	if len(sb.orderByCols) > 0 {
		buf.WriteString(" ORDER BY ")
		buf.WriteString(strings.Join(sb.orderByCols, ", "))
		if sb.order != "" {
			buf.WriteString(" ")
			buf.WriteString(sb.order)
		}
	}

	if sb.limit > 0 {
		buf.WriteString(" LIMIT ")
		buf.WriteString(strconv.Itoa(sb.limit))
		if sb.offset > 0 {
			buf.WriteString(" OFFSET ")
			buf.WriteString(strconv.Itoa(sb.offset))
		}
	}

	return buf.String(), nil
}
