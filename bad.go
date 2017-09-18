package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/gzip"
)

// BadController implements the bad resource.
type BadController struct {
	*goa.Controller
}

// NewBadController creates a bad controller.
func NewBadController(service *goa.Service) *BadController {
	c := service.NewController("BadController")
	c.Use(gzip.Middleware(1))
	return &BadController{Controller: c}
}
