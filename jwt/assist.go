package jwthandler

import "github.com/golang-jwt/jwt/v5"


func SystemAuthJwt() string {
	claims := jwt.MapClaims{
		"roles":   []string{"system"},
		"domains": []string{"internal"},
	}

	t ,err:= JwtEdSign( claims)

	if err != nil {
		panic(err)
	}
	return "Bearer " + t
}