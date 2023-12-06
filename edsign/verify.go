package edsign

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"strings"

	models "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

type PublicKey ed25519.PublicKey

func GetPublicKey() (PublicKey, error) {
	rv := redisHelper.GetValue(context.Background(), models.GetUtilsConf().PublicIntercomSecKey)
	var key any
	if rv.Val() != "" {
		println("ed25119 Public Key Load From Cache")

		p, err := getPemDecode(rv.Val())
		if err != nil {
			return nil, err
		}
		key, err = x509.ParsePKIXPublicKey(p)
		if err != nil {
			return nil, fmt.Errorf("key is not ed25519 key")
		}
	} else {
		// load From DB
		pulicKeyDb, err := GetConfigFromDb(models.GetUtilsConf().PublicIntercomSecKey)
		if err != nil {
			return nil, err
		}
		p, err := getPemDecode(pulicKeyDb)
		if err != nil {
			return nil, err
		}
		key, err = x509.ParsePKIXPublicKey(p)
		if err != nil {
			return nil, err
		}
		redisHelper.SaveKeyLifeTime(context.Background(), models.GetUtilsConf().PublicIntercomSecKey, pulicKeyDb)

	}
	edKey, ok := key.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not ed25519 key")
	}

	return PublicKey(edKey), nil

}

// Verify checks that input signature is valid. That is, if
// input file was signed by private key corresponding to input
// public key.
func Verify(message, signature string) (bool, error) {
	p, err := GetPublicKey()
	if err != nil {
		return false, err
	}
	byteSign, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	ok := ed25519.Verify(ed25519.PublicKey(p), []byte(message), byteSign)
	return ok, nil
}

func Verifier(signature string, message ...string) (bool, error) {
	messages := strings.Join(message, "::")
	return Verify(messages, signature)
}
