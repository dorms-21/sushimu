package handlers

import (
	"net/http"
	"strconv"

	"restaurante-go/internal/models"
	"restaurante-go/internal/repository"
)

func (a *App) HandleProducts(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "productos")
	if !ok {
		return
	}

	repo := repository.NewProductRepository(a.DB)
	products, err := repo.GetAll()
	if err != nil {
		a.render(w, "productos.html", PageData{
			Title:       "Productos",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudieron cargar los productos",
			Permissions: perms,
			Module:      "productos",
		})
		return
	}

	success := r.URL.Query().Get("success")

	a.render(w, "productos.html", PageData{
		Title:       "Productos",
		UserName:    username,
		UserRole:    role,
		Data:        products,
		Success:     success,
		Permissions: perms,
		Module:      "productos",
	})
}

func (a *App) HandleProductCreate(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "productos")
	if !ok {
		return
	}
	_ = username

	if !perms["create"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
		return
	}

	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))
	active := r.FormValue("active") == "on"

	repo := repository.NewProductRepository(a.DB)
	err := repo.Create(&models.Product{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
		ImageURL:    r.FormValue("image_url"),
		Active:      active,
	})
	if err != nil {
		a.render(w, "productos.html", PageData{
			Title:       "Productos",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudo crear el producto",
			Permissions: perms,
			Module:      "productos",
		})
		return
	}

	http.Redirect(w, r, "/productos?success=Producto creado correctamente", http.StatusSeeOther)
}

func (a *App) HandleProductEdit(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "productos")
	if !ok {
		return
	}
	_ = username

	if !perms["edit"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))
	active := r.FormValue("active") == "on"

	repo := repository.NewProductRepository(a.DB)
	err := repo.Update(&models.Product{
		ID:          id,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
		ImageURL:    r.FormValue("image_url"),
		Active:      active,
	})
	if err != nil {
		a.render(w, "productos.html", PageData{
			Title:       "Productos",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudo actualizar el producto",
			Permissions: perms,
			Module:      "productos",
		})
		return
	}

	http.Redirect(w, r, "/productos?success=Producto actualizado correctamente", http.StatusSeeOther)
}

func (a *App) HandleProductDelete(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "productos")
	if !ok {
		return
	}
	_ = username

	if !perms["delete"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))

	repo := repository.NewProductRepository(a.DB)
	err := repo.Delete(id)
	if err != nil {
		a.render(w, "productos.html", PageData{
			Title:       "Productos",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudo eliminar el producto",
			Permissions: perms,
			Module:      "productos",
		})
		return
	}

	http.Redirect(w, r, "/productos?success=Producto eliminado correctamente", http.StatusSeeOther)
}

func (a *App) HandlePOS(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "pos")
	if !ok {
		return
	}

	a.render(w, "pos.html", PageData{
		Title:       "Punto de venta",
		UserName:    username,
		UserRole:    role,
		Permissions: perms,
		Module:      "pos",
	})
}