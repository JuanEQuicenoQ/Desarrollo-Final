package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/JuanEQuicenoQ/Desarrollo-Final/models"
	repositorio "github.com/JuanEQuicenoQ/Desarrollo-Final/repository"
)

var (
	updateQuery = "UPDATE libros SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM libros WHERE id=$1;"
	selectQuery = "SELECT id, titulo, autor, edicion, pais, publicaciones, BestSeller FROM libros WHERE id=$1;"
	listQuery   = "SELECT id, titulo, autor, edicion, pais, publicaciones, BestSeller FROM libros limit $1 offset $2"
	createQuery = "INSERT INTO libros ( titulo, autor, edicion, pais, publicaciones, BestSeller) VALUES (:titulo, :autor, :edicion, :pais, :publicaciones, :BestSeller) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Libro]
}

func NewController(repo repositorio.Repository[models.Libro]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnLibro(reqBody []byte, id string) error {
	nuevosValoresLibro := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresLibro)
	if err != nil {
		log.Printf("fallo al actualizar un libro, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un libro, con error: %s", err.Error())
	}

	if len(nuevosValoresLibro) == 0 {
		log.Printf("fallo al actualizar un libro, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un libro, con error: %s", err.Error())
	}

	query := construirUpdateQuery(nuevosValoresLibro)
	nuevosValoresLibro["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresLibro)
	if err != nil {
		log.Printf("fallo al actualizar un libro, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un libro, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columns := []string{}
	for key := range nuevosValores {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) EliminarUnLibro(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un libro, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un libro, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnLibro(id string) ([]byte, error) {
	libro, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un libro, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un libro, con error: %s", err.Error())
	}

	libroJson, err := json.Marshal(libro)
	if err != nil {
		log.Printf("fallo al leer un libro, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un libro, con error: %s", err.Error())
	}
	return libroJson, nil
}

func (c *Controller) LeerLibros(limit, offset int) ([]byte, error) {
	libros, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer libros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer libros, con error: %s", err.Error())
	}

	jsonLibros, err := json.Marshal(libros)
	if err != nil {
		log.Printf("fallo al leer libros, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer libros, con error: %s", err.Error())
	}
	return jsonLibros, nil
}

func (c *Controller) CrearLibro(reqBody []byte) (int64, error) {
	nuevoLibro := &models.Libro{}
	err := json.Unmarshal(reqBody, nuevoLibro)
	if err != nil {
		log.Printf("fallo al crear un nuevo libro, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo libro, con error: %s", err.Error())
	}

	valoresColumnasNuevoLibro := map[string]any{
		"titulo":        nuevoLibro.Titulo,
		"autor":         nuevoLibro.Autor,
		"edicion":       nuevoLibro.Edicion,
		"pais":          nuevoLibro.Pais,
		"publicaciones": nuevoLibro.Publicacion,
		"BestSeller":    nuevoLibro.BestSeller,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoLibro)
	if err != nil {
		log.Printf("fallo al crear un nuevo libro, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo libro, con error: %s", err.Error())
	}
	return nuevoId, nil
}
