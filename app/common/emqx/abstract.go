package emqx

import (
	kernel "github.com/titrxw/emqx-sdk/src/Kernel"
)

type Abstract struct {
	EmqxClient *kernel.EmqxClient
}

func (s *Abstract) getEmqxClient() *kernel.EmqxClient {
	return s.EmqxClient
}
