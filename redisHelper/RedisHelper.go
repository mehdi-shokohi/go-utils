package redisHelper

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	utilsConfig "github.com/mehdi-shokohi/go-utils/config"
)

func SaveKey(context context.Context, key string, value interface{}, expire time.Duration) *redis.StatusCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)
	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.Set(context, key, value, expire)

}
func SaveKeyLifeTime(context context.Context, key string, value interface{}) *redis.StatusCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.Set(context, key, value, 0)

}
func SaveOnTable(context context.Context, key string, value ...interface{}) *redis.IntCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.HSet(context, key, value)

}

func GetFromTable(context context.Context, key string, field string) *redis.StringCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.HGet(context, key, field)
}

func GetValue(context context.Context, key string) *redis.StringCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.Get(context, key)
}

func GetTTL(context context.Context, key string) *redis.DurationCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)

	return redisDB.TTL(context, key)
}

func GetAll(context context.Context, key string) *redis.StringStringMapCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)
	return redisDB.HGetAll(context, key)
}

func DeleteKeys(ctx context.Context, key ...string) *redis.IntCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)
	return redisDB.Del(ctx, key...)
}
func CheckExists(ctx context.Context, key ...string) bool {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)
	ex, err := redisDB.Exists(ctx, key...).Result()
	if err != nil {
		return false
	}
	if ex == int64(len(key)) {
		return true
	}
	return false
}
func DeleteFromTable(ctx context.Context, key string, field ...string) *redis.IntCmd {
	rdb := GetRedis(utilsConfig.GetUtilsConf().RedisURI, utilsConfig.GetUtilsConf().RedisPassword,0)

	redisDB := rdb.GetConnection()
	defer rdb.Release(redisDB)
	return redisDB.HDel(ctx, key, field...)
}
