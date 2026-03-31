package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"restaurante-go/internal/repository"
)

type App struct {
	DB        *sql.DB
	Templates *template.Template
}

type PageData struct {
	Title       string
	UserName    string
	UserRole    string
	Data        any
	Error       string
	Success     string
	Permissions map[string]bool
	Module      string
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

func (a *App) getPermissions(role, module string) map[string]bool {
	repo := repository.NewPermissionRepository(a.DB)
	perms, err := repo.GetPermission(role, module)
	if err != nil {
		return map[string]bool{
			"view":   false,
			"create": false,
			"edit":   false,
			"delete": false,
		}
	}
	return perms
}

func (a *App) requireModuleAccess(w http.ResponseWriter, r *http.Request, module string) (string, string, map[string]bool, bool) {
	username, role := getSessionData(r)

	if username == "" || role == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return "", "", nil, false
	}

	perms := a.getPermissions(role, module)
	if !perms["view"] {
		http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
		return "", "", nil, false
	}

	return username, role, perms, true
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
	username, role, perms, ok := a.requireModuleAccess(w, r, "dashboard")
	if !ok {
		return
	}

	a.render(w, "dashboard.html", PageData{
		Title:       "Dashboard",
		UserName:    username,
		UserRole:    role,
		Permissions: perms,
		Module:      "dashboard",
	})
}