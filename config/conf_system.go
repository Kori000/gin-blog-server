package config

import "fmt"

type System struct {
	Host string `yarml:"host" default:"0.0.0.0"`
	Port int    `yarml:"port" default:"4000"`
	Env  string `yarml:"env" default:"release"`
}

func (c System) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
