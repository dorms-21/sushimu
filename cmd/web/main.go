package main

import (
	"log"
	"net/http"
	"os"

	"restaurante-go/internal/db"
	"restaurante-go/internal/handlers"
	"restaurante-go/internal/render"
)

func main() {
	// Conexión a la BD
	conn, err := db.Open()
	if err != nil {
		log.Fatal("error conectando a postgres: ", err)
	}
	defer conn.Close()

	// Templates
	tmpl, err := render.LoadTemplates()
	if err != nil {
		log.Fatal("error cargando templates: ", err)
	}

	// App
	app := handlers.NewApp(conn, tmpl)

	mux := http.NewServeMux()

	// Archivos estáticos
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Públicas
	mux.HandleFunc("/", app.HandleHome)
	mux.HandleFunc("/login", app.HandleLoginPage)
	mux.HandleFunc("/auth/login", app.HandleLoginPost)
	mux.HandleFunc("/logout", app.HandleLogout)
	mux.HandleFunc("/unauthorized", app.HandleUnauthorized)

	// Protegidas
	mux.HandleFunc("/dashboard", app.HandleDashboard)
	mux.HandleFunc("/productos", app.HandleProducts)
	mux.HandleFunc("/ordenes", app.HandleOrders)
	mux.HandleFunc("/admin", app.HandleAdmin)
	mux.HandleFunc("/ventas", app.HandleSales)
	mux.HandleFunc("/pos", app.HandlePOS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("servidor corriendo en http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}