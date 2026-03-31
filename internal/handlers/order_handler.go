package handlers

import (
	"net/http"
	"strconv"

	"restaurante-go/internal/models"
	"restaurante-go/internal/repository"
)

func (a *App) HandleOrders(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "ordenes")
	if !ok {
		return
	}

	repo := repository.NewOrderRepository(a.DB)
	items, err := repo.GetAll()
	if err != nil {
		a.render(w, "ordenes.html", PageData{
			Title:       "Órdenes",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudieron cargar las órdenes",
			Permissions: perms,
			Module:      "ordenes",
		})
		return
	}

	a.render(w, "ordenes.html", PageData{
		Title:       "Órdenes",
		UserName:    username,
		UserRole:    role,
		Data:        items,
		Success:     r.URL.Query().Get("success"),
		Permissions: perms,
		Module:      "ordenes",
	})
}

func (a *App) HandleOrderCreate(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "ordenes")
	if !ok {
		return
	}
	_ = username
	_ = role

	if !perms["create"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ordenes", http.StatusSeeOther)
		return
	}

	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	total, _ := strconv.ParseFloat(r.FormValue("total"), 64)

	repo := repository.NewOrderRepository(a.DB)
	err := repo.Create(&models.Order{
		TableNo: r.FormValue("table_no"),
		UserID:  userID,
		Status:  r.FormValue("status"),
		Total:   total,
	})
	if err != nil {
		http.Redirect(w, r, "/ordenes?success=Error al crear orden", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ordenes?success=Orden creada correctamente", http.StatusSeeOther)
}

func (a *App) HandleOrderEdit(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "ordenes")
	if !ok {
		return
	}
	_ = username
	_ = role

	if !perms["edit"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ordenes", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	userID, _ := strconv.Atoi(r.FormValue("user_id"))
	total, _ := strconv.ParseFloat(r.FormValue("total"), 64)

	repo := repository.NewOrderRepository(a.DB)
	err := repo.Update(&models.Order{
		ID:      id,
		TableNo: r.FormValue("table_no"),
		UserID:  userID,
		Status:  r.FormValue("status"),
		Total:   total,
	})
	if err != nil {
		http.Redirect(w, r, "/ordenes?success=Error al actualizar orden", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ordenes?success=Orden actualizada correctamente", http.StatusSeeOther)
}

func (a *App) HandleOrderDelete(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "ordenes")
	if !ok {
		return
	}

	if !perms["delete"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ordenes", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	repo := repository.NewOrderRepository(a.DB)
	err := repo.Delete(id)
	if err != nil {
		http.Redirect(w, r, "/ordenes?success=Error al eliminar orden", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ordenes?success=Orden eliminada correctamente", http.StatusSeeOther)
}