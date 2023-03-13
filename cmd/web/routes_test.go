package main

import (
	"fmt"
	"testing"

	"github.com/dikyayodihamzah/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		errStr := fmt.Sprintf("type is not *chi.Mux, but is %T", v)
		t.Error(errStr)
	}
}
