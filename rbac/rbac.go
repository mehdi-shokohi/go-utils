package rbac

import (
	"context"

	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	redisadapter "github.com/casbin/redis-adapter/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/mehdi-shokohi/go-utils/config"
	utilsConfig "github.com/mehdi-shokohi/go-utils/config"

	"github.com/mehdi-shokohi/go-utils/redisHelper"
)

const GlobalDomain = "global"

func getRbacModelConfig() string {
	cache := redisHelper.GetValue(context.Background(), "rbac_model_config")
	// print(cache.Val());
	casbinConf := cache.Val()
	if casbinConf == "" {
		casbinConf = getDbRbacModelConfig()
		redisHelper.SaveKeyLifeTime(context.Background(), "rbac_model_config", casbinConf)
		return casbinConf
	}

	return casbinConf
}

func getDbRbacModelConfig() string {
	// dbConfig := make(map[string]string)
	// r := mongoHelper.NewMongo(context.Background(), utilsConfig.GetUtilsConf().ConfigCollectionName, dbConfig)
	// rbacModel, err := r.FindOne(&bson.D{{Key: "key", Value: "rbac_model"}})
	// if err != nil {
	// 	panic(err)
	// }
	key := "rbac_model"
	cfg := config.GetUtilsConf().ConfigDb(key)
	if cfgVal, ok := cfg.(string); ok {
		println(fmt.Sprintf("Loaded %s key from DB", key))

		return cfgVal

	}
	return ""
}

func getEnforcer(c *fiber.Ctx) *casbin.Enforcer {

	m, err := model.NewModelFromString(getRbacModelConfig())
	if err != nil {
		panic(err)
	}
	if !redisHelper.CheckExists(context.Background(), "casbin_rules") {
		dbacl, err := mongodbadapter.NewAdapter(utilsConfig.GetUtilsConf().MongoURI) // Your MongoDB URL.
		if err != nil {
			panic(err)
		}

		e, err := casbin.NewEnforcer(m, dbacl)
		if err != nil {
			panic(err)
		}
		e.SetAdapter(dbacl)
		err = e.LoadPolicy()
		if err != nil {
			panic(err)
		}
		mrbac := redisadapter.NewAdapter("tcp", utilsConfig.GetUtilsConf().RedisURI)
		if err != nil {
			panic(err)
		}
		re, err := casbin.NewEnforcer(m, mrbac)
		if err != nil {
			panic(err)
		}
		re.AddPolicies(e.GetPolicy())
		// re.SetAdapter(dbacl)
		// re.LoadPolicy()
		// re.SetAdapter(mrbac)
		// re.LoadPolicy()

		err = re.SavePolicy()
		if err != nil {
			panic(err)
		}
	}
	rbacPolicy := redisadapter.NewAdapter("tcp", utilsConfig.GetUtilsConf().RedisURI)
	if err != nil {
		panic(err)
	}

	re, err := casbin.NewEnforcer(m, rbacPolicy)
	if err != nil {
		panic(err)
	}
	return re
}

func HasRole(r []string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		re := getEnforcer(c)
		user := c.Locals(utilsConfig.GetUtilsConf().JWTUserContext).(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		if userRoles, ok := claims["roles"].([]interface{}); ok {
			for _, urole := range userRoles {
				for _, role := range r {
					ur := urole.(string)
					if doms, ok := claims["domains"].([]interface{}); ok {
						for _, dom := range doms {
							domain := dom.(string)
							has, err := re.Enforce(ur, domain, role)
							if err != nil {
								panic(err)
							}
							if has {
								return c.Next()

							}
						}
					}

				}

			}
			// return c.JSON(fiber.Map{"status": "forbiden"})
			c.Status(fiber.StatusForbidden)
			return c.JSON(utilsConfig.Response{Data: nil, Error: fiber.Map{"status": "forbiden"}})

		}
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(utilsConfig.Response{Data: nil, Error: fiber.Map{"status": "UnAutorized"}})

	}
}
func HasRoutePermission() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		re := getEnforcer(c)
		user := c.Locals(utilsConfig.GetUtilsConf().JWTUserContext).(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		if userRoles, ok := claims["roles"].([]interface{}); ok {
			for _, ur := range userRoles {
				userRole := ur.(string)
				fmt.Println(c.Path(), c.Method())
				if doms, ok := claims["domains"].([]interface{}); ok {
					for _, dom := range doms {
						domain := dom.(string)
						okEnforce, err := re.Enforce(userRole, domain, c.Path(), c.Method())
						if okEnforce && err == nil {
							return c.Next()

						}
					}
				}
			}
			return c.JSON(utilsConfig.Response{Data: nil, Error: fiber.Map{"status": "forbiden"}})

		}
		return c.JSON(utilsConfig.Response{Data: nil, Error: fiber.Map{"status": "UnAutorized"}})
	}
}
func OwnerDefinedRolesValidate(permissions []string, action string) bool {
	return CheckInArray(permissions, action)
}
