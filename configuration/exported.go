package configuration

import (
	// stdlib
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	// other
	"gopkg.in/yaml.v2"
)

var (
	Cfg *config
)

// Initialize initializes package and parses configuration into struct.
func Initialize() {
	log.Println("Initializing configuration...")

	pathRaw, found := os.LookupEnv("GONEWS_CONFIG")
	if !found {
		log.Fatalln("Failed to read configuration - no GONEWS_CONFIG environment variable defined.")
	}

	// Normalize path.
	if strings.HasPrefix(pathRaw, "~") {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln("Failed to obtain user's home directory path: " + err.Error())
		}

		pathRaw = strings.Replace(pathRaw, "~", userHomeDir, 1)
	}
	absPath, err1 := filepath.Abs(pathRaw)
	if err1 != nil {
		log.Fatalln("Failed to get absolute path for configuration file: " + err1.Error())
	}

	// Read and parse configuration file.
	fileData, err2 := ioutil.ReadFile(absPath)
	if err2 != nil {
		log.Fatalln("Failed to read configuration file data: " + err2.Error())
	}

	Cfg = &config{}
	err3 := yaml.Unmarshal(fileData, Cfg)
	if err3 != nil {
		log.Fatalln("Failed to parse configuration file: " + err3.Error())
	}

	log.Printf("Configuration file parsed: %+v\n", Cfg)
}
