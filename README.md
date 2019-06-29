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

```
package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
	_ "github.com/CezarGarrido/sqllogs"
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
	db, err := sql.Open("sqllog:mysql", dbSource)
	if err != nil {
		panic(err)
	}
	birty := time.Now()
	name := "teste"
	age := 21
	height := 1.71
	_, err = db.Exec("insert tests set name=?, age=?,birty=?,height=?", name, age, birty, height)
	if err != nil {
		panic(err)
	}
	log, _ := sqllogs.Logs()
	fmt.Println(log)

}

```
## Authors
Cezar Garrido Britez  
[@CezarCgb18](https://twitter.com/CezarCgb18)

