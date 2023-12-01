package config

type Config struct {
	Port         string
	DatabaseName string `mapstructure:"database" validate:"required"     default:"challenge-verifymy"`
	URI          string `mapstructure:"uri"      validate:"required,uri" default:""`
}
