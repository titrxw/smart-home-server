package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/model"
)

type SessionData map[string]interface{}

type UserOauth struct {
}

func (userOauth UserOauth) SaveUserToSession(ctx *gin.Context, user *model.User) error {
	sessions.Default(ctx).Set("user_id", user.ID)
	return sessions.Default(ctx).Save()
}

func (userOauth UserOauth) RemoveUserFromSession(ctx *gin.Context) error {
	sessions.Default(ctx).Delete("user_id")
	return sessions.Default(ctx).Save()
}

func (userOauth UserOauth) GetUserIdFromSession(ctx *gin.Context) model.UID {
	userId, ok := sessions.Default(ctx).Get("user_id").(uint)
	if !ok {
		return 0
	}

	return model.UID(userId)
}
