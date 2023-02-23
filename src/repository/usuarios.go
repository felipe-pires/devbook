package repository

import (
	"api/src/model"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

// UsuarioRepository criar repositorio de usuarios
func UsuarioRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere usuario no banco de dados
func (repositorio usuarios) Criar(usuario model.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare("insert into usuarios(nome, nick, email, senha) values (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ID), nil
}
