package configs

import (
	"os"
	"strconv"
)

type ConfigList struct {
	Env                 string
	DBHost              string
	DBPort              int
	DBDriver            string
	DBName              string
	DBUser              string
	DBPassword          string
	APICorsAllowOrigins []string
}

func (c *ConfigList) IsDevelopment() bool {
	return c.Env == "development"
}

var Config ConfigList

func LoadEnv() error {
	DBPort, error := strconv.Atoi(GetnEnvDefault("MYSQL_PORT", "3306"))
	if error != nil {
		return error
	}

	Config = ConfigList{
		Env:                 GetnEnvDefault("APP_ENV", "development"),
		DBDriver:            GetnEnvDefault("DB_DRIVER", "mysql"),
		DBHost:              GetnEnvDefault("DB_HOST", "0.0.0.0"),
		DBPort:              DBPort,
		DBUser:              GetnEnvDefault("DB_USER", "app"),
		DBPassword:          GetnEnvDefault("DB_PASSWORD", "password"),
		DBName:              GetnEnvDefault("DB_NAME", "api_datebase"),
		APICorsAllowOrigins: []string{"http://0.0.0.0:8001"},
	}

	return nil
}

func init() {
	if err := LoadEnv(); err != nil {
		panic(err)
	}
}

func GetnEnvDefault(key, defVal string) string {
	val, err := os.LookupEnv(key)
	if !err {
		return defVal
	}
	return val
}
