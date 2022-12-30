package user

import (
	"github.com/gin-gonic/gin"
	global "github.com/titrxw/go-framework/src/Global"
	model "github.com/titrxw/smart-home-server/app/Model"
)

type SessionData map[string]interface{}

type UserOauth struct {
}

func (userOauth UserOauth) SaveUserToSession(ctx *gin.Context, user *model.User) error {
	return global.FHttpServer.Session.Set(ctx, "user_id", user.ID)
}

func (userOauth UserOauth) RemoveUserFromSession(ctx *gin.Context) error {
	return global.FHttpServer.Session.Delete(ctx, "user_id")
}

func (userOauth UserOauth) GetUserIdFromSession(ctx *gin.Context) model.UID {
	userId, ok := global.FHttpServer.Session.Get(ctx, "user_id").(uint)
	if !ok {
		return 0
	}

	return model.UID(userId)
}
