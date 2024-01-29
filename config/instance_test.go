package config

import (


)


// func TestDbConf(t *testing.T){
// 	InitUtils(UtilsConfig{
// 		JWTUserContext:         "1",
// 		PrivateIntercomSecKey:  "1",
// 		PublicIntercomSecKey:   "1",
// 		JWTPrivateKey:          "1",
// 		JWTPublicKey:           "1",
// 		UserHeaderFiberContext: "1",
// 		ConfigDb:   getDbConfig,
// 	})
// }

// func getDbConfig(key string)interface{}{
	
// 		dbConfig := make(map[string]interface{})
// 		r := mongoHelper.NewMongo(context.Background(), "configs", dbConfig)
// 		_, err := r.FindOne(&bson.D{{Key: "key", Value: key}})
// 		if err != nil {
// 			return nil
// 		}
// 		if v,ok:=dbConfig["value"];ok{
// 			return v
// 		}
// 		return nil
	
// }

