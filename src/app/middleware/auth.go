package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

func GetUserIdFromJwtToken(ctx context.Context, c *app.RequestContext) int64 {
	claims, err := JwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		log.Println(err.Error())
		// JwtMiddleware.Unauthorized(ctx, c, 1, "Get UserId from jwt token failed")
		return -1
	}
	user := claims[IdentityKey].(map[string]interface{})
	id := int64(user["Id"].(float64))
	return id
}
