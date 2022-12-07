package repository

import (
	"database/sql"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
