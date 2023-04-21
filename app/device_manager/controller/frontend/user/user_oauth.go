package user

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type SessionData map[string]interface{}

type UserOauth struct {
}

func (userOauth UserOauth) SaveUserToSession(ctx *gin.Context, user *model.User) error {
	return server.GHttpServer.Session.Set(ctx, "user_id", user.ID)
}

func (userOauth UserOauth) RemoveUserFromSession(ctx *gin.Context) error {
	return server.GHttpServer.Session.Delete(ctx, "user_id")
}

func (userOauth UserOauth) GetUserIdFromSession(ctx *gin.Context) model.UID {
	userId, ok := server.GHttpServer.Session.Get(ctx, "user_id").(uint)
	if !ok {
		return 0
	}

	return model.UID(userId)
}
