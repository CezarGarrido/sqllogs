package sqllogs

import (
	"database/sql/driver"
	//"github.com/CezarGarrido/sqllogs"
)

type loggingStmt struct {
	wrappedStmt driver.Stmt
	QueryValue  string
}

func (s *loggingStmt) Close() error {
	if err := s.wrappedStmt.Close(); err != nil {
		return err
	}
	return nil
}
func (s *loggingStmt) Exec(args []driver.Value) (driver.Result, error) {
	go func() {
		result := Parse(s.QueryValue, args)
		LOGS <- result
	}()
	result, err := s.wrappedStmt.Exec(args)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *loggingStmt) NumInput() int {
	numInput := s.wrappedStmt.NumInput()
	return numInput
}
func (s *loggingStmt) Query(args []driver.Value) (driver.Rows, error) {
	go func() {
		result := Parse(s.QueryValue, args)
		LOGS <- result
	}()
	rows, err := s.wrappedStmt.Query(args)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
