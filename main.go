package main

import (
	"fmt"
	"os"
	"encoding/base64"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

	fmt.Println("Initializing ENV-2-FS...")

	envB64 := os.Getenv("ENV_FS_B64")
	if envB64 == "" {
		fmt.Println("ENV_FS_B64 is empty!")
		os.Exit(1)
	}

	fmt.Println("[DEBUG] ENV_FS_B64:", envB64)

	envDecoded, err := base64.StdEncoding.DecodeString(envB64)
	handleError(err)

	fmt.Println("[DEBUG] []byte:", envDecoded, "\n")

	var files Files
	err = yaml.Unmarshal(envDecoded, &files)
	handleError(err)

	fmt.Println("Num of Files:", len(files.File))

	for path, content := range files.File {
		fmt.Println("Writing file:", path)
		d := []byte(content)
		err = ioutil.WriteFile(path, d, 0644)
		handleError(err)
	}

}
