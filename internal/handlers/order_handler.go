package handlers

import (
	"net/http"

	"restaurante-go/internal/middleware"
	"restaurante-go/internal/repository"
)

func (a *App) HandleOrders(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		username, role := getSessionData(r)

		repo := repository.NewOrderRepository(a.DB)
		orders, err := repo.GetAll()
		if err != nil {
			a.render(w, "ordenes.html", PageData{
				Title:    "Órdenes",
				UserName: username,
				UserRole: role,
				Error:    "No se pudieron cargar las órdenes",
			})
			return
		}

		a.render(w, "ordenes.html", PageData{
			Title:    "Órdenes",
			UserName: username,
			UserRole: role,
			Data:     orders,
		})
	})
	handler(w, r)
}

func (a *App) HandleSales(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		username, role := getSessionData(r)
		a.render(w, "ventas.html", PageData{
			Title:    "Ventas",
			UserName: username,
			UserRole: role,
		})
	})
	handler(w, r)
}