package jwthandler

import (
	"context"
	"crypto"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	utilsConfig "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/edsign"
	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

const Bearer = "Bearer"

func JwtEdVerify() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		reqToken := c.Get(fiber.HeaderAuthorization)
		token, err := GetJwtToken(reqToken)
		if err != nil {
			return fiber.ErrUnauthorized

		}
		c.Locals(utilsConfig.GetUtilsConf().JWTUserContext, token)
		return c.Next()

	}
}
func JwtEdVerifyInline(cin *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		reqToken := c.Get(fiber.HeaderAuthorization)
		token, err := GetJwtToken(reqToken)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		c.Locals(utilsConfig.GetUtilsConf().JWTUserContext, token)
		return c.Next()

	}(cin)
}
func getPublicKey() (interface{}, error) {
	rv := redisHelper.GetValue(context.Background(), utilsConfig.GetUtilsConf().JWTPublicKey)
	var key any
	if rv.Val() != "" {
		println("ed25119 JWT Public Key Load From Cache")

		var errPem error
		key, errPem = jwt.ParseEdPublicKeyFromPEM([]byte(rv.Val()))

		if errPem != nil {
			return nil, fmt.Errorf("key is not ed25519 key")
		}
	} else {
		// load From DB
		pulicKeyDb, err := edsign.GetConfigFromDb(utilsConfig.GetUtilsConf().JWTPublicKey)
		if err != nil {
			return nil, err
		}

		var errPem error
		key, errPem = jwt.ParseEdPublicKeyFromPEM([]byte(pulicKeyDb))

		if errPem != nil {
			return nil, fmt.Errorf("key is not ed25519 key")
		}
		redisHelper.SaveKeyLifeTime(context.Background(), utilsConfig.GetUtilsConf().JWTPublicKey, pulicKeyDb)

	}

	return key, nil

}
func keyfunc() jwt.Keyfunc {
	return func(*jwt.Token) (interface{}, error) {
		eddPubKey, err := getPublicKey()
		if err != nil {
			return nil, fiber.ErrInternalServerError
		}
		return eddPubKey, nil
	}
}
func JwtEdSign(claims jwt.MapClaims) (string, error) {
	ed25519PrivateKey, err := getPrivateKey()
	if err != nil {
		return "", fiber.ErrInternalServerError
	}

	method := jwt.GetSigningMethod(jwt.SigningMethodEdDSA.Alg())

	token := jwt.NewWithClaims(method, claims)
	signer := ed25519PrivateKey.(crypto.Signer)

	sig, err := token.SignedString(signer)
	if err != nil {
		return "", fiber.ErrInternalServerError
	}
	return sig, nil
}

func getPrivateKey() (interface{}, error) {
	rv := redisHelper.GetValue(context.Background(), utilsConfig.GetUtilsConf().JWTPrivateKey)
	var key any
	if rv.Val() != "" {
		println("ed25119 Private Key Load From Cache")
		var err error
		key, err = jwt.ParseEdPrivateKeyFromPEM([]byte(rv.Val()))
		if err != nil {
			return nil, err
		}

	} else {
		// load From DB
		privateKeyDb := os.Getenv(utilsConfig.GetUtilsConf().JWTPrivateKey)
		if privateKeyDb == "" {
			return nil, errors.New("private key not found")
		}
		var err error
		key, err = jwt.ParseEdPrivateKeyFromPEM([]byte(privateKeyDb))
		if err != nil {
			return nil, err
		}
		redisHelper.SaveKeyLifeTime(context.Background(), utilsConfig.GetUtilsConf().JWTPrivateKey, privateKeyDb)

	}
	return key, nil
}

func GetJwtToken(reqToken string) (*jwt.Token, error) {
	bearer := strings.Split(reqToken, " ")
	if len(bearer) == 2 {
		if bearer[0] != Bearer {
			return nil, errors.New("jwt malformed")
		}
	} else {
		return nil, errors.New("jwt malformed")
	}

	token, err := jwt.Parse(bearer[1], keyfunc())
	if err != nil {
		return nil, err
	}

	return token, nil
}
