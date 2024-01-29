package userHandler

import (
	"context"
	"errors"
	

	models "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

func UsersIdValidations(c context.Context, header models.IJWTHeader) error {
	cacheAdapter := redisHelper.GetRedis(models.GetUtilsConf().RedisURI, models.GetUtilsConf().RedisPassword,0)
	res := cacheAdapter.GetConnection().Get(c, models.UserBannedRedisPrefix+header.GetUserId())
	if res.Val() == "" {
		return errors.New("you are banned  access to system")
	}
	return nil
}

func UserCheckActive(c context.Context, headerStatus string) error {
	if headerStatus == models.InActive {
		return errors.New("user is inActive")
	}
	return nil
}
