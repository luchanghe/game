package configManage

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once

var v *viper.Viper

func GetConfig() *viper.Viper {
	once.Do(func() {
		path, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		v = viper.New()
		v.SetConfigName("config")
		v.AddConfigPath(filepath.Clean(path + "/config"))
		err = v.ReadInConfig()
		if err != nil {
			panic(err.Error())
		}
	})
	return v
}
