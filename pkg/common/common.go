package common

import (
	"net/url"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

// IsEnvironmentSet checks if the provided env variables are setting a value
func IsEnvironmentSet(env []string, key string) bool {
	for _, envvar := range env {
		if strings.HasPrefix(envvar, key+"=") {
			return true
		}
	}

	return false
}

// IsEnvironmentSetTo checks if the provided env variables are setting a value
func IsEnvironmentSetTo(env []string, key string, value string) bool {
	for _, envvar := range env {
		z := strings.SplitN(envvar, "=", 2)
		if strings.ToLower(key) == strings.ToLower(z[0]) && strings.ToLower(value) == strings.ToLower(z[1]) {
			return true
		}
	}

	return false
}

// GetEnvironment gets the value of a environment property
func GetEnvironment(env []string, key string) string {
	for _, envvar := range env {
		z := strings.SplitN(envvar, "=", 2)
		if key == z[0] {
			return z[1]
		}

	}

	return ""
}

// GetSlug turns the provided value into a slug
func GetSlug(value string) string {
	slug := value

	// ToLower
	slug = strings.ToLower(slug)

	// Replace
	slug = strings.Replace(slug, "_", "", -1)
	slug = strings.Replace(slug, "/", "-", -1)

	return slug
	// everything except 0-9 and a-z replaced with -. No leading / trailing -.
}

// AssertThatEnvEquals is a helper function that asserts that a env key has a specific value
func AssertThatEnvEquals(t *testing.T, env []string, key string, value string) {
	if IsEnvironmentSetTo(env, key, value) == false {
		t.Errorf(key + " should be " + value)
	}
}

// CheckForError checks if a error happend and logs it, and ends the process
func CheckForError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetHostFromURL gets the host from a url
func GetHostFromURL(addr string) string {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return u.Host
}

// SetupTestLogger prepares the logger for test execution
func SetupTestLogger() {
	// Logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}