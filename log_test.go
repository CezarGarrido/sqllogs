package sqllogs

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		"",
		"localhost",
		"3306",
		"financeiro",
	)
	SetDebug(true)
	db, err := sql.Open("sqllog:mysql", dbSource)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("Select * from receitas where esc_id=?", 43)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("Select * from categorias where esc_id=?", 43)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("Select * from despesas where esc_id=?", 43)
	if err != nil {
		panic(err)
	}
	log := ExecLogs()
	fmt.Println("All logs ->",log)

	fmt.Println("Fim.")
}
