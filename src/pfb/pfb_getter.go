package pfb

import "context"

func GetPFB() string {
	if pfbGetter != nil {
		return pfbGetter()
	}
	return ""
}

func InjectPfbGetter(getter PfbGetterFunc) {
	pfbGetter = getter
}

//PfbGetterFunc 获取当前上下文（一般指处理请求/消息的协程上下文）的SourceType的函数类型
type PfbGetterFunc func() string

//获取当前上下文（一般指处理请求/消息的协程上下文）的SourceType的函数实例变量
var pfbGetter PfbGetterFunc

//GetPFBByNamingService 通过查询naming service获取PFB
func GetPFBByNamingService(ctx context.Context, platformUserID string) string {
	if pfbGetterByNamingService != nil {
		return pfbGetterByNamingService(ctx, platformUserID)
	}
	return ""
}

func InjectPfbGetterByNamingService(getter PfbGetterByNamingServiceFunc) {
	pfbGetterByNamingService = getter
}

//PfbGetterByNamingServiceFunc 获取当前上下文（一般指处理请求/消息的协程上下文）的PFB的函数类型
type PfbGetterByNamingServiceFunc func(ctx context.Context, platformUserID string) string

//获取当前上下文（一般指处理请求/消息的协程上下文）的PFB的函数实例变量
var pfbGetterByNamingService PfbGetterByNamingServiceFunc
