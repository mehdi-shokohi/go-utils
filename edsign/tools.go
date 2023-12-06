package edsign

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"os"

	"github.com/mehdi-shokohi/go-utils/config"
)

func GetConfigFromDb(key string) (string, error) {

	cfg := config.GetUtilsConf().ConfigDb(key)
	if cfgVal, ok := cfg.(string); ok {
		println(fmt.Sprintf("Loaded %s key from DB", key))

		return cfgVal, nil
	}
	return "", errors.New(fmt.Sprintf("%s key not found", key))
}

func GenerateSaveEd25519(fb string) error {

	var (
		err   error
		b     []byte
		block *pem.Block
		pub   ed25519.PublicKey
		priv  ed25519.PrivateKey
	)

	pub, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = os.WriteFile(fb, pem.EncodeToMemory(block), 0600)
	if err != nil {
		return err
	}

	// public key
	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}

	fileName := fb + ".pub"
	err = os.WriteFile(fileName, pem.EncodeToMemory(block), 0644)
	return err

}
