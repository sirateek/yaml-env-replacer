package main

import (
	"bytes"
	"flag"
	"os"
	"regexp"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type filePath struct {
	ConfigFile string
	EnvFile    string
	OutputFile string
}

func ParseFlag() filePath {
	// Read env from file based on inputted flag.
	configFilePath := flag.String("config-file", "", "A Path to your yaml file for a replace.")
	envFilePath := flag.String("env-file", "", "A Path to your .env file containing env var for a replace.")
	outFilePath := flag.String("out", "", "A Path to place where the output of this replacement will be on.")

	flag.Parse()
	logger.Info("ConfigFile Path: ", *configFilePath)
	logger.Info("EnvFile Path: ", *envFilePath)
	logger.Info("OutFile Path: ", *outFilePath)

	return filePath{
		ConfigFile: *configFilePath,
		EnvFile:    *envFilePath,
		OutputFile: *outFilePath,
	}

}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err)
	}
	return data, err
}

func GetEnvReplaceSyntaxRegExp() *regexp.Regexp {
	r, _ := regexp.Compile(`\${(.*)}`)
	return r
}

func ReadEnvConfigFile(envFilePath string) error {
	envData, err := ReadFile(envFilePath)
	if err != nil {
		return err
	}

	viper.SetConfigType("env")
	err = viper.ReadConfig(bytes.NewBuffer(envData))
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
