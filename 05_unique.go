package tips

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

// Как реализовать требование уникальности?

func _() {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		err := tx.Rollback()
		if !errors.Is(err, sql.ErrTxDone) {
			log.Println(err.Error())
		}
	}()

	if err != nil {
		log.Println(err.Error())
	}
	_, err = tx.QueryContext(ctx, "SELECT uniq FROM urls")
	if !errors.Is(err, sql.ErrNoRows) {
		return
	}
	_, _ = tx.ExecContext(ctx, "INSERT uniq INTO urls")
	if err := tx.Commit(); err != nil {
		log.Println(err.Error())
	}
}

func _() {
	_, _ = db.ExecContext(ctx, "ALTER TABLE urls ADD CONSTRAINT unique_constraint_name UNIQUE (url)")

	_, err := db.ExecContext(ctx, "INSERT INTO urls (url) VALUES (?)", "url")
	// ERROR:  duplicate key value violates unique constraint "unique_constraint"
	if err != nil && strings.Contains(err.Error(), "unique_constraint_name") {
		// дубликат
	}

	res, _ := db.ExecContext(ctx, "INSERT INTO urls (url) VALUES (?) ON CONFLICT DO NOTHING", "url")
	row, _ := res.RowsAffected()
	if row < 1 {
		// дубликат
	}
}
