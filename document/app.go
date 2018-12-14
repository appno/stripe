package document

import (
	"fmt"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	defaultHome     = "$HOME/.stripe"
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

// GetAppHome : Read application directory var
func GetAppHome() (string, error) {
	home, ok := os.LookupEnv("STRIPE_HOME")
	if !ok {
		home = defaultHome
	}

	appHome, err := homedir.Expand(os.ExpandEnv(home))
	if err != nil {
		return "", err
	}

	return appHome, nil
}

// GetDeadline : Read deadline environment var
func GetDeadline() string {
	value, ok := os.LookupEnv("STRIPE_DEADLINE")
	if !ok {
		value = defaultDeadline
	}
	return value
}

// GetDeadlineDuration : Get deadline duration
func GetDeadlineDuration() time.Duration {
	deadline := GetDeadline()

	duration, err := time.ParseDuration(deadline)
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
