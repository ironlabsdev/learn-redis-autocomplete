package router

import (
	oauthMiddleware "autocomplete/api/middleware"
	db "autocomplete/database/generated"
	"autocomplete/utils/env"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Controller struct {
	Conf    *env.Conf
	Pool    *pgxpool.Pool
	Router  *chi.Mux
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func (c *Controller) RegisterUses() {
	c.Router.Use(oauthMiddleware.RequestID)
	c.Router.Use(oauthMiddleware.SetEnvConfig)
	c.Router.Use(middleware.Logger)
}

func (c *Controller) RegisterRoutes() {

}
