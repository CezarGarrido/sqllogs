package sqllogs

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

var (
	LOGSEXEC  = []string{}
	LOGSQUERY = []string{}
	DEBUG     = false
)

func SetDebug(isDebug bool) {
	DEBUG = isDebug
}
func ExecLogs() []string {
	var logs []string
	logs = LOGSEXEC
	LOGSEXEC = nil
	return logs
}

func QueryLogs() []string {
	var logs []string
	logs = LOGSQUERY
	LOGSQUERY = nil
	return logs
}

func AddQueryLog(log string) {
	LOGSQUERY = append(LOGSQUERY, log)
}

func AddExecLog(log string) {
	LOGSEXEC = append(LOGSEXEC, log)
}
//Query:
//https://groups.google.com/forum/#!topic/golang-nuts/zIwClvZFWIs
func FormatSQL(query string, args []driver.Value) string {
	var buffer bytes.Buffer
	nArgs := len(args)
	for i, part := range strings.Split(query, "?") {
		buffer.WriteString(part)
		if i < nArgs {
			switch a := args[i].(type) {
			case *int64:
				val := args[i]
				if val.(*int64) != nil {
					buffer.WriteString(fmt.Sprintf("%d", *val.(*int64)))
				} else {
					buffer.WriteString("NULL")
				}
			case *int:
				val := args[i]
				if val.(*int) != nil {
					buffer.WriteString(fmt.Sprintf("%d", *val.(*int)))
				} else {
					buffer.WriteString("NULL")
				}
			case *float64:
				val := args[i]
				if val.(*float64) != nil {
					buffer.WriteString(fmt.Sprintf("%f", *val.(*float64)))
				} else {
					buffer.WriteString("NULL")
				}
			case *bool:
				val := args[i]
				if val.(*bool) != nil {
					buffer.WriteString(fmt.Sprintf("%t", *val.(*bool)))
				} else {
					buffer.WriteString("NULL")
				}
			case *string:
				val := args[i]
				if val.(*string) != nil {
					buffer.WriteString(fmt.Sprintf("'%q'", *val.(*string)))
				} else {
					buffer.WriteString("NULL")
				}
			case *time.Time:
				val := args[i]
				if val.(*time.Time) != nil {
					time := *val.(*time.Time)
					buffer.WriteString(fmt.Sprintf("'%v'", time.Format("2006-01-02 15:04:05")))
				} else {
					buffer.WriteString("NULL")
				}
			case int, int8, int16, int32, int64,
				uint, uint8, uint16, uint32, uint64:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case float64:
				buffer.WriteString(fmt.Sprintf("%f", a))
			case bool:
				buffer.WriteString(fmt.Sprintf("%t", a))
			case time.Time:
				buffer.WriteString(fmt.Sprintf("'%v'", a.Format("2006-01-02 15:04:05")))
			case sql.NullBool:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%t", a.Bool))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullInt64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%d", a.Int64))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullString:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%q", a.String))
				} else {
					buffer.WriteString("NULL")
				}

			case nil:
				buffer.WriteString("NULL")
			default:
				buffer.WriteString(fmt.Sprintf("'%v'", a))
			}
		}
	}
	return buffer.String()
}
