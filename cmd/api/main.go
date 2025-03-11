package api

import (
	"enkya.org/playground/internal/config"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg := config.NewConfig()
	envconfig.MustProcess("", cfg)
}
