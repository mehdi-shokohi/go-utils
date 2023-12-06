package config

import (
	"errors"
	"strings"
	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

)
const (
	JWTUserContext        = "jwtUser"
	PrivateIntercomSecKey = "ipc_private_key"
	PublicIntercomSecKey  = "ipc_public_key"
	JWTPrivateKey         = "jwt_private_key"
	JWTPublicKey          = "jwt_public_key"
	UserHeaderFiberContext    = "userHeader"
)
const (
	Active      = "active"
	InActive    = "inactive"
	BlockedUser = "blocked"
)
const (
	SessionLogout         = "ban_sessions"
	SessionRedisPrefix    = "sess_id_"
	UserBanned            = "ban_users"
	UserBannedRedisPrefix = "ban_user_"
)
const (
	defaultPage    = 1
	defaultPerPage = 5
)
type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidationError struct {
	Field string
	Rule  string
	Param string
	Message string
}

type PaginationParams struct {
	Page    uint   `query:"page"`
	PerPage uint   `query:"per_page"`
	Search  string `query:"search"`
}
func (i *PaginationParams) Validate() error {
	i.Search = strings.Trim(i.Search," ")
	if i.Page == 0 {
		i.Page = defaultPage
	}

	if i.PerPage == 0 {
		i.PerPage = defaultPerPage
	}
	if i.Search!="" && len(i.Search)<5{
		return errors.New("bad request")
	}
	return nil
}




type IJWTHeader interface{
	GetExpireTime()int64
	GetUserId()string
	GetDomain()[]string
	GetSessionId()string
	GetStatus()string
}

func LoadJwtHeader[T any](c *fiber.Ctx) *T {
	if session, ok := c.Locals(GetUtilsConf().JWTUserContext).(*jwt.Token); ok {
		claims := session.Claims.(jwt.MapClaims)
		userjwt, _ := json.Marshal(claims)
		jwtUser := new(T)
		json.Unmarshal(userjwt, jwtUser)
		return jwtUser
	}
	return nil
}