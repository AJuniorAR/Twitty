package routers

import (
	"errors"
	"strings"

	"github.com/AJuniorAR/Twitty/bd"
	"github.com/AJuniorAR/Twitty/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email valor del email usado en tdos los endpoint
var Email string

//IDUsuarios es el ID devuelto del modelo,que se usara en todos los endpoints
var IDUsuario string

//ProcesoToken proceso token para extraer sus valores

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, IDUsuario, err
}
