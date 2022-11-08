package config

import (
	"flag"
	"os"
)

type Config struct {
    
    username   string
	password   string
    
}

func Get() *Config {
    conf := &Config{}

    
    flag.StringVar(&conf.username, "username", os.Getenv("USERNAME"), "twilio Username")
	flag.StringVar(&conf.password, "password", os.Getenv("PASSWORD"), "twilio Password")
    
	flag.Parse()

	return conf
}


func (c *Config) GetUsername() string {
	return c.username
}

func (c *Config) GetPassword() string {
	return c.password
}

