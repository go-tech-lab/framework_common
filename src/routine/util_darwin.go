package routine

import (
	"runtime"
	"strings"
)

/**
* 协程相关工具方法
* author changlin.guo
* date 2019-08-22
 */


// RoutineUniqNo 获取gorouine的协程唯一编号（不是 goroutine id）
func RoutineUniqNo() int64 {
	if strings.Contains(runtime.GOOS, "darwin") {
		return NotLinuxGetGID()
	} else {
		panic("your os is not darwin123, process exit")
	}
}

/**
* 获取gorouine的退出回调函数
 */
func RegisterExitCallback(callback func()) {
	callback()
}
