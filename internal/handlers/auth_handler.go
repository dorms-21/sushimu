package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"restaurante-go/internal/repository"
)

func (a *App) HandleLoginPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	repo := repository.NewUserRepository(a.DB)
	user, err := repo.FindByUsername(username)
	if err != nil {
		a.render(w, "login.html", PageData{
			Title: "Login",
			Error: "Usuario o contraseña incorrectos",
		})
		return
	}

	if !user.Active {
		a.render(w, "login.html", PageData{
			Title: "Login",
			Error: "Tu usuario está inactivo",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		a.render(w, "login.html", PageData{
			Title: "Login",
			Error: "Usuario o contraseña incorrectos",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_user",
		Value:    user.Username,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_role",
		Value:    user.Role,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	switch user.Role {
	case "admin":
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	case "mesero":
		http.Redirect(w, r, "/ordenes", http.StatusSeeOther)
	case "cliente":
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (a *App) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:   "session_role",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}