package rpc

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// IRpcWrapper rpc wrapper: 增强rpc功能的wrapper
type IRpcWrapper interface {
	// ClientWrapper rpc 客户端 wrapper
	ClientWrapper() client.CallWrapper
	// ServerWrapper rpc 服务端 wrapper
	ServerWrapper() server.HandlerWrapper
}
