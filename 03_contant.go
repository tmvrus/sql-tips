package tips

func _() {
	// подготовка аргументов
	_, _ = db.QueryContext(ctx, query)
	// работа с ошибками
	// работа с результатом
}

const query = `WITH res AS (
    SELECT 
		money, 
		bonus, 
		balance
    FROM 
		money
    WHERE id = $2
    FOR UPDATE
)
INSET INTO 
	history (tx_id, user, money, bonus, source, reason, currency)
VALUES
	($1, $2, $3, 0, $4, $5, COALESCE((SELECT balance FROM res), 0) + $3, COALESCE((SELECT bonus FROM res), 0))
ON CONFLICT DO NOTHING`
