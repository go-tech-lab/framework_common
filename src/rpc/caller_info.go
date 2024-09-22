package rpc

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/micro/go-micro/metadata"
)

var upstreamServiceInfoTransportKey = "upstream-service-info"

var activeUpstreamServiceInfoKey = upstreamServiceInfoKey{}

type upstreamServiceInfoKey struct {
}
type CallerServiceInfo struct {
	Env                  string
	Region               string
	CallerServiceName    string
	CallerServiceMethod  string
	CallerServiceGroup   string
	CallerServiceProject string
	CallerServiceProduct string
	CallerServiceBizName string
	CallerServiceHost    string
	CallerServiceIdc     string
}

//ExtractUpstreamServiceInfo 从ctx中提取上游调用方信息
func ExtractUpstreamServiceInfo(ctx context.Context) *CallerServiceInfo {
	callerServiceInfo := &CallerServiceInfo{}
	data, ok := metadata.FromContext(ctx)
	if !ok {
		return callerServiceInfo
	}
	upstreamServiceInfoJson, _ := data[upstreamServiceInfoTransportKey]
	if len(upstreamServiceInfoJson) <= 0 {
		return callerServiceInfo
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(upstreamServiceInfoJson), callerServiceInfo)
	if err != nil {
		return callerServiceInfo
	}
	return callerServiceInfo
}
