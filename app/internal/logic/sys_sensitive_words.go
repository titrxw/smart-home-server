package logic

import (
	"github.com/titrxw/smart-home-server/app/internal/repository"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/logic"
)

type SysSensitiveWords struct {
	logic.Abstract
}

const SysSensitiveWordsSettingKey = "sys:sensitive:words"

func (l SysSensitiveWords) GetSysSensitiveWords() []string {
	return repository.Repository.Setting.Get(l.GetDefaultDb(), SysSensitiveWordsSettingKey, []string{}).([]string)
}

func (l SysSensitiveWords) GetSensitiveWord(word string) []string {
	return helper.GetSensitiveWord(word, l.GetSysSensitiveWords())
}
