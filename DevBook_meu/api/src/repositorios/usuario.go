package repositorios

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

//Usuario representa um repositório de usuário
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Cria insere um usuário

func (repositorio Usuarios) Criar(Usuario model.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare("INSERT INTO public.usuarios(nome, nick, email, senha) VALUES($1, $2, $3, $4);")

	if erro != nil {
		return 0, erro

	}

	defer statement.Close()
	resultado, erro := statement.Exec(Usuario.Nome, Usuario.Nick, Usuario.Email, Usuario.Senha)

	if erro != nil {
		return 0, erro

	}

	ultimoIdInserido, erro := resultado.RowsAffected()

	if erro != nil {
		return 0, erro

	}

	return uint64(ultimoIdInserido), nil

}

func (repositorio Usuarios) Buscar(nomeOunick string) ([]model.Usuario, error) {

	nomeOunick = fmt.Sprintf("%%%s%%", nomeOunick) //%nomeOunick%

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoem FROM public.usuarios where lower(nome) LIKE $1 or lower(nick) LIKE $2", nomeOunick, nomeOunick)

	if erro != nil {

		return nil, erro

	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {

		var usuario model.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriandoEm); erro != nil {
			return nil, erro

		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil

}

/*Buscar por Id, traz um unico usuário do banco de dados */
func (repositorio Usuarios) BuscarporID(Idusuario uint64) (model.Usuario, error) {

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoem FROM public.usuarios where id = $1", Idusuario)

	if erro != nil {

		return model.Usuario{}, erro

	}

	defer linhas.Close()
	var usuario model.Usuario
	if linhas.Next() {

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriandoEm); erro != nil {
			return model.Usuario{}, erro

		}

	}
	return usuario, nil

}

/*Altera as Informações de um usuário*/
func (repositorio Usuarios) Atualizar(ID uint64, usuario model.Usuario) error {
	statement, erro := repositorio.db.Prepare("update public.usuarios set nome = $1, nick = $2, email = $3 where id = $4")

	if erro != nil {

		return erro

	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {

		return erro
	}

	return nil

}

func (repositorio Usuarios) DeletarUsuario(ID uint64, usuario model.Usuario) error {
	statement, erro := repositorio.db.Prepare("delete from public.usuarios where id = $1")

	if erro != nil {

		return erro

	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {

		return erro
	}

	return nil

}

// BuscarPorEmail retorna o seu id, e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id,senha FROM public.usuarios where email = $1", email)

	if erro != nil {
		return model.Usuario{}, erro

	}

	var usuario model.Usuario

	defer linha.Close()

	if linha.Next() {

		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {

			return model.Usuario{}, erro

		}

	}
	return usuario, nil
}
