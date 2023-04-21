package logic

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/redis/go-redis/v9"
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/device_manager/utils"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"net/smtp"
	"time"
)

type Email struct {
	Abstract
}

const ServiceEmailCode = "service:email:code:email:%s:verify_type:%s"

func (l Email) formatCacheKey(email string, verifyType string) string {
	return fmt.Sprintf(ServiceEmailCode, email, verifyType)
}

func (l Email) SendEmail(userEmail string, htmlContent string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("发件人 <%s>", app.GApp.GetConfig().GetString("email.from_user_name"))
	e.To = []string{userEmail}
	e.HTML = []byte(htmlContent)
	return e.Send(
		app.GApp.GetConfig().GetString("email.host")+":"+app.GApp.GetConfig().GetString("email.port"),
		smtp.PlainAuth(
			app.GApp.GetConfig().GetString("email.identify"),
			app.GApp.GetConfig().GetString("email.user_name"),
			app.GApp.GetConfig().GetString("email.password"),
			app.GApp.GetConfig().GetString("email.host"),
		),
	)
}

func (l Email) SendVerifyCode(ctx context.Context, email string, verifyType string) error {
	code := helper.RandomNumber(6)

	content := fmt.Sprintf(`
	<html><div>
		<div>
			您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div></html>
	`, time.Now().Format(model.TimeFormat), code)

	err := l.SendEmail(email, content)
	if err != nil {
		return err
	}

	return l.GetDefaultRedis().Set(ctx, l.formatCacheKey(email, verifyType), code, time.Second*300).Err()
}

func (l Email) VerifyCode(ctx context.Context, email string, emailVerifyCode string, verifyType string) error {
	cacheKey := l.formatCacheKey(email, verifyType)

	return utils.RetryLimit{}.Try(ctx, l.GetDefaultRedis(), func(context context.Context, client redis.Cmdable, curNum int64) error {
		result := client.Get(context, cacheKey)
		if result.Err() != nil && result.Err() != redis.Nil {
			return result.Err()
		}
		if emailVerifyCode == result.Val() {
			return client.Del(context, cacheKey).Err()
		}

		return exception.NewResponseError("email验证码错误")
	}, cacheKey, 5, time.Second*300)
}
