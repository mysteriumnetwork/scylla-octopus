package entity

import (
	"github.com/caarlos0/env/v6"
)
// NodeBinaries paths to scylladb executables on database node
type NodeBinaries struct {
	Cqlsh    string
	Nodetool string
}

// Credentials scylladb credentials
type Credentials struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
}
