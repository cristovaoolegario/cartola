package repository

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}

func RunDbInit(relativePath string, db *sql.DB) error {
	path := filepath.Join(relativePath, "sql", "migrations", "20221127164400_init.up.sql")

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	sql := string(c)
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func RegisterRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
