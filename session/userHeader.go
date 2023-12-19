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
	cacheAdapter := redisHelper.GetRedis(models.GetUtilsConf().RedisURI, "")
	res := cacheAdapter.GetConnection().Get(c, models.SessionRedisPrefix+header.GetSessionId())
	if res.Val() == "" {
		return errors.New("invalid session")
	}
	return nil
}
