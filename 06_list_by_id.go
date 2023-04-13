package tips

import "github.com/lib/pq"

// Как получить список значений из БД за один запрос ?

func _() {
	// динамическое структурирование  запроса в зависимости от количества ID
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE id = $1 OR id = $2 OR id = $3 ", 1, 2, 3)

	// использование IN/ANY
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE id IN ($1, $2, $3) ", 1, 2, 3)

	// использование специализированных типов релизущующих driver.Value
	arr := pq.Array([]uint32{1, 2, 3})
	_, _ = db.QueryContext(ctx, "SELECT url FROM urls WHERE id IN ($1) ", arr)

}
