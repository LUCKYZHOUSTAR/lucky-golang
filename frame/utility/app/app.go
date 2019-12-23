package app

import "flag"

var (
	profileFlag = flag.String("p", "", `application active profiles`)
	signalFlag  = flag.String("s", "", `send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown
		reload — reloading the configuration file`)
)

// GetProfileFlag 获取激活的环境设置参数
func GetProfileFlag() *string {
	parseFlag()
	return profileFlag
}

// GetSignalFlag 返回 signal 参数
func GetSignalFlag() *string {
	parseFlag()
	return signalFlag
}

func parseFlag() {
	if !flag.Parsed() {
		flag.Parse()
	}
}
