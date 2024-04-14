package repositories

import (
	"database/sql"
	"io"
)

type RepositoryDatabase interface {
	io.Closer
	GetDB() *sql.DB
}
