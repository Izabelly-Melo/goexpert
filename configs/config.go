package configs

import (
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

func NewConfig() *conf {
	return cfg
}

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiration string `mapstructure:"JWT_EXPIRATION"`
	JwtExpiresIn  int
	TokenAuth     *jwtauth.JWTAuth
}

func (c *conf) GetDBDrive() string {
	return c.DBDriver
}

// ymal, json, env, flags, etc.
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("api_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	expiresIn, convErr := strconv.Atoi(cfg.JWTExpiration)
	if convErr != nil {
		expiresIn = 300
	}
	cfg.JwtExpiresIn = expiresIn

	return cfg, err
}
