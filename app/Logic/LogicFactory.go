package logic

type LogicFactory struct {
	AppLogic
	DeviceLogic
	DeviceOperateLogic
	DeviceReportLogic
	EmqxLogic
	UserLogic
	SysSensitiveWordsLogic
	EmailLogic
}

var Logic = new(LogicFactory)
