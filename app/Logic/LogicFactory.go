package logic

type LogicFactory struct {
	AppLogic
	DeviceLogic
	DeviceGatewayLogic
	DeviceOperateLogic
	DeviceReportLogic
	EmqxLogic
	MessageLogic
	UserLogic
	SysSensitiveWordsLogic
	EmailLogic
	AttachLogic
}

var Logic = new(LogicFactory)
