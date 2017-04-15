package main

import (
	"fmt"
	"os"
	"encoding/base64"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	appName	     = "go-envfs"
	envVarName   = "ENVFS_YAML_B64"
	version      = "0.1"
)

type Files struct {
	File map[string]string `yaml:"files"`
}

func handleError(e error) {
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}
}

func main() {

	fmt.Println("Initializing", appName, "version", version, "...")

	envB64 := os.Getenv(envVarName)
	if envB64 == "" {
		fmt.Println(envVarName, "is empty. Nothing to do here.")
		os.Exit(1)
	}


	fmt.Println("Attempting base64 decode of:", envVarName, "...")
	envDecoded, err := base64.StdEncoding.DecodeString(envB64)
	handleError(err)

	//fmt.Println("[DEBUG] []byte:", envDecoded, "\n")

	var files Files
	fmt.Println("Attempting to Unmarshal yaml ...")
	err = yaml.Unmarshal(envDecoded, &files)
	handleError(err)

	fmt.Println("Number of files to write:", len(files.File))

	for path, content := range files.File {
		fmt.Println("Writing file with mode 0644:", path)
		d := []byte(content)
		err = ioutil.WriteFile(path, d, 0644)
		handleError(err)
	}

	fmt.Println(appName, "has completed execution. Exiting.")
}
