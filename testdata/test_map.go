package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
)

type Server struct {
	Name        string        `json:"name,omitempty"`
	ID          int           `structs:"-"`
	Enabled     bool          `structs:"-"`
	http.Server `structs:"-"` // embedded
}

func main() {

	server := Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	m := structs.Map(server)
	fmt.Println(m)
}
