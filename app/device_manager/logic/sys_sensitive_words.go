package logic

import (
	"github.com/titrxw/smart-home-server/app/common/helper"
	"github.com/titrxw/smart-home-server/app/device_manager/repository"
)

type SysSensitiveWords struct {
	Abstract
}

const SysSensitiveWordsSettingKey = "sys:sensitive:words"

func (l SysSensitiveWords) GetSysSensitiveWords() []string {
	return repository.Repository.Setting.Get(l.GetDefaultDb(), SysSensitiveWordsSettingKey, []string{}).([]string)
}

func (l SysSensitiveWords) GetSensitiveWord(word string) []string {
	return helper.GetSensitiveWord(word, l.GetSysSensitiveWords())
}
