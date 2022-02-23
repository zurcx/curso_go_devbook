package modelos

// DadosAutenticação contém id e token do usuário autenticado.
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
