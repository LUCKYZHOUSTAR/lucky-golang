package log

//接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
// 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。

const (
	logLevelDebug = iota
	logLevelInfo
	logLevelWarning
	logLevelError
	logLevelFatal
)

//interface{}代表什么类型都可以传入
type ILogger interface {
	SetLevel(level int)
	Debug(category, msg string, args ...interface{})
	Info(category, msg string, args ...interface{})
	Warning(category, msg string, args ...interface{})
	Error(category, msg string, args ...interface{})
	Fatal(category, msg string, args ...interface{})
	Raw(category, msg string, args ...interface{})
}
