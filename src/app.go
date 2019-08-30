package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"

	"github.com/PhilippHeuer/normalize-ci/pkg/azuredevops"
	"github.com/PhilippHeuer/normalize-ci/pkg/githubactions"
	"github.com/PhilippHeuer/normalize-ci/pkg/gitlabci"
)

// App Properties
var Version string

// Init Hook
func init() {
	// Logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

// CLI Main Entrypoint
func main() {
	// detect debug mode
	debugValue, debugIsSet := os.LookupEnv("DEBUG")
	if debugIsSet && strings.ToLower(debugValue) == "true" {
		log.SetLevel(log.TraceLevel)
	}

	// initialize normalizers
	var normalizers []Normalizer
	normalizers = append(normalizers, azuredevops.NewNormalizer())
	normalizers = append(normalizers, githubactions.NewNormalizer())
	normalizers = append(normalizers, gitlabci.NewNormalizer())

	// get all environment variables
	var env []string // current environment
	for _, entry := range os.Environ() {
		env = append(env, entry)
		log.Debug("ENV: ", entry)
	}

	// normalize (iterate over all supported systems and normalize variables if possible)
	var normalized []string
	for _, normalizer := range normalizers {
		if normalizer.Check(env) == true {
			log.Debug("Matched " + normalizer.GetName() + ", not checking for any other matches.")
			normalized = normalizer.Normalize(env)
			break
		} else {
			log.Debug("Didn't match in " + normalizer.GetName())
		}
	}

	// set normalized variables in current session
	for _, entry := range normalized {
		entrySplit := strings.SplitN(entry, "=", 2)
		log.Debug("Setting var in current session: " + entry)

		err := os.Setenv(entrySplit[0], entrySplit[1])
		common.CheckForError(err)

		// print via stdout
		s := fmt.Sprintf("export %s=%s\n", entrySplit[0], entrySplit[1])
		io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	}
}