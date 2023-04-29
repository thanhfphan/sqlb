package sqlb

import (
	"fmt"
	"testing"
)

func Test_Example(t *testing.T) {
	sb := NewSelectBuilder()
	sb.Select("id").
		From("student s").
		Where(
			sb.NotIn("s.name", "long", "thanh", "thao"),
		).GroupBy("s.id", "s.name").
		Having(
			sb.LessEqualThan("s.id", 100),
		).
		OrderBy("s.id").Asc()

	fmt.Println(sb.String())

	t.Error("trigger")
}
