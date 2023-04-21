package user

import (
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/logic"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type RegisterRequest struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,password"`
	EmailVerifyCode string `form:"email_code" binding:"required,len=6"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,password"`
}

type Oauth struct {
	frontend.Abstract
	UserOauth
}

func (c Oauth) SendRegisterEmailCode(ctx *gin.Context) {
	registerEmailRequest := RegisterEmailRequest{}
	if !c.Validate(ctx, &registerEmailRequest) {
		return
	}

	user := logic.Logic.User.GetByEmail(registerEmailRequest.Email)
	if user != nil {
		c.JsonResponseWithServerError(ctx, exception.NewResponseError("该email已存在"))
		return
	}

	err := logic.Logic.Email.SendVerifyCode(ctx.Request.Context(), registerEmailRequest.Email, "register")
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}

func (c Oauth) Register(ctx *gin.Context) {
	registerRequest := RegisterRequest{}
	if !c.Validate(ctx, &registerRequest) {
		return
	}
	//if words := logic.Logic.SysSensitiveWordsLogic.GetSensitiveWord(registerRequest.Email); len(words) > 0 {
	//	oauthController.JsonResponseWithServerError(ctx, exception.NewLogicError("用户名包含敏感字符 "+strings.Join(words, ",")))
	//	return
	//}
	err := logic.Logic.Email.VerifyCode(ctx.Request.Context(), registerRequest.Email, registerRequest.EmailVerifyCode, "register")
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	user, err := logic.Logic.User.CreateUser(registerRequest.Email, registerRequest.Email, registerRequest.Password)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.triggerLogin(ctx, user)
}

func (c Oauth) Login(ctx *gin.Context) {
	loginRequest := LoginRequest{}
	if !c.Validate(ctx, &loginRequest) {
		return
	}

	user, err := logic.Logic.User.GetByEmailAndPwd(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}
	if user.IsDisable() {
		c.JsonResponseWithServerError(ctx, exception.NewResponseError("用户已被禁用"))
		return
	}

	c.triggerLogin(ctx, user)
}

func (c Oauth) Logout(ctx *gin.Context) {
	err := c.RemoveUserFromSession(ctx)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}

func (c Oauth) triggerLogin(ctx *gin.Context, user *model.User) {
	user.LastIp = ctx.ClientIP()
	user.LatestVisit = time.Now().Format(model.TimeFormat)
	err := logic.Logic.User.UpdateUser(user)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}
	err = logic.Logic.User.ResetUserCache(ctx.Request.Context(), user)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	err = c.SaveUserToSession(ctx, user)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonSuccessResponse(ctx)
}
