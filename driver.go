package sqllogs

import (
	"database/sql"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

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
