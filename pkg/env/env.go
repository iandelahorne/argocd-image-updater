package env

import (
	"os"
	"strings"
)

// Package env provides some utility functions to interact with the environment
// of the process.

// GetBoolVal retrieves a boolean value from given environment envVar.
// Returns default value if envVar is not set.
func GetBoolVal(envVar string, defaultValue bool) bool {
	if val := os.Getenv(envVar); val != "" {
		if strings.ToLower(val) == "true" {
			return true
		} else if strings.ToLower(val) == "false" {
			return false
		}
	}
	return defaultValue
}

// GetStringVal retrieves a string value from given environment envVar
// Returns default value if envVar is not set.
func GetStringVal(envVar string, defaultValue string) string {
	if val := os.Getenv(envVar); val != "" {
		return val
	} else {
		return defaultValue
	}
}

// StringsFromEnv parses given value from the environment as a list of strings,
// using seperator as the delimeter, and returns them as a slice. The strings
// in the returned slice will have leading and trailing white space removed.
func StringsFromEnv(env string, defaultValue []string, separator string) []string {
	if str := os.Getenv(env); str != "" {
		ss := strings.Split(str, separator)
		for i, s := range ss {
			ss[i] = strings.TrimSpace(s)
		}
		return ss
	}
	return defaultValue
}
