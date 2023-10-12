package config

type Logger struct {
	Level        string `yaml:"level" default:"info"`
	Prefix       string `yaml:"prefix" default:"[gvb]"`
	Director     string `yaml:"director" default:"log"`
	Showline     bool   `yaml:"show_line" default:"true"`
	Loginconsole bool   `yaml:"log_in_console" default:"true"`
}
