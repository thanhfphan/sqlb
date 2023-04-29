package sqlb

import (
	"fmt"
	"strings"
)

type Cond struct {
}

func (c *Cond) Equal(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" = ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) NotEqual(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" <> ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) GreaterThan(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" > ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) GreaterEqualThan(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" >= ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) LessThan(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" < ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) LessEqualThan(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" <= ")
	buf.WriteString(fmt.Sprintf("$%v", value))
	return buf.String()
}

func (c *Cond) In(field string, values ...interface{}) string {
	vals := []string{}
	for _, item := range values {
		vals = append(vals, fmt.Sprintf("$%v", item))
	}
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" IN (")
	buf.WriteString(strings.Join(vals, ", "))
	buf.WriteString(")")
	return buf.String()
}

func (c *Cond) NotIn(field string, values ...interface{}) string {
	vals := []string{}
	for _, item := range values {
		vals = append(vals, fmt.Sprintf("$%v", item))
	}
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" NOT IN (")
	buf.WriteString(strings.Join(vals, ", "))
	buf.WriteString(")")
	return buf.String()
}

func (c *Cond) Like(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" LIKE ")
	buf.WriteString(fmt.Sprintf("#%v", value))
	return buf.String()
}

func (c *Cond) NotLike(field string, value interface{}) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" NOT LIKE ")
	buf.WriteString(fmt.Sprintf("#%v", value))
	return buf.String()
}

func (c *Cond) IsNull(field string) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" IS NULL")
	return buf.String()
}

func (c *Cond) IsNotNull(field string) string {
	buf := &strings.Builder{}
	buf.WriteString(field)
	buf.WriteString(" IS NOT NULL")
	return buf.String()
}

func (c *Cond) Or(orExprs ...string) string {
	buf := &strings.Builder{}
	buf.WriteString("(")
	buf.WriteString(strings.Join(orExprs, " OR "))
	buf.WriteString(")")
	return buf.String()
}

func (c *Cond) And(orExprs ...string) string {
	buf := &strings.Builder{}
	buf.WriteString("(")
	buf.WriteString(strings.Join(orExprs, " AND "))
	buf.WriteString(")")
	return buf.String()
}
