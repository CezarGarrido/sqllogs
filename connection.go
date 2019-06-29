package sqllogs

import (
	"database/sql/driver"
)

type loggingConn struct {
	wrappedConn driver.Conn
	Queries string
}

func (c *loggingConn) Begin() (driver.Tx, error) {
	tx, err := c.wrappedConn.Begin()
	if err != nil {
		return nil, err
	}
	return &loggingTx{wrappedTx: tx}, nil
}

func (c *loggingConn) Close() error {
	if err := c.wrappedConn.Close(); err != nil {
		return err
	}
	return nil
}

func (c *loggingConn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := c.wrappedConn.Prepare(query)
	if err != nil {
		return nil, err
	}
	loggin:= &loggingStmt{wrappedStmt: stmt, QueryValue: query}
	return loggin, nil
}
