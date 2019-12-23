package log

var (
	_Logger ILogger
	_Level  int
)

func GetLevel() int {
	//如果logger还没有初始化，就进行初始化操作
	if _Logger == nil {

		//TODO:初始化代码操作
	}

	return _Level
}

func getLogger() ILogger {
	if _Logger == nil {

	}

	return _Logger
}

func Debug(category, msg string, args ...interface{}) {

	getLogger().Debug(category, msg, args)
}

func Info(category, msg string, args ...interface{}) {
	getLogger().Info(category, msg, args)
}

func Warning(category, msg string, args ...interface{}) {
	getLogger().Warning(category, msg, args)
}

func Error(category, msg string, args ...interface{}) {
	getLogger().Error(category, msg, args)

}
func Fatal(category, msg string, args ...interface{}) {
	getLogger().Fatal(category, msg, args)

}
func Raw(category, msg string, args ...interface{}) {
	getLogger().Raw(category, msg, args)

}
