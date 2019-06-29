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

	db, err := sql.Open("sqllog:mysql", dbSource)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("Select * from receitas where esc_id=?", 43)
	if err != nil {
		panic(err)
	}
	log, _ := Logs()
	fmt.Println(log)

}
