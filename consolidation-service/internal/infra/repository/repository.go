package repository

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}

func RunDbInit(db *sql.DB) error {
	path := filepath.Join("../../../sql", "migrations", "20221127164400_init.up.sql")

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
