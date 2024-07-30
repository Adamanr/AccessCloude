package main

import (
	"accessCloude/internal/config"
	"accessCloude/internal/storage"
	"log"
	"net"
	"net/http"
	"os"

	"log/slog"

	api "accessCloude/internal/handler"

	"github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	cfg, err := config.GetConfigs()
	if err != nil {
		slog.Info("Failed to get config")
		panic(err)
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil
	r := chi.NewRouter()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	log.Println("Database URL: ", dbURL)
	db := storage.NewDatabase(dbURL)

	accessCloude := api.NewAccessCloude(db)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*", "http://localhost:5173"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)
	api.HandlerFromMux(accessCloude, r)

	h := httpmiddleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort(cfg.CS.Host, cfg.CS.Port),
	}

	slog.Info("Starting server on", cfg.CS.Host, cfg.CS.Port)
	log.Fatal(s.ListenAndServe())
}
