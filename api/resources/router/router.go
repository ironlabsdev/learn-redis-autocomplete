package router

import (
	"net/http"

	oauthMiddleware "autocomplete/api/middleware"
	"autocomplete/api/requestlog"
	"autocomplete/api/resources/autocomplete"
	db "autocomplete/database/generated"
	"autocomplete/internal/database"
	"autocomplete/utils/env"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Controller struct {
	Router *chi.Mux

	Conf   *env.Conf
	Logger *zerolog.Logger

	Pool    *pgxpool.Pool
	Queries *db.Queries

	RedisClient *database.RedisClient
}

func (c *Controller) RegisterUses() {
	c.Router.Use(oauthMiddleware.RequestID)
	c.Router.Use(oauthMiddleware.SetEnvConfig)
	c.Router.Use(middleware.Logger)
}

func (c *Controller) RegisterRoutes() {
	c.Router.Route("/api/v1", func(r chi.Router) {
		autocompleteHandler := autocomplete.New(c.Logger, c.Queries, c.RedisClient)
		r.Method(http.MethodGet, "/books/autocomplete", requestlog.NewHandler(autocompleteHandler.Autocomplete, c.Logger))
	})
}
