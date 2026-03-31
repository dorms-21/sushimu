package handlers

import (
	"net/http"

	"restaurante-go/internal/middleware"
	"restaurante-go/internal/repository"
)

func (a *App) HandleProducts(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		username, role := getSessionData(r)

		repo := repository.NewProductRepository(a.DB)
		products, err := repo.GetAll()
		if err != nil {
			a.render(w, "productos.html", PageData{
				Title:    "Productos",
				UserName: username,
				UserRole: role,
				Error:    "No se pudieron cargar los productos",
			})
			return
		}

		a.render(w, "productos.html", PageData{
			Title:    "Productos",
			UserName: username,
			UserRole: role,
			Data:     products,
		})
	})
	handler(w, r)
}

func (a *App) HandlePOS(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		username, role := getSessionData(r)
		a.render(w, "pos.html", PageData{
			Title:    "Punto de venta",
			UserName: username,
			UserRole: role,
		})
	})
	handler(w, r)
}