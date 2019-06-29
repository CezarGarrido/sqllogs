package sqllogs

import (
	"database/sql/driver"
)

type loggingTx struct {
	wrappedTx driver.Tx
}
func (tx *loggingTx) Commit() error {
	if err := tx.wrappedTx.Commit(); err != nil {
		return err
	}
	return nil
}
func (tx *loggingTx) Rollback() error {
	if err := tx.wrappedTx.Rollback(); err != nil {
		return err
	}
	return nil
}
