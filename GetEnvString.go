package env

import "os"

func String(name string, required bool, defaultValue string, usage string) string {
	value := os.Getenv(name)
	if value == "" {
		if required {
			panic("missing required environment variable: " + name)
		}
		return defaultValue
	}
	return value
}
