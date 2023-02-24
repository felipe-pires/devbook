package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	totalMessagesByOrigin = "totalMessagesByOrigin"
	totalContextsByOrigin = "totalContextsByOrigin"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = usuario.Prepapar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = json.Unmarshal(body, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.UsuarioRepository(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
}

// BuscarUsuario busca um usuario pelo id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
}

// AtualizarUsuario altera informacoes de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
}

// DeletarUsuario deleta um usuario pelo id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
}
