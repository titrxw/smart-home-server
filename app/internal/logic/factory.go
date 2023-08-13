package logic

type Factory struct {
	App
	Device
	DeviceGateway
	DeviceOperate
	DeviceReport
	Emqx
	Message
	User
	SysSensitiveWords
	Email
	Attach
	Captcha
}

var Logic = new(Factory)
