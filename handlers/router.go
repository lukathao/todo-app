package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			router.Get("/healthcheck", healthCheck)
			router.Get("/todos/{id}", getTodoById)
			router.Get("/todos", getTodos)
			router.Post("/todos/create", createTodo)
			router.Put("/todos/update/{id}", updateTodo)
			router.Delete("/todos/delete/{id}", deleteTodo)
		})
	})

	return router
}
