package main

import (
	"os"
	"regexp"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// Read the flag from command line.
	path, err := ParseFlag()
	if err != nil {
		os.Exit(1)
		return
	}

	// Read the configFilePath
	configData, err := ReadFile(path.ConfigFile)
	if err != nil {
		os.Exit(1)
		return
	}
	configDataString := string(configData)

	// Check if there is really a replace syntax
	envReplaceSyntax := GetEnvReplaceSyntaxRegExp()
	searchResult := envReplaceSyntax.FindAllString(configDataString, -1)
	logger.Info("Total Replace: ", len(searchResult))
	logger.Debug(searchResult)
	if len(searchResult) == 0 {
		logger.Warn("There is no replacing syntax `${...}` in the config file. Kindly Exiting... :-)")
		os.Exit(0)
		return
	}

	// Read Env file.
	err = ReadEnvConfigFile(path.EnvFile)
	if err != nil {
		os.Exit(1)
		return
	}

	// Replace the `${NAME}` by using RegExp.
	for _, value := range searchResult {
		// Substring value to remove ${ and }
		lengthOfEnvKey := len(value)
		envKey := value[2 : lengthOfEnvKey-1]
		logger.Debug("Getting value for key: ", envKey)

		//  Get env from viper.
		envValue := viper.GetString(envKey)
		logger.Debug("Value is: ", envValue)
		if envValue == "" {
			logger.Error("There is no replacement for this key: ", envKey)
			os.Exit(1)
			return
		}
		logger.Info("Replacing the ", value, " value with ", envValue)

		// Replace that with value from viper.
		localMatcher, _ := regexp.Compile(value)
		configDataString = localMatcher.ReplaceAllString(configDataString, envValue)
	}
	logger.Info("Done replacing.")

	// Write to file.
	err = os.WriteFile(path.OutputFile, []byte(configDataString), 0777)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	logger.Info("Done the process. Goodbye :-)")
	os.Exit(0)
}
