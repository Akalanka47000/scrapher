package config

import (
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type Config struct {
	Port              int    `mapstructure:"PORT"`
	Host              string `mapstructure:"HOST"`
	FrontendBaseUrl   string `mapstructure:"FRONTEND_BASE_URL" validate:"url"`
	ChromePath        string `mapstructure:"CHROME_PATH" validate:"required"`         // Path to the Chrome executable
	ServiceRequestKey string `mapstructure:"SERVICE_REQUEST_KEY" validate:"required"` // Key to protect internal routes
}

var Env *Config

func setDefaults() {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("FRONTEND_BASE_URL", "http://localhost:5173")
}

func Load() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	_, b, _, _ := runtime.Caller(0)
	viper.AddConfigPath(filepath.Dir(b) + "/../..") // Set like this so it can be loaded by test suites as well

	if err := viper.ReadInConfig(); err != nil {
		typ := reflect.TypeOf(Env).Elem()
		for i := range typ.NumField() {
			viper.BindEnv(typ.Field(i).Tag.Get("mapstructure"))
		}
	}

	lo.Try(func() error {
		setDefaults()
		return nil
	})

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatal(err)
	}

	if errs := validator.New().Struct(Env); errs != nil {
		log.Fatal("Invalid environment configuration\n", errs)
	}
}
