package config

import "os"

type UtilsConfig struct {
	JWTUserContext         string
	JWTPrivateKey          string
	JWTPublicKey           string
	PrivateIntercomSecKey  string
	PublicIntercomSecKey   string
	UserHeaderFiberContext string
	RedisURI               string
	MongoURI               string
	RedisPassword string
	RedisDb int

	ConfigDb func(key string) interface{}
}

var config *UtilsConfig

func InitUtils(uc UtilsConfig) {
	config = &uc
	if config.JWTPrivateKey == "" {
		config.JWTPrivateKey = JWTPrivateKey
	}
	if config.JWTUserContext == "" {
		config.JWTUserContext = JWTUserContext
	}
	if config.JWTPublicKey == "" {
		config.JWTPublicKey = JWTPublicKey
	}
	if config.PrivateIntercomSecKey == "" {
		config.PrivateIntercomSecKey = PrivateIntercomSecKey
	}
	if config.PublicIntercomSecKey == "" {
		config.PublicIntercomSecKey = PublicIntercomSecKey
	}
	if config.UserHeaderFiberContext == "" {
		config.UserHeaderFiberContext = UserHeaderFiberContext
	}
	if config.MongoURI == "" {
		config.MongoURI = os.Getenv("MONGO_ADDRESS")
	}
	if config.RedisURI == "" {
		config.RedisURI = os.Getenv("KEYDB_ADDRESS")
	}

}

func GetUtilsConf() *UtilsConfig {
	return config
}
