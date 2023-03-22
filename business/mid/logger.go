package mid

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/rahulgopher/service/foundation/web"
)

// Logger

func Logger(log *log.Logger) web.Middleware {

	m := func(before web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			log.Printf("%s : started 	 : %s %s -> %s",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr)

			err := before(ctx, w, r)

			log.Printf("%s : completed(%d) 	 : %s %s -> %s (%s)",
				v.TraceID, v.StatusCode,
				r.Method, r.URL.Path,
				r.RemoteAddr, time.Since(v.Now),
			)

			// Return the error so it can be handled further up the chain.
			return err

		}
		return h
	}

	return m
}

// Why logging :
// Step 1: Take handler and return handler
// step 2: Pass logger middleware to web app initialization
// step 3: log request after handler
