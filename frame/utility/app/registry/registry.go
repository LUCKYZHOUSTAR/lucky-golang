package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"mtime.com/framework/config"
	"mtime.com/framework/etcd/etcdx"
	"mtime.com/framework/log"
	"mtime.com/framework/utility/convert"
)

// Register 注册应用到配置中心
func Register(fn func() map[string]interface{}) {
	// 过期时间(单位秒)
	const ttl = 60

	c := etcdx.GetClient()

	appName := config.GetAppConfig().App.AppName
	if appName == "" {
		log.Warning("app", "register app failed: app_name is not configured")
		return
	}

	ip := config.GetGlobalConfig().RPCRegisterIP
	if appName == "" {
		log.Warning("app", "register app failed: ip is not configured")
		return
	}

	props := make(map[string]interface{})
	if fn == nil {
		props = make(map[string]interface{})
		props["version"] = config.GetAppConfig().App.Version
		props["debug"] = config.GetAppConfig().App.Debug
		props["start_time"] = convert.TimeToString(time.Now())
		props["process"] = os.Getpid()
	} else {
		props = fn()
	}

	go func() {
		for {
			key := fmt.Sprintf("app/%s/%s", appName, ip)
			value, _ := json.Marshal(props)
			_, err := c.Set(key, string(value), ttl+5)
			if err == nil {
				log.Info("app", "register app success: %s", value)
			} else {
				log.Error("app", "register app failed: %v", err)
			}

			time.Sleep(time.Second * ttl)
		}
	}()

}
