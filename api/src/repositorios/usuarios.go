package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarPorId(usuarioId uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioId)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(Id uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, email = ?, nick = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, usuario.Nick, Id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(Id uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) PararDeSeguir(usuarioId, seguidorId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u
		inner join seguidores s on u.id = s.seguidor_id
		where s.usuario_id = ?
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

func (repositorio usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u
		inner join seguidores s on u.id = s.usuario_id
		where s.seguidor_id = ?
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}
