package handler

import (
	"goTodo/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()

	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/sign-up", h.signUp)
		auth.Post("/sign-in", h.signIn)
	})

	router.Route("/api", func(api chi.Router) {

		api.Route("/lists", func(lists chi.Router) {
			lists.Post("/", h.createList)
			lists.Get("/", h.getAllLists)
			lists.Get("/{id}", h.getListById)
			lists.Put("/{id}", h.updateList)
			lists.Delete("/{id}", h.deleteList)

			lists.Route("/{id}/items", func(items chi.Router) {
				items.Post("/", h.createItem)
				items.Get("/", h.getAllItems)
				items.Get("/{item_id}", h.getItemById)
				items.Put("/{item_id}", h.updateItem)
				items.Delete("/{item_id}", h.deleteItem)
			})
		})
	})

	return router
}
