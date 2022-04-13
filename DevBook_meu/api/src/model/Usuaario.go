package model

import (
	security "api/src/Security"
	"errors"
	"regexp"
	"strings"
	"time"
)

type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriandoEm time.Time `json:"CriandoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {

	if erro := usuario.validar(etapa); erro != nil {
		return erro

	}

	if erro := usuario.formatar(etapa); erro != nil {

		return erro

	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {

		return errors.New("Usuário não pode estar em branco!")
	}
	if usuario.Nick == "" {

		return errors.New("Nick não pode estar em branco!")
	}
	if usuario.Email == "" {

		return errors.New("Email não pode estar em branco!")
	}
	if etapa == "Cadastro" && usuario.Senha == "" {

		return errors.New("Senha não pode estar em branco!")
	}

	if match, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, usuario.Email); match != true {

		return errors.New("Email Inválido")

	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

	if etapa == "Cadastro" {

		senhaComHash, erro := security.Hash(usuario.Senha)

		if erro != nil {

			return erro
		}

		usuario.Senha = string(senhaComHash)

	}
	return nil
}
