package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_Level string `yaml:"log_level"`
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	Showline     bool   `yaml:"show_line"`
	Loginconsole bool   `yaml:"log_in_console"`
}

type System struct {
	Host string `yarml:"host"`
	Port int    `yarml:"port"`
	Env  string `yarml:"env"`
}
