package main

import (
	"context"
	"database/sql"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}
	repository.RegisterRepositories(uow)
}
