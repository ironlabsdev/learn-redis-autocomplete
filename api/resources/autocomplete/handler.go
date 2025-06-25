package autocomplete

import (
	"encoding/json"
	"net/http"

	"autocomplete/api/resources/common/errhandler"
	l "autocomplete/api/resources/common/log"
	db "autocomplete/database/generated"
	"autocomplete/internal/database"
	ctxUtil "autocomplete/utils/ctx"
	"github.com/rs/zerolog"
)

type API struct {
	Logger      *zerolog.Logger
	Repository  *db.Queries
	RedisClient *database.RedisClient
}

func New(logger *zerolog.Logger, db *db.Queries, redisClient *database.RedisClient) *API {
	return &API{
		Logger:      logger,
		Repository:  db,
		RedisClient: redisClient,
	}
}

func (a *API) Autocomplete(w http.ResponseWriter, r *http.Request) {
	reqID := ctxUtil.RequestID(r.Context())

	query := r.URL.Query().Get("query")
	if query == "" {
		a.Logger.Info().Str(l.KeyReqID, reqID).Msg("Empty autocomplete query received")
		errhandler.BadRequest(w, errhandler.RespInvalidURLParamID)
		return
	}

	// todo - implement autocomplete search
	var results []string

	if err := json.NewEncoder(w).Encode([]byte(`{"error": "todo - implement autocomplete search"}`)); err != nil {
		a.Logger.Error().Str(l.KeyReqID, reqID).Err(err).Msg("Error encoding autocomplete response")
		errhandler.ServerError(w, errhandler.RespJSONEncodeFailure)
		return
	}

	a.Logger.Info().Str(l.KeyReqID, reqID).Str("query", query).Int("results", len(results)).Msg("Autocomplete query processed")
}
