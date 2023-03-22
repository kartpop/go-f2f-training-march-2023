package web

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/google/uuid"

	"github.com/dimfeld/httptreemux/v5"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

// App ...
type App struct {
	*httptreemux.ContextMux // compositon
	shutdown                chan os.Signal
	mw                      []Middleware
}

// Handler
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {

	app := App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}

	return &app
}

// overriding existing Handle method
func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {

	// First wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// Add the application's general middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {

		//before
		// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: uuid.New().String(), // replace with opentelemetry later
			Now:     time.Now(),
		}

		ctx := context.WithValue(r.Context(), KeyValues, &v)
		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown() // here we can't log as foundation so better to just shutdown
			return
		}

		//after
	}

	a.ContextMux.Handle(method, path, h)
}

// SignalShutdown - graceful shutdown
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
