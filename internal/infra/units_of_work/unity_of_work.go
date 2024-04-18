package units_of_work

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UowInterface interface {
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
	}
}

func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}
	u.Tx = nil
	return nil
}
