package globUtils

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/mehdi-shokohi/go-utils/config"
)

func SendErrorRecoveryHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {

				if e, ok := r.(error); ok {
					c.SendStatus(400)
					c.JSON(models.Response{
						Data:  nil,
						Error: models.Error{Message: e.Error(), Code: -1},
					})

				} else if v, ok := r.(map[string]error); ok {
					t := ""
					for eKey, errMessage := range v {
						t += eKey + errMessage.Error() + " , "
					}
					c.SendStatus(400)
					c.JSON(models.Response{Data: nil, Error: models.Error{Message: t, Code: -1}})
				}
				c.SendStatus(500)
				c.JSON(models.Response{Data: nil, Error: r})

			}

		}()
		return c.Next()
	}

}
