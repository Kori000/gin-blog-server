package config

import "fmt"

type System struct {
	Host string `yarml:"host"`
	Port int    `yarml:"port"`
	Env  string `yarml:"env"`
}

func (c System) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
