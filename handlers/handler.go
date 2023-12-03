package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JuanEQuicenoQ/Desarrollo-Final/controllers"
	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarUnLibro(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un libro, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un libro, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarUnLibro(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un libro, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un libro, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarUnLibro(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.controller.EliminarUnLibro(id)
	if err != nil {
		log.Printf("fallo al eliminar un libro, con error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("fallo al eliminar un libro con id %s", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) LeerUnLibro(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	libro, err := h.controller.LeerUnLibro(id)
	if err != nil {
		log.Printf("fallo al leer un libro, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el libro con id %s no se pudo encontrar", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(libro)
}

func (h *Handler) LeerLibros(writer http.ResponseWriter, req *http.Request) {
	Libros, err := h.controller.LeerLibros(100, 0)
	if err != nil {
		log.Printf("fallo al leer Libros, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los libros", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(Libros)
}

func (h *Handler) CrearLibro(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo libro, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo libro", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoId, err := h.controller.CrearLibro(body)
	if err != nil {
		log.Println("fallo al crear un nuevo Libro, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo libro", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nuevo libro: %d", nuevoId)))
}
