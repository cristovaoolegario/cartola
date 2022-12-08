package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/db"
	"github.com/go-chi/chi"
)

func ListPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := queries.FindAllPlayers(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(players)
	}
}

func ListMyTeamPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := chi.URLParam(r, "teamID")
		players, err := queries.GetPlayersByMyTeamID(ctx, teamID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(players)
	}
}
