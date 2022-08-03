package logic

type LogicFactory struct {
	DeviceLogic
	DeviceOperateLogic
	DeviceReportLogic
	EmqxLogic
	UserLogic
	SysSensitiveWordsLogic
}

var Logic = new(LogicFactory)
