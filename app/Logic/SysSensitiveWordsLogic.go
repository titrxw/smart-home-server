package logic

import (
	helper "github.com/titrxw/smart-home-server/app/Helper"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

type SysSensitiveWordsLogic struct {
	LogicAbstract
}

const SYS_SENSITIVE_WORDS_SETTING_KEY = "sys:sensitive:words"

func (this SysSensitiveWordsLogic) GetSysSensitiveWords() []string {
	return repository.Repository.SettingRepository.Get(this.GetDefaultDb(), SYS_SENSITIVE_WORDS_SETTING_KEY, []string{}).([]string)
}

func (this SysSensitiveWordsLogic) GetSensitiveWord(word string) []string {
	return helper.GetSensitiveWord(word, this.GetSysSensitiveWords())
}
