package context

import (
	"github.com/go-tech-lab/framework_common/src/routine"
	"path/filepath"
	"runtime"
	"strings"
)

type FuncContext struct {
	//函数名
	FuncName string
	//文件名
	FileName string
	//代码行
	Line int
	//RoutineNo
	routineNo int64
}


// GetFuncContext 查找函数执行上下文信息
// param skip 忽略函数调用栈层次
// return  文件名，代码行数，函数名
func GetFuncContext(skip int) *FuncContext {
	fileName, line, funcName := "???", 0, "???"
	pc, fileName, line, ok := runtime.Caller(skip)

	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		fileName = filepath.Base(fileName)
	}
	routineNo := routine.RoutineUniqNo()
	return &FuncContext{FuncName: funcName, FileName: fileName, Line: line, routineNo: routineNo}
}
