package user

import (
	exception "github.com/titrxw/smart-home-server/app/Exception"
	model "github.com/titrxw/smart-home-server/app/Model"
	"time"

	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"

	"github.com/gin-gonic/gin"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type RegisterEmailRequest struct {
	base.RequestAbstract
	Email string `form:"email" binding:"required,email"`
}

type RegisterRequest struct {
	base.RequestAbstract
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,password"`
	EmailVerifyCode string `form:"email_code" binding:"required,len=6"`
}

type LoginRequest struct {
	base.RequestAbstract
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,password"`
}

type OauthController struct {
	frontend.ControllerAbstract
	UserOauth
}

func (oauthController OauthController) SendRegisterEmailCode(ctx *gin.Context) {
	registerEmailRequest := RegisterEmailRequest{}
	if !oauthController.ValidateFormPost(ctx, &registerEmailRequest) {
		return
	}

	user := logic.Logic.UserLogic.GetByEmail(registerEmailRequest.Email)
	if user != nil {
		oauthController.JsonResponseWithServerError(ctx, exception.NewLogicError("该email已存在"))
		return
	}

	err := logic.Logic.EmailLogic.SendVerifyCode(ctx.Request.Context(), registerEmailRequest.Email, "register")
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	oauthController.JsonSuccessResponse(ctx)
}

func (oauthController OauthController) Register(ctx *gin.Context) {
	registerRequest := RegisterRequest{}
	if !oauthController.ValidateFormPost(ctx, &registerRequest) {
		return
	}
	//if words := logic.Logic.SysSensitiveWordsLogic.GetSensitiveWord(registerRequest.Email); len(words) > 0 {
	//	oauthController.JsonResponseWithServerError(ctx, exception.NewLogicError("用户名包含敏感字符 "+strings.Join(words, ",")))
	//	return
	//}
	err := logic.Logic.EmailLogic.VerifyCode(ctx.Request.Context(), registerRequest.Email, registerRequest.EmailVerifyCode, "register")
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	user, err := logic.Logic.UserLogic.CreateUser(registerRequest.Email, registerRequest.Email, registerRequest.Password)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	oauthController.triggerLogin(ctx, user)
}

func (oauthController OauthController) Login(ctx *gin.Context) {
	loginRequest := LoginRequest{}
	if !oauthController.ValidateFormPost(ctx, &loginRequest) {
		return
	}

	user, err := logic.Logic.UserLogic.GetByEmailAndPwd(loginRequest.Email, loginRequest.Password)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}
	if user.IsDisable() {
		oauthController.JsonResponseWithServerError(ctx, exception.NewLogicError("用户已被禁用"))
		return
	}

	oauthController.triggerLogin(ctx, user)
}

func (oauthController OauthController) Logout(ctx *gin.Context) {
	err := oauthController.RemoveUserFromSession(ctx)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}

	oauthController.JsonSuccessResponse(ctx)
}

func (oauthController OauthController) triggerLogin(ctx *gin.Context, user *model.User) {
	user.LastIp = ctx.ClientIP()
	user.LatestVisit = time.Now().Format(model.TimeFormat)
	err := logic.Logic.UserLogic.UpdateUser(user)
	if err != nil {
		oauthController.JsonResponseWithServerError(ctx, err)
		return
	}
	err = logic.Logic.UserLogic.ResetUserCache(ctx.Request.Context(), user)
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
