package manage

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

type ConfigManage struct {
	*viper.Viper
}

var configManageOnce sync.Once
var configManageCache *ConfigManage

func GetConfigManage() *ConfigManage {
	configManageOnce.Do(func() {
		path, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		v := viper.New()
		v.SetConfigName("config")
		v.AddConfigPath(filepath.Clean(path + "/config"))
		err = v.ReadInConfig()
		if err != nil {
			panic(err.Error())
		}
		configManageCache = &ConfigManage{Viper: v}
	})
	return configManageCache

}
