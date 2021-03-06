package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/tonyalaribe/todoapi/withmgonclosures/internal/config"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo(configuration))
	router.Delete("/{todoID}", DeleteTodo(configuration))
	router.Post("/", CreateTodo(configuration))
	router.Get("/", GetAllTodos(configuration))
	return router
}

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetATodo(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		todos := Todo{
			Slug:  todoID,
			Title: "Hello world from PORT: " + configuration.Constants.PORT,
			Body:  "Heloo world from planet earth",
		}
		render.JSON(w, r, todos) // A chi router helper for serializing and returning json
	}
}

func DeleteTodo(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Deleted TODO successfully"
		render.JSON(w, r, response) // Return some demo response
	}
}

func CreateTodo(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created TODO successfully"
		render.JSON(w, r, response) // Return some demo response
	}
}

func GetAllTodos(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := []Todo{
			{
				Slug:  "slug",
				Title: "Hello world",
				Body:  "Heloo world from planet earth",
			},
		}
		render.JSON(w, r, todos) // A chi router helper for serializing and returning json
	}
}
