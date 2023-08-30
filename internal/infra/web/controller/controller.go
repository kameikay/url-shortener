package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/kameikay/url-shortener/internal/infra/web/handlers"
)

type Controller struct {
	router  chi.Router
	handler *handlers.Handler
}

func NewController(router chi.Router, handler *handlers.Handler) *Controller {
	return &Controller{
		router:  router,
		handler: handler,
	}
}

func (c *Controller) Route() {
	c.router.Post("/", c.handler.ReturnCodeHandler)
	c.router.Get("/{code}", c.handler.RedirectToUrl)
}
