package config

import (
	"flag"
	"os"
)

var (
	debug         = flag.Bool("debug", os.Getenv("DEBUG") == "True", "enable debugging")
	MongoPassword = flag.String("MongoPassword", os.Getenv("MONGO_INITDB_ROOT_PASSWORD"), "The database password")
	MongoAddress  = flag.String("MongoAddress", os.Getenv("MONGO_ADDRESS"), "The database server")
	MongoUser     = flag.String("MongoUser", os.Getenv("MONGO_INITDB_ROOT_USERNAME"), "The database user")
	MongoDatabase = flag.String("MongoDatabase", os.Getenv("MONGO_DATABASE"), "The database chema")
	LogPath       = flag.String("LogPath", os.Getenv("LOG_PATH"), "The log path")
)
