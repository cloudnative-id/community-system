package settings

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	DatabaseHost     string `envconfig:"DATABASE_HOST" required:"true"`
	DatabaseUser     string `envconfig:"DATABASE_USER" required:"true"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" required:"true"`
	DatabaseName     string `envconfig:"DATABASE_NAME" required:"true"`
	DatabasePort     string `envconfig:"DATABASE_PORT"`
}

func NewSettings() Settings {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		log.Fatalln(err)
	}

	return settings
}
