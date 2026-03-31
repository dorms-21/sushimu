package middleware

import "net/http"

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_user")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func RequireRole(allowedRoles ...string) func(http.HandlerFunc) http.HandlerFunc {
	allowed := make(map[string]bool)
	for _, role := range allowedRoles {
		allowed[role] = true
	}

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			roleCookie, err := r.Cookie("session_role")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			if !allowed[roleCookie.Value] {
				http.Redirect(w, r, "/unauthorized", http.StatusSeeOther)
				return
			}

			next(w, r)
		}
	}
}