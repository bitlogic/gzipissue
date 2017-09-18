//go:generate goagen bootstrap -d github.com/bitlogic/gzipissue/design

package main

import (
	"github.com/bitlogic/gzipissue/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("test")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bad" controller
	c := NewBadController(service)
	app.MountBadController(service, c)
	// Mount "good" controller
	c2 := NewGoodController(service)
	app.MountGoodController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
