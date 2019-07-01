package sqllogs

import (
	"database/sql/driver"
	"fmt"
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
		q := FormatSQL(s.QueryValue, args)
		if DEBUG {
			fmt.Println("sqllog:Exec ->", q)
		}
		ExecLog <- q
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
		q := FormatSQL(s.QueryValue, args)
		if DEBUG {
			fmt.Println("sqllog:Query ->", q)
		}
		QueryLog <- q
	}()
	rows, err := s.wrappedStmt.Query(args)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
