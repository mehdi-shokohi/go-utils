package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	models "github.com/mehdi-shokohi/go-utils/config"

	sessionHandler "github.com/mehdi-shokohi/go-utils/session"
	userHandler "github.com/mehdi-shokohi/go-utils/user"
)

func Run(c context.Context, header models.IJWTHeader) error {
	err := sessionHandler.SessionExpireValidation(header)
	if err != nil {
		return err
	}
	err = sessionHandler.SessionIdValidations(c, header)
	if err != nil {
		return err
	}
	// err = userHandler.UsersIdValidations(c, header)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func RunFiber[T models.IJWTHeader]() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userHeader := getUserHeader[T](c)
		err := Run(c.Context(), userHeader)
		if err == nil {
			c.Next()
			return nil
		}
		c.SendStatus(401)
		return c.JSON(models.Response{Data: nil, Error: err.Error()})

	}
}

func RunFiberUserActive[T models.IJWTHeader]() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userHeader := getUserHeader[T](c)
		err := userHandler.UserCheckActive(c.Context(), (userHeader).GetStatus())
		if err == nil {
			c.Next()
		}
		c.SendStatus(fiber.StatusNotAcceptable)
		return c.JSON(models.Response{Data: nil, Error: err.Error()})

	}
}

func getUserHeader[T any](c *fiber.Ctx) T {

	if us, ok := c.Locals(models.GetUtilsConf().UserHeaderFiberContext).(T); ok {
		return us
	}
	userHeader := models.LoadJwtHeader[T](c)
	c.Locals(models.GetUtilsConf().UserHeaderFiberContext, userHeader)
	return *userHeader

}
