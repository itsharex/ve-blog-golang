// Code generated by goctl. DO NOT EDIT.
// Source: config.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/logic/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

type ConfigRpcServer struct {
	svcCtx *svc.ServiceContext
	configrpc.UnimplementedConfigRpcServer
}

func NewConfigRpcServer(svcCtx *svc.ServiceContext) *ConfigRpcServer {
	return &ConfigRpcServer{
		svcCtx: svcCtx,
	}
}

// 保存配置
func (s *ConfigRpcServer) SaveConfig(ctx context.Context, in *configrpc.SaveConfigReq) (*configrpc.SaveConfigResp, error) {
	l := configrpclogic.NewSaveConfigLogic(ctx, s.svcCtx)
	return l.SaveConfig(in)
}

// 查询配置
func (s *ConfigRpcServer) FindConfig(ctx context.Context, in *configrpc.FindConfigReq) (*configrpc.FindConfigResp, error) {
	l := configrpclogic.NewFindConfigLogic(ctx, s.svcCtx)
	return l.FindConfig(in)
}
