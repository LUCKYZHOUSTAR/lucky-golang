package config

import (
	"go/frame/utility/app"
	"os"
	"path/filepath"
	"runtime"
)

//　　Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头
var (
	ConfigDir       string
	globalConfigDir string
)

const (
	globalConfigDirEnvKey = "LUCKY_GLOBAL_CONFIG_PATH"
)

//GetConfigPath 获取全局配置信息
func GetConfigPath(fileName string) string {

	return filepath.Join(getConfigDir(), fileName)

}

//GetConfigPaths 获取指定模式下面的配置列表信息s
func GetConfigPaths(fileNamePattern string) []string {
	pattern := filepath.Join(getConfigDir(), fileNamePattern)
	paths, err := filepath.Glob(pattern)

	if err != nil {
		panic(err)
	}

	return paths

}

//FindConfigPath 查找某个目录是否存在
func FindConfigPath(fileName string, exts ...string) string {

	for _, ext := range exts {
		p := filepath.Join(getConfigDir(), fileName+ext)
		//判断是不是一个目录
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return ""
}

//GetGlobalConfirPath 获取全局配置的地址信息
func GetGlobalConfirPath(fileName string) string {
	return filepath.Join(getGlobalConfigDir(), fileName)
}

//拼接当前的目录操作，当前的文件绝对路径，拼接config
func getConfigDir() string {
	if ConfigDir == "" {
		folder, err := app.GetAppFolder()
		if err != nil {
			//挂断当前进程
			panic(err)
		}

		ConfigDir = filepath.Join(folder, "config")
	}
	return ConfigDir
}

//获取全局的默认的配置路径信息
func getGlobalConfigDir() string {

	if globalConfigDir == "" {
		//获取环境变量的默认地址信息
		globalConfigDir = os.Getenv(globalConfigDirEnvKey)
		//如果环境变量的默认地址的信息为空的话，通过不同的操作系统来指定不同的目录结构
		if globalConfigDir == "" {

			switch runtime.GOOS {
			case "windows":
				globalConfigDir = `D:\etc`
			case "darwin":
				globalConfigDir = `/etc/lucky`
			default:
				globalConfigDir = `/home/lucky/etc`
			}
		}
	}

	return ""

}
