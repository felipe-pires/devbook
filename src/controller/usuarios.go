package controller

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario..."))

	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario model.Usuario

	if erro = json.Unmarshal(body, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	repositorio := repository.UsuarioRepository(db)
	ID, erro := repositorio.Criar(usuario)

	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Id inserido: %d", ID)))
}

// BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios..."))
}

// BuscarUsuario busca um usuario pelo id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario com id: "))
}

// AtualizarUsuario altera informacoes de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario..."))
}

// DeletarUsuario deleta um usuario pelo id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario com id: "))
}
