package sqllogs

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	LOGS = make(chan string)
)

func Logs() (string, error) {
	select {
	case msg := <-LOGS:
		return msg, nil
	default:
		return "", errors.New("Failed log mysql")
	}
}

type LoggingDriver struct {
	Driver string
}

func (d LoggingDriver) Open(dsn string) (driver.Conn, error) {
	mysqlDriver := &mysql.MySQLDriver{}
	conn, err := mysqlDriver.Open(dsn)
	if err != nil {
		return nil, err
	}
	con := &loggingConn{wrappedConn: conn}
	return con, nil
}

func init() {
	sql.Register("sqllog:mysql", &LoggingDriver{"mysql"})
}
