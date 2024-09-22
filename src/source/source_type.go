package source

const (
	SourceType_None   = SourceType("")
	SourceType_Normal = SourceType("normal")
	SourceType_Stress_Read = SourceType("stress_read")
	SourceType_Stress_Write = SourceType("stress_write")
)

//请求来源类型
type SourceType string

//获取当前上下文（一般指处理请求/消息的协程上下文）的SourceType
func GetSourceType() SourceType {
	if sourceTypeGetter != nil {
		return sourceTypeGetter()
	}
	return SourceType_None
}

//设置 sourceTypeGetter的函数实例
func InjectSourceTypeGetter(getter SourceTypeGetterFunc) {
	sourceTypeGetter = getter
}

//获取当前上下文（一般指处理请求/消息的协程上下文）的SourceType的函数类型
type SourceTypeGetterFunc func() SourceType

//获取当前上下文（一般指处理请求/消息的协程上下文）的SourceType的函数实例变量
var sourceTypeGetter SourceTypeGetterFunc
