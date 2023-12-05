package main

type config struct {
	Update struct {
	} `cmd:"" help:"Update Exchange Rates"`
	API struct {
		Port int32 `flag:"" default:"8080" env:"PORT"`
	} `cmd:"" help:"Run api"`
	DbURL string `flag:"" env:"DB_URL"`
}
