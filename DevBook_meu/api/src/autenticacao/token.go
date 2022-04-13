package autenticacao

import (
	"api/src/config"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuerioID int64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuerioID"] = usuerioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) // secret
}

/*Verifica se o token passado na request é válido*/
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	return nil

}

func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	// bearer 123
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
