package tips

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
	ctx = context.Background()
)

func _() {
	db, err = sql.Open("pg", "postgres://user:password@host:5432/dbname?sslmode=disable&timeout=1s&pool=10&param=valur")
	if err != nil {
		log.Fatalln(err.Error())
	}

	sql.Register("driver", nil) // регистрирует драйвер

	//  интерфесы которые трубуются для реализации своего драйвера
	var _ driver.Driver
	var _ driver.Conn
	var _ driver.Stmt

	// db -  это фасад для работы с пулом соединений

	// Конфигурация
	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(time.Second)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	err = db.Close() // закрывает все соединеняи в пуле, это важно для произвотельности: pgbouncer, pgpool
	if err != nil {
		log.Print(err.Error())
	}
}
