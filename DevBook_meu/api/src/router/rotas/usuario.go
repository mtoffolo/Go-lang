package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:              "/usuario",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarTodosUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizadoUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletaUsuario,
		RequerAutenticao: true,
	},
	{
		URI:              "/usuario/{usuarioID}/seguir",
		Metodo:           http.MethodPost,
		Funcao:           controllers.SeguirUsuario,
		RequerAutenticao: true,
	},
}
