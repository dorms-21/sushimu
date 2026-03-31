package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"restaurante-go/internal/middleware"
)

type App struct {
	DB        *sql.DB
	Templates *template.Template
}

type PageData struct {
	Title    string
	UserName string
	UserRole string
	Data     any
	Error    string
}

func NewApp(db *sql.DB, templates *template.Template) *App {
	return &App{
		DB:        db,
		Templates: templates,
	}
}

func (a *App) render(w http.ResponseWriter, page string, data PageData) {
	err := a.Templates.ExecuteTemplate(w, page, data)
	if err != nil {
		http.Error(w, "error renderizando template: "+err.Error(), http.StatusInternalServerError)
	}
}

func getSessionData(r *http.Request) (string, string) {
	var username string
	var role string

	if c, err := r.Cookie("session_user"); err == nil {
		username = c.Value
	}

	if c, err := r.Cookie("session_role"); err == nil {
		role = c.Value
	}

	return username, role
}

func (a *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (a *App) HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	a.render(w, "login.html", PageData{
		Title: "Login",
	})
}

func (a *App) HandleUnauthorized(w http.ResponseWriter, r *http.Request) {
	a.render(w, "unauthorized.html", PageData{
		Title: "Sin autorización",
	})
}

func (a *App) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		username, role := getSessionData(r)
		a.render(w, "dashboard.html", PageData{
			Title:    "Dashboard",
			UserName: username,
			UserRole: role,
		})
	})
	handler(w, r)
}

func (a *App) HandleAdmin(w http.ResponseWriter, r *http.Request) {
	handler := middleware.RequireAuth(
		middleware.RequireRole("admin")(func(w http.ResponseWriter, r *http.Request) {
			username, role := getSessionData(r)
			a.render(w, "admin.html", PageData{
				Title:    "Administración",
				UserName: username,
				UserRole: role,
			})
		}),
	)
	handler(w, r)
}