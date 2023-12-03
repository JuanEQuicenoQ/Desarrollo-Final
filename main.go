package main

import (
	"log"
	"net/http"

	"github.com/JuanEQuicenoQ/Desarrollo-Final/controllers"
	"github.com/JuanEQuicenoQ/Desarrollo-Final/handlers"
	"github.com/JuanEQuicenoQ/Desarrollo-Final/models"
	repositorio "github.com/JuanEQuicenoQ/Desarrollo-Final/repository" /* importando el paquete de repositorio */
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB("aquí_va_la_URL_de_conexión_de_tu_instancia_de_PostgreSQL", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
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

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/libros", http.HandlerFunc(handler.LeerLibros)).Methods(http.MethodGet)
	router.Handle("/libros", http.HandlerFunc(handler.CrearLibro)).Methods(http.MethodPost)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.LeerUnLibro)).Methods(http.MethodGet)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.ActualizarUnLibro)).Methods(http.MethodPatch)
	router.Handle("/libros/{id}", http.HandlerFunc(handler.EliminarUnLibro)).Methods(http.MethodDelete)

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	http.ListenAndServe(":8080", router)
}
