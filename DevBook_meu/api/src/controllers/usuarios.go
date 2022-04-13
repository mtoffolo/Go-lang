package controllers

import (
	database "api/src/Database"
	"api/src/model"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoResquest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {

		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario

	if erro = json.Unmarshal(corpoResquest, &usuario); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)
		return

	}
	if erro := usuario.Preparar("Cadastro"); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)
		return

	}

	db, erro := database.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	respostas.JSON(w, http.StatusCreated, usuario)

}
func BuscarTodosUsuario(w http.ResponseWriter, r *http.Request) {

	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := database.Conectar()

	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	fmt.Println(parametros)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return

	}
	db, erro := database.Conectar()

	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuario, erro := repositorio.BuscarporID(usuarioID)

	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)

		return

	}

	respostas.JSON(w, http.StatusOK, usuario)

}
func AtualizadoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return

	}

	corpoResquisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)

		return

	}

	var usuario model.Usuario

	if erro = json.Unmarshal(corpoResquisicao, &usuario); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)

		return

	}

	if erro = usuario.Preparar("edicao"); erro != nil {

		respostas.Erro(w, http.StatusBadRequest, erro)

		return

	}
	db, erro := database.Conectar()

	if erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)

		return

	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)

		return

	}

	db, erro := database.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}

	defer db.Close()

	var usuario model.Usuario

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if erro = repositorio.DeletarUsuario(usuarioID, usuario); erro != nil {

		respostas.Erro(w, http.StatusInternalServerError, erro)

		return

	}

}
