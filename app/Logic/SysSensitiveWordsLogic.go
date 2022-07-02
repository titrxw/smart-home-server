package logic

import (
	helper "github.com/titrxw/smart-home-server/app/Helper"
	repository "github.com/titrxw/smart-home-server/app/Repository"
)

type SysSensitiveWordsLogic struct {
	LogicAbstract
}

const SYS_SENSITIVE_WORDS_SETTING_KEY = "sys:sensitive:words"

func (swLogic SysSensitiveWordsLogic) GetSysSensitiveWords() []string {
	return repository.Repository.SettingRepository.Get(swLogic.GetDefaultDb(), SYS_SENSITIVE_WORDS_SETTING_KEY, []string{}).([]string)
}

func (swLogic SysSensitiveWordsLogic) GetSensitiveWord(word string) []string {
	return helper.GetSensitiveWord(word, swLogic.GetSysSensitiveWords())
}
