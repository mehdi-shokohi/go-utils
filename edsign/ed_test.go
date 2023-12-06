package edsign

import (
	"fmt"
	"os"
	"testing"

	utilsConfig "github.com/mehdi-shokohi/go-utils/config"
	"github.com/mehdi-shokohi/mongoHelper/config"
)

func TestSign(t *testing.T) {
	os.Setenv("KEYDB_ADDRESS", "localhost:6040")
	os.Setenv("MONGO_ADDRESS", "mongodb://localhost:27018/inflowenger?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	os.Setenv("MONGODB_NAME", "inflowenger")
	config.SetConfig(config.Config{
		MongoAddress: utilsConfig.GetUtilsConf().MongoURI,
		MongoDbName:  os.Getenv("MONGODB_NAME"),
	}) // once run

}

func TestSignEntity(t *testing.T) {
	os.Setenv("KEYDB_ADDRESS", "localhost:6040")
	os.Setenv("MONGO_ADDRESS", "mongodb://localhost:27018/inflowenger?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	os.Setenv("MONGODB_NAME", "inflowenger")
	config.SetConfig(config.Config{
		MongoAddress: utilsConfig.GetUtilsConf().MongoURI,
		MongoDbName:  os.Getenv("MONGODB_NAME"),
	}) // once run

}

func TestSinerEmpty(t *testing.T) {
	os.Setenv("KEYDB_ADDRESS", "localhost:6060")
	os.Setenv("MONGO_ADDRESS", "mongodb://localhost:27018/inflowenger?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	os.Setenv("MONGODB_NAME", "inflowenger")
	config.SetConfig(config.Config{
		MongoAddress: utilsConfig.GetUtilsConf().MongoURI,
		MongoDbName:  os.Getenv("MONGODB_NAME"),
	}) // once run
	data := []string{}
	s := fmt.Sprintf("%v", data)
	fmt.Println(s)
	signed := Signer(s, "1")
	fmt.Println(signed)
	res, _ := Verifier(signed, s, "1")
	println(res)

}
