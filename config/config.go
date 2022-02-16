package config

type Configuration struct {
	ControllerPath string `mapstructure:"controllerPath" json:"controllerPath" yaml:"controllerPath"`
	ModelPath      string `mapstructure:"modelPath" json:"modelPath" yaml:"modelPath"`
	EventPath      string `mapstructure:"eventPath" json:"eventPath" yaml:"eventPath"`
	ListenerPath   string `mapstructure:"listenerPath" json:"listenerPath" yaml:"listenerPath"`
	Mysql          Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

var Config *Configuration
