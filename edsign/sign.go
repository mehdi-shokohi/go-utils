package edsign

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"

	models "github.com/mehdi-shokohi/go-utils/config"

	globUtils "github.com/mehdi-shokohi/go-utils/global"
	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

// PrivateKey is ed25519.PrivateKey.
type PrivateKey ed25519.PrivateKey

// GetPrivateKey reads the private key from input file and
// returns the initialized PrivateKey.
func GetPrivateKey() (PrivateKey, error) {
	rv := redisHelper.GetValue(context.Background(), models.GetUtilsConf().PrivateIntercomSecKey)
	var key any
	if rv.Val() != "" {
		println("ed25119 Private Key Load From Cache")
		p, err := getPemDecode(rv.Val())
		if err != nil {
			return nil, err
		}
		key, err = x509.ParsePKCS8PrivateKey(p)
		if err != nil {
			return nil, fmt.Errorf("key is not ed25519 key")
		}
	} else {
		// load From DB
		privateKeyDb, err := GetConfigFromDb(models.GetUtilsConf().PrivateIntercomSecKey)
		if err != nil {
			return nil, err
		}
		p, err := getPemDecode(privateKeyDb)
		if err != nil {
			return nil, err
		}
		key, err = x509.ParsePKCS8PrivateKey(p)
		if err != nil {
			return nil, err
		}
		redisHelper.SaveKeyLifeTime(context.Background(), models.GetUtilsConf().PrivateIntercomSecKey, privateKeyDb)

	}
	edKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not ed25519 key")
	}

	return PrivateKey(edKey), nil
}

// Sign reads the input file and compute the ED25519 signature
// using the private key.
func Sign(message string) (string, error) {
	p, err := GetPrivateKey()
	if err != nil {
		return "", err
	}
	signature := ed25519.Sign(ed25519.PrivateKey(p), []byte(message))
	return hex.EncodeToString(signature), nil
}

func getPemDecode(p string) ([]byte, error) {
	kEnc, _ := pem.Decode([]byte(p))
	if kEnc == nil {
		return nil, fmt.Errorf("no pem block found")
	}
	return kEnc.Bytes, nil
}

func Signer(s ...string) string {
	message := strings.Join(s, "::")
	signed, err := Sign(message)
	if err != nil {
		globUtils.LoggerAlert(err.Error())
	}
	return signed
}
