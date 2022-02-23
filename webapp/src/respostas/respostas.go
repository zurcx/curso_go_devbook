package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato json para a requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}

	// if dados != nil {
	// 	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
	// 		log.Fatal(erro)
	// 	}

	// }

}

// TratarStatusCodeDeErro trata as requisicoes com status 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}