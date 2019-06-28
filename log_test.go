package sqllogs

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	birty := time.Now()
	name := "teste"
	age := 21
	height := 1.71
	result1 := Query("insert teste set name=?, age=?,birty=?,height=?", name, age, birty, height)
	fmt.Printf("%+v\n", result1)
}

func TestSelect(t *testing.T) {
	name := "teste"
	result2 := Query("select * from teste where name=?", name)
	fmt.Printf("%+v\n", result2)
}
func TestPointerInsert(t *testing.T) {
	day:=time.Now()
	birty := &day
	name := "teste"
	age := 21
	height := 1.71
	result1 := Query("insert teste set name=?, age=?,birty=?,height=?", name, age, birty, height)
	fmt.Printf("%+v\n", result1)
}
