package database

import (
	"database/sql"
	"log"

	"github.com/BurntSushi/toml"
)

type databaseConfig struct {
	InUse      string     `toml:"inuse"`
	Postgresql postgresql `toml"postgresql"`
}

//package wise global variables
var (
	aDatabaseConfig                 databaseConfig
	aRelationalDatabaseManipulation relationalDatabaseManipulation
	conn                            *sql.DB
)

func init() {
	_, err := toml.DecodeFile("../config/DatabaseRelated.toml", &aDatabaseConfig)
	if err != nil {
		log.Fatal("Fail to load Database related configurations.")
	}
	log.Println(aDatabaseConfig)

	switch aDatabaseConfig.InUse {
	case "postgresql":
		aRelationalDatabaseManipulation = aDatabaseConfig.Postgresql
	default:
		log.Fatal("Cannot find match database with configuration.")
	}

	conn = aRelationalDatabaseManipulation.Connect()
}

func Get() relationalDatabaseManipulation {
	return aRelationalDatabaseManipulation
}
