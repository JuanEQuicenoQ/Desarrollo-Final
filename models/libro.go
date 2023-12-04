package models

type Libro struct {
	Id          int    `db:"id" json:"id"`
	Titulo      string `db:"titulo" json:"titulo"`
	Autor       string `db:"autor" json:"autor"`
	Edicion     uint   `db:"edicion" json:"edicion"`
	Pais        string `db:"pais" json:"pais"`
	Publicacion int    `db:"publicacion" json:"publicacion"`
	BestSeller  bool   `db:"bestseller" json:"bestseller"`
}
