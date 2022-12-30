package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jordan-wright/email"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	helper "github.com/titrxw/smart-home-server/app/Helper"
	model "github.com/titrxw/smart-home-server/app/Model"
	utils "github.com/titrxw/smart-home-server/app/Utils"
	"github.com/titrxw/smart-home-server/config"
	"net/smtp"
	"time"
)

type EmailLogic struct {
	LogicAbstract
}

const SERVICE_EMAIL_CODE = "service:email:code:email:%s:verify_type:%s"

func (emailLogic EmailLogic) formatCacheKey(email string, verifyType string) string {
	return fmt.Sprintf(SERVICE_EMAIL_CODE, email, verifyType)
}

func (emailLogic EmailLogic) SendEmail(userEmail string, htmlContent string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("发件人 <%s>", config.GConfig.Email.FromUserName)
	e.To = []string{userEmail}
	e.HTML = []byte(htmlContent)
	return e.Send(config.GConfig.Email.Host+":"+config.GConfig.Email.Port, smtp.PlainAuth(config.GConfig.Email.Identify, config.GConfig.Email.UserName, config.GConfig.Email.Password, config.GConfig.Email.Host))
}

func (emailLogic EmailLogic) SendVerifyCode(ctx context.Context, email string, verifyType string) error {
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

	err := emailLogic.SendEmail(email, content)
	if err != nil {
		return err
	}

	return emailLogic.GetDefaultRedis().Set(ctx, emailLogic.formatCacheKey(email, verifyType), code, time.Second*300).Err()
}

func (emailLogic EmailLogic) VerifyCode(ctx context.Context, email string, emailVerifyCode string, verifyType string) error {
	cacheKey := emailLogic.formatCacheKey(email, verifyType)

	return utils.RetryLimit{}.Try(ctx, emailLogic.GetDefaultRedis(), func(context context.Context, client *redis.Client, curNum int64) error {
		result := client.Get(context, cacheKey)
		if result.Err() != nil && result.Err() != redis.Nil {
			return result.Err()
		}
		if emailVerifyCode == result.Val() {
			return client.Del(context, cacheKey).Err()
		}

		return exception.NewLogicError("email验证码错误")
	}, cacheKey, 5, time.Second*300)
}
