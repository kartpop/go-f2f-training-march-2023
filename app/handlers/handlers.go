// handlers ...
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/rahulgopher/service/business/mid"
	"github.com/rahulgopher/service/foundation/web"
)

// API ....
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness/:id", check.readiness)

	return app
}

// 1. Build entire new mux - not feasible
// 2. Extend existing piece- returning existing web app reference
