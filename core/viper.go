package core

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"altman/global"
)

func Viper() *viper.Viper {

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.AmConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.AmConfig); err != nil {
		fmt.Println(err)
	}

	return v
}
