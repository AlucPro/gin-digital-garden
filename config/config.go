package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	DbType  string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
}

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
