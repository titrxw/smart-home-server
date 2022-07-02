package user

import (
	"errors"
	"strings"

	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"

	"github.com/gin-gonic/gin"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type RegisterRequest struct {
	base.RequestAbstract
	Username string `form:"user_name" binding:"required,user_name"`
	Mobile   string `form:"mobile" binding:"required,mobile"`
	Password string `form:"password" binding:"required,password"`
}

type LoginRequest struct {
	base.RequestAbstract
	Mobile   string `form:"mobile" binding:"required,mobile"`
	Password string `form:"password" binding:"required,password"`
}

type OauthController struct {
	frontend.ControllerAbstract
	UserOauth
}

func (oauthController OauthController) Register(ctx *gin.Context) {
	registerRequest := RegisterRequest{}
	if !oauthController.ValidateFormPost(ctx, &registerRequest) {
		return
	}
	if words := logic.Logic.SysSensitiveWordsLogic.GetSensitiveWord(registerRequest.Username); len(words) > 0 {
		oauthController.JsonResponseWithServerError(ctx, errors.New("用户名包含敏感字符 "+strings.Join(words, ",")))
		return
	}

	user, err := logic.Logic.UserLogic.CreateUser(registerRequest.Username, registerRequest.Mobile, registerRequest.Password)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	err = oauthController.SaveUserToSession(ctx, user)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	oauthController.JsonSuccessResponse(ctx)
}

func (oauthController OauthController) Login(ctx *gin.Context) {
	loginRequest := LoginRequest{}
	if !oauthController.ValidateFormPost(ctx, &loginRequest) {
		return
	}

	user, err := logic.Logic.UserLogic.GetByMobileAndPwd(loginRequest.Mobile, loginRequest.Password)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}
	if user.IsDisable() {
		oauthController.JsonResponseWithServerError(ctx, "用户已被禁用")
		return
	}
	user.LastIp = ctx.ClientIP()
	err = logic.Logic.UserLogic.UpdateUser(user)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	err = oauthController.SaveUserToSession(ctx, user)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	oauthController.JsonSuccessResponse(ctx)
}
