package global

import (
	"altman/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

var (
	AmViper  *viper.Viper
	AmConfig config.Config
	DB       *gorm.DB
	Wg       sync.WaitGroup
)
