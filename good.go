package main

import (
	"github.com/bitlogic/gzipissue/gzip"
	"github.com/goadesign/goa"
)

// GoodController implements the good resource.
type GoodController struct {
	*goa.Controller
}

// NewGoodController creates a good controller.
func NewGoodController(service *goa.Service) *GoodController {
	c := service.NewController("GoodController")
	c.Use(gzip.Middleware(1))
	return &GoodController{Controller: c}
}
