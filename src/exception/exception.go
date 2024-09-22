package exception

// IException 异常接口
type IException interface {
	// GetCode get error code
	GetCode() int32

	// GetMessage error message
	GetMessage() string

	//Error error interface
	Error() string
}
