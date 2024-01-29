package sessionHandler

import (
	"context"
	"errors"
	"time"

	models "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

func SessionExpireValidation(header models.IJWTHeader) error {
	if header.GetExpireTime() > time.Now().Unix() {
		return nil
	}
	return errors.New("session has been expired")
}

func SessionIdValidations(c context.Context, header models.IJWTHeader) error {
	cacheAdapter := redisHelper.GetRedis(models.GetUtilsConf().RedisURI, models.GetUtilsConf().RedisPassword,0)
	res := cacheAdapter.GetConnection().HGet(c, models.SessionRedisPrefix+header.GetSessionId(),header.GetUserId())
	if res.Val() == "1" {
		return nil
	}
	if res.Val() == "" {
		// get ensure , previous function have done validate of expiration time.
		cacheAdapter.GetConnection().HSet(c,models.SessionRedisPrefix+header.GetSessionId(),header.GetUserId(),1)
		cacheAdapter.GetConnection().ExpireAt(c,models.SessionRedisPrefix+header.GetSessionId(),time.Unix(header.GetExpireTime()+10,0))
		return SessionIdValidations(c,header)
	}
	return errors.New("invalid session")
}
