package config

type MySQL struct {
	Hostname string `mapstructure:"hostname"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Prefix   string `mapstructure:"prefix"`
}
