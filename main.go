package main

import (
	"log"
	"net/http"

	"github.com/JuanEQuicenoQ/Desarrollo-Final/controllers"
	myhandlers "github.com/JuanEQuicenoQ/Desarrollo-Final/handlers"
	"github.com/JuanEQuicenoQ/Desarrollo-Final/models"
	repositorio "github.com/JuanEQuicenoQ/Desarrollo-Final/repository"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl)
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	db, err := ConectarDB("postgres://rdsbcqik:24AR6Ji_JD5czCO7SFTbr4akyj85WdVK@batyr.db.elephantsql.com/rdsbcqik", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	repo, err := repositorio.NewRepository[models.Libro](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := myhandlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	router := mux.NewRouter()

	router.Handle("/libros", http.HandlerFunc(handler.LeerLibros)).Methods(http.MethodGet)
	router.Handle("/libros", http.HandlerFunc(handler.CrearLibro)).Methods(http.MethodPost)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.LeerUnLibro)).Methods(http.MethodGet)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.ActualizarUnLibro)).Methods(http.MethodPatch)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.EliminarUnLibro)).Methods(http.MethodDelete)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}
