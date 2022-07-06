package config

type Config struct {
	Base  Base  `mapstructure:"base"`
	JWT   JWT   `mapstructure:"jwt"`
	MySQL MySQL `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
}
