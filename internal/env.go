package internal

import "os"

// GetEnvstring get environment var and return default if none.
func GetEnvstring(name string, def string) string {
	a := os.Getenv(name)
	if a != "" {
		return a
	}
	return def
}

// GetEnvbool get env var and return default bool if none.
func GetEnvbool(name string, def bool) bool {
	defaults := []string{"TRUE", "true", "1", "enabled", "ENABLED"}
	a := os.Getenv(name)
	for _, v := range defaults {
		if v == a {
			return true
		}
	}
	return def
}
