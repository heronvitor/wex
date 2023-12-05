package main

import "time"

type config struct {
	Update struct {
		UpdateInterval time.Duration `flag:"" default:"24h"`
		RetryInterval  time.Duration `flag:"" default:"1h"`
	} `cmd:"" help:"Update Exchange Rates"`
	API struct {
		Port int32 `flag:"" default:"8080" env:"PORT"`
	} `cmd:"" help:"Run api"`
	DbURL string `flag:"" env:"DB_URL"`
}
