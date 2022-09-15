package main

import (
	"bytes"
	"flag"
	"os"
	"regexp"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Example flag:
// -config-file=$HOME/project/configs/test.yaml -env-file=$HOME/project/envs/prod.env -out=$HOME/project/out/test.yaml

func main() {
	// Read env from file based on inputted flag.
	configFilePath := flag.String("config-file", "", "A Path to your yaml file for a replace.")
	envFilePath := flag.String("env-file", "", "A Path to your .env file containing env var for a replace.")
	outFilePath := flag.String("out", "", "A Path to place where the output of this replacement will be on.")

	flag.Parse()
	logger.Info("ConfigFile Path: ", *configFilePath)
	logger.Info("EnvFile Path: ", *envFilePath)
	logger.Info("OutFile Path: ", *outFilePath)

	// Read the configFilePath
	configData, err := os.ReadFile(*configFilePath)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	configDataString := string(configData)

	// Check if there is really a replace syntax
	r, _ := regexp.Compile(`\${(.*)}`)
	searchResult := r.FindAllString(configDataString, -1)
	logger.Info("RegRex Result: ", searchResult)
	if len(searchResult) == 0 {
		logger.Info("There is no replacing syntax `${...}` in the config file. Kindly Exiting... :-)")
		os.Exit(0)
	}

	//  Read target file.
	envData, err := os.ReadFile(*envFilePath)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	viper.SetConfigType("env")
	err = viper.ReadConfig(bytes.NewBuffer(envData))
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// Replace the `${NAME}` by using RegEx.
	for _, value := range searchResult {
		// Substring value to remove ${ and }
		lengthOfEnvKey := len(value)
		envKey := value[2 : lengthOfEnvKey-1]

		//  Get env from viper.
		envValue := viper.GetString(envKey)
		if envValue == "" {
			logger.Error("There is no replacement for this key: ", envKey)
			os.Exit(1)
		}

		// Replace that with value from viper.
		localMatcher, _ := regexp.Compile(value)
		configDataString = localMatcher.ReplaceAllString(configDataString, envValue)
	}

	logger.Info(configDataString)
	// Write to file.
	err = os.WriteFile(*outFilePath, []byte(configDataString), 0777)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	os.Exit(0)
}
