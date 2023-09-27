package config

type System struct {
	Host string `yarml:"host"`
	Port int    `yarml:"port"`
	Env  string `yarml:"env"`
}
