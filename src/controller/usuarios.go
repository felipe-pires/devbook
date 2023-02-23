package controller

import "net/http"

//CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuario..."))
}

//BuscarUsuarios busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios..."))
}

//BuscarUsuario busca um usuario pelo id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario com id: "))
}

//AtualizarUsuario altera informacoes de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario..."))
}

//DeletarUsuario deleta um usuario pelo id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario com id: "))
}
