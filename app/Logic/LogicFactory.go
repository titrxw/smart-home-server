package logic

type LogicFactory struct {
	AppLogic
	DeviceLogic
	DeviceOperateLogic
	DeviceReportLogic
	EmqxLogic
	UserLogic
	SysSensitiveWordsLogic
}

var Logic = new(LogicFactory)
