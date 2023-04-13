package tips

import (
	"encoding/json"
	"fmt"
	"log"
)

func _() {
	// оставлены для совместимости со старым кодом, нужно использовать методы с context.Context
	_ = db.QueryRow
	_ = db.Query

	var v string
	// "ситаксический сахар" для получения 1 записи
	row := db.QueryRowContext(ctx, "SELECT url FROM urls LIMIT 1")
	// row - всегда != nil
	// row.Scan - возвращет ошибку если значений нет
	_ = row.Scan(&v)

	rows, err := db.QueryContext(ctx, "SELECT url FROM urls")
	if err != nil {
		log.Println(err.Error())
	}

	var urls []string
	for rows.Next() {
		if err := rows.Scan(&v); err != nil {
			log.Println(err)
		}
		urls = append(urls, v)
	}

	_ = rows.Close() // нужно вызывать если только вычитались не все данные или произошла ошибка
}

func _() {
	id := "ID"
	_ = db.QueryRowContext(ctx, "SELECT url FROM urls WHERE id = "+id)

	_ = db.QueryRowContext(ctx, "SELECT url FROM urls WHERE id = $1", id) // аргументы передаются отдельно от кода
}

// {"city":"London", "country":"UK", "phone":["123", "456", "789"]}
type address struct {
	City    string   `json:"city"`
	Country uint32   `json:"country"`
	Phone   []string `json:"phone"`
}

func _() {
	var a address
	if err := db.QueryRowContext(ctx, "SELECT address FROM persons WHERE id = 100").Scan(&a); err != nil {
		log.Println(err.Error())
	}
}

// Scan релизует интерфейс sql.Scanner
// это позволяет преобразовывать любрые данные из базы
func (a *address) Scan(i any) error {
	data, ok := i.([]byte)
	if !ok {
		return fmt.Errorf("unexpected type %T", i)
	}

	if err := json.Unmarshal(data, a); err != nil {
		return fmt.Errorf("unvalid json: %w", err)
	}
	return nil
}
