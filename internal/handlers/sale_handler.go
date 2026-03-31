package handlers

import (
	"net/http"
	"strconv"

	"restaurante-go/internal/models"
	"restaurante-go/internal/repository"
)

func (a *App) HandleSales(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "ventas")
	if !ok {
		return
	}

	repo := repository.NewSaleRepository(a.DB)
	items, err := repo.GetAll()
	if err != nil {
		a.render(w, "ventas.html", PageData{
			Title:       "Ventas",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudieron cargar las ventas",
			Permissions: perms,
			Module:      "ventas",
		})
		return
	}

	a.render(w, "ventas.html", PageData{
		Title:       "Ventas",
		UserName:    username,
		UserRole:    role,
		Data:        items,
		Success:     r.URL.Query().Get("success"),
		Permissions: perms,
		Module:      "ventas",
	})
}

func (a *App) HandleSaleCreate(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "ventas")
	if !ok {
		return
	}

	if !perms["create"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ventas", http.StatusSeeOther)
		return
	}

	orderID, _ := strconv.Atoi(r.FormValue("order_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	repo := repository.NewSaleRepository(a.DB)
	err := repo.Create(&models.Sale{
		OrderID:       orderID,
		PaymentMethod: r.FormValue("payment_method"),
		Amount:        amount,
		Status:        r.FormValue("status"),
	})
	if err != nil {
		http.Redirect(w, r, "/ventas?success=Error al crear venta", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ventas?success=Venta creada correctamente", http.StatusSeeOther)
}

func (a *App) HandleSaleEdit(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "ventas")
	if !ok {
		return
	}

	if !perms["edit"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ventas", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	orderID, _ := strconv.Atoi(r.FormValue("order_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	repo := repository.NewSaleRepository(a.DB)
	err := repo.Update(&models.Sale{
		ID:            id,
		OrderID:       orderID,
		PaymentMethod: r.FormValue("payment_method"),
		Amount:        amount,
		Status:        r.FormValue("status"),
	})
	if err != nil {
		http.Redirect(w, r, "/ventas?success=Error al actualizar venta", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ventas?success=Venta actualizada correctamente", http.StatusSeeOther)
}

func (a *App) HandleSaleDelete(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "ventas")
	if !ok {
		return
	}

	if !perms["delete"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/ventas", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	repo := repository.NewSaleRepository(a.DB)
	err := repo.Delete(id)
	if err != nil {
		http.Redirect(w, r, "/ventas?success=Error al eliminar venta", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/ventas?success=Venta eliminada correctamente", http.StatusSeeOther)
}