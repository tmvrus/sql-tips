package tips

import "log"

func _() {
	// запросы который часто выполняется без изменений самого кода
	// компилируется и проверяется один раз
	// привязывается к соединению

	_ = db.Prepare // оставлен для совместисмости

	stm, err := db.PrepareContext(ctx, "SELECT url FROM urls WHERE id = $1 AND user_id = $2 AND is_active = $3")
	if err != nil {
		log.Println(err.Error())
	}

	_, _ = stm.QueryContext(ctx, "url_1", "user_1", false)
	_, _ = stm.QueryContext(ctx, "url_2", "user_2", false)
	_, _ = stm.QueryContext(ctx, "url_3", "user_3", false)

	if err := stm.Close(); err != nil {
		log.Println(err.Error())
	}
}
