# sqllogs 

Go package for parsing SQL queries.

## Getting Started

### Dependencies

* Golang (https://golang.org/dl/).

### Installing

```
   go get https://github.com/CezarGarrido/sqllogs

```
### Usage

```
import (
	"fmt"
	"time"
    "github.com/CezarGarrido/sqllogs"
)
	birty := time.Now()
	name := "teste"
	age := 21
	height := 1.71
	result1 := Query("insert teste set name=?, age=?,birty=?,height=?", name, age, birty, height)
	fmt.Printf("%+v\n", result1)
```
## Authors
Cezar Garrido Britez  
[@CezarCgb18](https://twitter.com/CezarCgb18)

