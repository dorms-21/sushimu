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
	conn, err := db.Open()
	if err != nil {
		log.Fatal("error conectando a postgres: ", err)
	}
	defer conn.Close()

	tmpl, err := render.LoadTemplates()
	if err != nil {
		log.Fatal("error cargando templates: ", err)
	}

	app := handlers.NewApp(conn, tmpl)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", app.HandleHome)
	mux.HandleFunc("/login", app.HandleLoginPage)
	mux.HandleFunc("/auth/login", app.HandleLoginPost)
	mux.HandleFunc("/logout", app.HandleLogout)
	mux.HandleFunc("/unauthorized", app.HandleUnauthorized)

	mux.HandleFunc("/dashboard", app.HandleDashboard)

	mux.HandleFunc("/productos", app.HandleProducts)
	mux.HandleFunc("/productos/crear", app.HandleProductCreate)
	mux.HandleFunc("/productos/editar", app.HandleProductEdit)
	mux.HandleFunc("/productos/eliminar", app.HandleProductDelete)

	mux.HandleFunc("/ordenes", app.HandleOrders)
	mux.HandleFunc("/ordenes/crear", app.HandleOrderCreate)
	mux.HandleFunc("/ordenes/editar", app.HandleOrderEdit)
	mux.HandleFunc("/ordenes/eliminar", app.HandleOrderDelete)

	mux.HandleFunc("/ventas", app.HandleSales)
	mux.HandleFunc("/ventas/crear", app.HandleSaleCreate)
	mux.HandleFunc("/ventas/editar", app.HandleSaleEdit)
	mux.HandleFunc("/ventas/eliminar", app.HandleSaleDelete)

	mux.HandleFunc("/admin", app.HandleAdminUsers)
	mux.HandleFunc("/admin/usuarios/crear", app.HandleAdminUserCreate)
	mux.HandleFunc("/admin/usuarios/editar", app.HandleAdminUserEdit)
	mux.HandleFunc("/admin/usuarios/eliminar", app.HandleAdminUserDelete)

	mux.HandleFunc("/pos", app.HandlePOS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("servidor corriendo en http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}