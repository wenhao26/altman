package config

type Base struct {
	Name       string `mapstructure:"name"`
	ListenPort int    `mapstructure:"listen_port"`
}
