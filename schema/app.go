package schema

import (
	"fmt"
	"os"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	defaultHome     = ".stripe"
	defaultDeadline = "60s"
	defaultPort     = "8082"
)

// GetConfigString : Read environment variables
func GetConfigString() string {
	var format = `STRIPE ENVIRONMENT VARIABLES:
	STRIPE_HOME: %s
	STRIPE_DEADLINE: %s
	STRIPE_PORT: %s`

	home, err := GetAppHome()
	if err != nil {
		home = "Unable to read var"
	}
	return fmt.Sprintf(format, home, GetDeadline(), GetPort())
}

// GetAppHome : Get application directory and creates it if necessary
func GetAppHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	configDir, ok := os.LookupEnv("STRIPE_HOME")
	if !ok {
		configDir = path.Join(home, defaultHome)
	}

	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return "", nil
	}

	return configDir, nil
}

// GetDeadline : Get deadline duration
func GetDeadline() time.Duration {
	value, ok := os.LookupEnv("STRIPE_DEADLINE")
	if !ok {
		value = defaultDeadline
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		duration, _ = time.ParseDuration(defaultDeadline)
	}

	return duration
}

// GetPort : Read HTTP server port environment var
func GetPort() string {
	value, ok := os.LookupEnv("STRIPE_PORT")
	if !ok {
		value = defaultPort
	}
	return value
}
