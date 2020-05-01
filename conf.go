package main

import "github.com/spider-pigs/envlookup"

type config struct {
	Port int
}

func newConf() *config {
	conf := &config{
		Port: envlookup.MustInt(envlookup.Int("PORT", 8080)),
	}
	return conf
}
