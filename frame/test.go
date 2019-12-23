package main

import (
	"fmt"
	"go/frame/config"
	"go/frame/errors"
	"path/filepath"
	"time"
)

func main() {

	// dd:=config.GetConfigDir()

	// fmt.Println(dd)

	paths, err := filepath.Glob("go")

	if err != nil {
		panic(err)
	}
	fmt.Println(paths)

	dd := errors.NewC("test exception", 234)

	fmt.Println(dd.Code())

	var zz = make(config.SettingMap, 45)

	zz["sdfasdf"] = "2019-12-31"

	fmt.Println(zz.Time("sdfasdf", time.Time{}))

	var aa = config.SettingMap{
		"sfdasf": "234234",
	}

	fmt.Println(aa.String("234", "234"))

}
