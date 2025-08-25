package models

// DadosAuth{} contém o id e o token do usuário autenticado.
type DadosAuth struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
