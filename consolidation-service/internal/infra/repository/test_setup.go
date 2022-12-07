package repository

import (
	"context"
	"database/sql"

	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
	_ "github.com/mattn/go-sqlite3"
)

func SetupTestDb(relativePath string) (context.Context, *sql.DB) {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}

	err = RunDbInit(relativePath, db)

	if err != nil {
		panic(err.Error())
	}

	return ctx, db
}

func SetupTestUoW(ctx context.Context, db *sql.DB) *uow.Uow {
	uow, err := uow.NewUow(ctx, db)
	if err != nil {
		panic(err.Error())
	}
	RegisterRepositories(uow)
	return uow
}
