package util

import (
	"context"

	"github.com/micro/go-micro/metadata"
)

// InjectRequestAttr 将key-value注入到context中的metadata中，一般用于请求链路上下文传递
func InjectRequestAttr(ctx context.Context, key string, value string) context.Context {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		ctx = metadata.NewContext(ctx, metadata.Metadata{})
		data, _ = metadata.FromContext(ctx)
	}
	data[key] = value
	return ctx
}

// ExtractRequestAttr 将context中的metadata中的key-value提取出来
func ExtractRequestAttr(ctx context.Context, key string) (value string, ok bool) {
	data, existed := metadata.FromContext(ctx)
	if !existed {
		return "", false
	}
	value, ok = data[key]
	return
}

// NewContext 将context中的metadata中的key-value提取出来
func NewContext(ctx context.Context) context.Context {
	data, existed := metadata.FromContext(ctx)
	if !existed {
		data = metadata.Metadata{}
	}
	ctx = metadata.NewContext(context.TODO(), data)
	return ctx
}

const (
	throughLogFilterTagHeader = "through_log_filter_tag_header"

	callThroughMethodKey = "x-call-through-method"

	serviceMethodKey = "service-method"
)

//DisableThroughLogFilter 将LogFilterTag设置到context中，用于跨服务传递
func DisableThroughLogFilter(ctx context.Context) context.Context {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		ctx = metadata.NewContext(ctx, metadata.Metadata{})
		data, _ = metadata.FromContext(ctx)
	}
	data[throughLogFilterTagHeader] = "false"
	return ctx
}

//InjectServiceMethod 将接口名注入到ctx，方便从下游能从ctx获取到上游接口名
func InjectServiceMethod(ctx context.Context, serviceMethod string) context.Context {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		ctx = metadata.NewContext(ctx, metadata.Metadata{})
		data, _ = metadata.FromContext(ctx)
	}
	data[serviceMethodKey] = serviceMethod
	return ctx
}

//ExtractServiceMethod ctx获取到上游接口名
func ExtractServiceMethod(ctx context.Context) string {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		return ""
	}
	return data[serviceMethodKey]
}

//InjectCallThroughMethod 将接口名注入到ctx，方面从下游能从ctx获取到上游接口名
func InjectCallThroughMethod(ctx context.Context, serviceMethod string) context.Context {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		ctx = metadata.NewContext(ctx, metadata.Metadata{})
		data, _ = metadata.FromContext(ctx)
	}
	data[callThroughMethodKey] = serviceMethod
	return ctx
}

//ExtractCallThroughMethod ctx获取到上游接口名
func ExtractCallThroughMethod(ctx context.Context) string {
	data, ok := metadata.FromContext(ctx)
	if !ok {
		return ""
	}
	return data[callThroughMethodKey]
}
