package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"

	ServerPort = "server.port"
	ServerHost = "server.host"
)

func init() {
	// defaults.
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(ServerPort, "SER_PORT")
	_ = viper.BindEnv(ServerHost, "SER_HOST")

	viper.SetDefault(DbName, "vetPartners")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(ServerHost, "localhost")
	viper.SetDefault(ServerPort, "8080")
}
