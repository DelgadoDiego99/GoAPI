package todo

type Todo struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Finalizado  bool   `json:"finalizado"`
}

var lista = []Todo{
	{ID: 1, Titulo: "Lavar platos", Descripcion: "Lavar platos", Finalizado: false},
	{ID: 1, Titulo: "Lavar ropa", Descripcion: "Colocar la ropa a lavar", Finalizado: false},
	{ID: 1, Titulo: "Regar plantas", Descripcion: "Echar agua a las plantas", Finalizado: false},
}