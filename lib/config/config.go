package config


type config struct {
	Environment		string		`mapstructure:"ENV"`
	LogLevel		string		`mapstructure:"LOG_LEVEL"`
	DBConnString	string		`mapstructure:"DB_CONN_STRING"`
	HTTPServerAddress string	`mapstructure:"HTTP_SERVER_ADDRESS"`
	PASETOSecret string `mapstructure:"PASETO_SYMMETRIC_KEY"`
}