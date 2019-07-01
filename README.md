# sqllogs 

Go package for parsing SQL queries and Log Queries.

## Getting Started

### Dependencies

* Golang (https://golang.org/dl/).

### Installing

```
   go get github.com/CezarGarrido/sqllogs

```
## Usage

```go
package main

import (
	"database/sql"
	"fmt"
	"time"
	"github.com/CezarGarrido/sqllogs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"username",
		"password",
		"hostname",
		"port",
		"dbname",
	)
	sqllogs.SetDebug(true)
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
	log := sqllogs.ExecLogs()
	fmt.Println("All logs ->",log)
	
	//output:
	//sqllog:Exec -> Select * from receitas where esc_id=43
    //sqllog:Exec -> Select * from categorias where esc_id=43
    //sqllog:Exec -> Select * from despesas where esc_id=43
    //All logs -> [Select * from receitas where esc_id=43 Select * from categorias where esc_id=43 Select * from despesas where esc_id=43]

}

```
## Authors
Cezar Garrido Britez  
[@CezarCgb18](https://twitter.com/CezarCgb18)

