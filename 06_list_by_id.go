package tips

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
)

// Как получить список значений из БД за один запрос ?

func builder(ll []int) (string, []any) {
	ss := make([]string, len(ll))
	for i := range ll {
		ss = append(ss, fmt.Sprintf("id = $%d", i+1))
	}

	return strings.Join(ss, " OR "), []any{ll}
}

func _() {
	// динамическое структурирование  запроса в зависимости от количества ID
	q, a := builder([]int{1, 2, 3})
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE"+q, a...)

	// использование IN/ANY
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE id IN ($1, $2, $3) ", 1, 2, 3)

	// использование специализированных типов релизущующих driver.Value
	arr := pq.Array([]uint32{1, 2, 3})
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE id IN ($1) ", arr)

}
