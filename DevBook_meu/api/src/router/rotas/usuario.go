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
		RequerAutenticao: false,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizadoUsuario,
		RequerAutenticao: false,
	},
	{
		URI:              "/usuario/{usuarioId}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.DeletaUsuario,
		RequerAutenticao: false,
	},
}
