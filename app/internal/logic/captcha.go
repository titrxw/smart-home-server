package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/titrxw/smart-home-server/app/pkg/captcha"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
)

type Captcha struct {
	logic.Abstract
}

func (l *Captcha) GenerateCaptcha(ctx context.Context) (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, captcha.NewDefaultRedisStore(l.GetDefaultRedis(), ctx))
	return cp.Generate()
}

func (l *Captcha) ValidateCaptcha(ctx context.Context, captchaId string, captchaValue string) bool {
	return captcha.NewDefaultRedisStore(l.GetDefaultRedis(), ctx).Verify(captchaId, captchaValue, true)
}
