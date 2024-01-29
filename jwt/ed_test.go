package jwthandler

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"

	"github.com/mehdi-shokohi/go-utils/edsign"
	_ "github.com/mehdi-shokohi/go-utils/edsign"
)

func TestED(t *testing.T) {
	os.Setenv("KEYDB_ADDRESS", "localhost:6040")
	os.Setenv("MONGO_ADDRESS", "mongodb://localhost:27018/inflowenger?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	os.Setenv("MONGODB_NAME", "inflowenger")
	// config.SetConfig(config.Config{
	// 	MongoAddress: utilsConfig.GetUtilsConf().MongoURI,
	// 	MongoDbName:  os.Getenv("MONGODB_NAME"),
	// }) // once run
	signed, err := JwtEdSign(jwt.MapClaims{
		"sess_id3": 111, "sess_id7": 111, "sess_id8": 111, "username": "3233333", "username4": 222,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	print(signed)

	ut, err := GetJwtToken(signed)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(ut.Claims.(jwt.MapClaims))
}
func TestGenEd25519(t *testing.T) {
	edsign.GenerateSaveEd25519("m1")

	// pu, pr, _ := ed25519.GenerateKey(rand.Reader)
	// fmt.Println(base64.StdEncoding.EncodeToString(pu))
	// fmt.Println(base64.StdEncoding.EncodeToString(pr))
	// fmt.Println(hex.EncodeToString(pu))
	// fmt.Println(hex.EncodeToString(pr))
}
