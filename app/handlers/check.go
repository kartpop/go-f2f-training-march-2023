package handlers

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"

	"github.com/rahulgopher/service/foundation/web"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// decode demo later
	if n := rand.Intn(10); n%2 == 0 {
		//return errors.New("untrusted error")
		return web.NewRequestError(errors.New("untrusted error"), 200)
		//panic("forcefully")
	}

	// http://localhost:3000/readiness/10
	m := web.Params(r)
	id := m["id"]

	return web.Respond(ctx, w, map[string]string{"id": id}, http.StatusOK) // internally	json.NewEncoder(w).Encode(status)
}
