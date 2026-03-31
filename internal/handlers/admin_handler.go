package handlers

import (
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"restaurante-go/internal/models"
	"restaurante-go/internal/repository"
)

func (a *App) HandleAdminUsers(w http.ResponseWriter, r *http.Request) {
	username, role, perms, ok := a.requireModuleAccess(w, r, "admin")
	if !ok {
		return
	}

	repo := repository.NewUserRepository(a.DB)
	items, err := repo.GetAll()
	if err != nil {
		a.render(w, "admin.html", PageData{
			Title:       "Administración",
			UserName:    username,
			UserRole:    role,
			Error:       "No se pudieron cargar los usuarios",
			Permissions: perms,
			Module:      "admin",
		})
		return
	}

	a.render(w, "admin.html", PageData{
		Title:       "Administración",
		UserName:    username,
		UserRole:    role,
		Data:        items,
		Success:     r.URL.Query().Get("success"),
		Permissions: perms,
		Module:      "admin",
	})
}

func (a *App) HandleAdminUserCreate(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "admin")
	if !ok {
		return
	}

	if !perms["create"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		http.Redirect(w, r, "/admin?success=Error al generar contraseña", http.StatusSeeOther)
		return
	}

	active := r.FormValue("active") == "on"

	repo := repository.NewUserRepository(a.DB)
	err = repo.Create(&models.User{
		Username:     r.FormValue("username"),
		PasswordHash: string(hash),
		Name:         r.FormValue("name"),
		Role:         r.FormValue("role"),
		Active:       active,
	})
	if err != nil {
		http.Redirect(w, r, "/admin?success=Error al crear usuario", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin?success=Usuario creado correctamente", http.StatusSeeOther)
}

func (a *App) HandleAdminUserEdit(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "admin")
	if !ok {
		return
	}

	if !perms["edit"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	active := r.FormValue("active") == "on"

	repo := repository.NewUserRepository(a.DB)
	err := repo.Update(&models.User{
		ID:       id,
		Username: r.FormValue("username"),
		Name:     r.FormValue("name"),
		Role:     r.FormValue("role"),
		Active:   active,
	})
	if err != nil {
		http.Redirect(w, r, "/admin?success=Error al actualizar usuario", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin?success=Usuario actualizado correctamente", http.StatusSeeOther)
}

func (a *App) HandleAdminUserDelete(w http.ResponseWriter, r *http.Request) {
	_, _, perms, ok := a.requireModuleAccess(w, r, "admin")
	if !ok {
		return
	}

	if !perms["delete"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	repo := repository.NewUserRepository(a.DB)
	err := repo.Delete(id)
	if err != nil {
		http.Redirect(w, r, "/admin?success=Error al eliminar usuario", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin?success=Usuario eliminado correctamente", http.StatusSeeOther)
}