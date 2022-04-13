package controllers

import (
	database "api/src/Database"
	security "api/src/Security"
	"api/src/autenticacao"
	"api/src/model"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*Resposável por Autenticar um usuário na API*/
func Login(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {

		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	db, erro := database.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	if erro = security.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {

		respostas.Erro(w, http.StatusUnauthorized, erro)
		return

	}
	token, erro := autenticacao.CriarToken(int64(usuarioSalvoNoBanco.ID))

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	w.Write([]byte(token))

}
