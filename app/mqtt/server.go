package mqtt

import (
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type Server struct {
	server.Server
	config           *viper.Viper
	subscribeManager *SubscribeManager
}

func NewMqttSubServer(config *viper.Viper, subscribeManager *SubscribeManager) *Server {
	return &Server{
		config:           config,
		subscribeManager: subscribeManager,
	}
}

func (s Server) GetSubscribeManager() *SubscribeManager {
	return s.subscribeManager
}

func (s Server) GetServerName() string {
	return "mqtt-sub"
}

func (s Server) GetOptions() map[string]string {
	return map[string]string{
		"Host": s.config.GetString("server.mqtt.host"),
		"Port": s.config.GetString("server.mqtt.port"),
	}
}

func (s Server) Start() {
	s.subscribeManager.Start(
		s.config.GetString("server.mqtt.host"),
		s.config.GetString("server.mqtt.port"),
		s.config.GetString("server.mqtt.user_name"),
		s.config.GetString("server.mqtt.password"),
	)
}
