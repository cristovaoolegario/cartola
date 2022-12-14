package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
	httphandler "github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/http"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/kafka/consumer"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}
	repository.RegisterRepositories(uow)

	router := chi.NewRouter()
	router.Get("/my-teams/{teamID}/players", httphandler.ListMyTeamPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/players", httphandler.ListPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/my-teams/{teamID}/balance", httphandler.GetMyTeamBalanceHandler(ctx, *db.New(dtb)))
	router.Get("/matches", httphandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dtb)))
	router.Get("/matches/{matchID}", httphandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dtb)))

	go http.ListenAndServe(":8080", router)

	msgChan := make(chan *kafka.Message)
	var topics = []string{"newMatch", "chooseTeam", "newPlayer", "matchUpdateResult", "newAction"}
	go consumer.Consume(topics, "broker:9094", msgChan)
	consumer.ProcessEvents(ctx, msgChan, uow)
}
