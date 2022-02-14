package config

import (
	"log"
	"os"
)

var Config *config

type config struct {
	Mysql *Mysql
}

type Mysql struct {
	DBTCPHost              string
	Port                   string
	User                   string
	Password               string
	Database               string
	InstanceConnectionName string
}

func InitConfig() {
	Config = &config{
		Mysql: &Mysql{
			Port:                   mustGetenv("DB_PORT"),
			User:                   mustGetenv("DB_USER"),
			Database:               mustGetenv("DB_NAME"),
			Password:               os.Getenv("DB_PASSWORD"),
			DBTCPHost:              os.Getenv("DB_TCP_HOST"),
			InstanceConnectionName: os.Getenv("INSTANCE_CONNECTION_NAME"),
		},
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
