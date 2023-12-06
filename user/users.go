package userHandler

import (
	"context"
	"errors"
	"os"

	models "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

func UsersIdValidations(c context.Context, header models.IJWTHeader) error {
	cacheAdapter := redisHelper.GetRedis(os.Getenv("KEYDBADDRESS"), "")
	res := cacheAdapter.GetConnection().Get(c, models.UserBannedRedisPrefix+header.GetUserId())
	if res.Val() != "" {
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