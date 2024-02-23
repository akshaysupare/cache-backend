package config

import "github.com/dgraph-io/ristretto"

const (
	Port = "9000"
)

var (
	Cache *ristretto.Cache
	Err   error
)
