package lib

import (
	"fmt"
	"strings"
)

// Site struct
type Site struct {
	Name     string
	ID       int
	Bindings []string
	State    string
	Path     string
	Proto    string
}

func (s Site) String() string {
	return fmt.Sprintf("%s,%s,%s", s.Path, s.Name, strings.Join(s.Bindings, ";"))
}
