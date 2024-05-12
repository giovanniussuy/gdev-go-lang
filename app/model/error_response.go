package model

type ErrorResponse struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
	Status    int    `json:"status"`
}
