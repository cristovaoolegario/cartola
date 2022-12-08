package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
)

func ListPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := queries.FindAllPlayers(ctx)
		if err != nil {
			fmt.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(players)
	}
}
