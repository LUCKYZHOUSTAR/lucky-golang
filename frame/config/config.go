package config

import (
	"go/frame/utility/app"
	"path/filepath"
)

//　　Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头
var (
	ConfigDir       string
	globalConfigDir string
)

const (
	globalConfigDirEvnKey = "LUCKY_GLOBAL_CONFIG_PATH"
)

func GetConfigDir() string {
	if ConfigDir == "" {
		folder, err := app.GetAppFolder()
		if err != nil {
			panic(err)
		}

		ConfigDir = filepath.Join(folder, "config")
	}
	return ConfigDir

}
