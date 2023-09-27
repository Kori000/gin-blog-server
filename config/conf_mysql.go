package config

import "strconv"

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DB        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Log_Level string `yaml:"log_level"`
	Config    string `yaml:"config"` // 高级配置; 例如 charset
}

func (m *Mysql) Dsn() string {
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
	return dsn
}
