package logic

type LogicFactory struct {
	DeviceLogic
	DeviceOperateLogic
	EmqxLogic
	UserLogic
	SysSensitiveWordsLogic
}

var Logic = new(LogicFactory)
