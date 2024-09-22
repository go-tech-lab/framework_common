package context

type RequestContext struct {
	//接口名
	Method string
	//来源ip
	ClientIp string
	//来源服务
	FromServer string
	//traceId
	TraceId string
	//用户身份
	User string
	//日志标签
	LogTags map[string]string
}
