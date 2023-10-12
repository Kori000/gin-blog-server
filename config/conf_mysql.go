package config

import "strconv"

type Mysql struct {
	Host      string `yaml:"host" default:"127.0.0.1"`
	Port      int    `yaml:"port" default:"3306"`
	DB        string `yaml:"db" default:"gvb_db"`
	User      string `yaml:"user" default:"root"`
	Password  string `yaml:"password" default:"qq123123"`
	Log_Level string `yaml:"log_level" default:"dev"`
	Config    string `yaml:"config" default:"charset=utf8mb4&parseTime=True&loc=Local"` // 高级配置; 例如 charset
}

func (m *Mysql) Dsn() string {
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
	return dsn
}
