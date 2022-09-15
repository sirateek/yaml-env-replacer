package main

import "flag"

// Example flag:
// -config-file=$HOME/project/configs/test.yaml -env-file=$HOME/project/envs/prod.env -out=$HOME/project/out/test.yaml

func main() {
	// Read env from file based on inputted flag.
	configFilePath := flag.String("config-file", "", "A Path to your yaml file for a replace.")
	envFilePath := flag.String("env-file", "", "A Path to your .env file containing env var for a replace.")
	outFilePath := flag.String("out", "", "A Path to place where the output of this replacement will be on.")

	// TODO: Read target file.

	// TODO: Replace the `${NAME}` by using RegEx.

	// TODO: Write to file.
}
