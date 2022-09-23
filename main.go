package main

import (
	"errors"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"log"
	"net/url"
	"os"
)

func main() {
	openapiFilePath, exists := os.LookupEnv("INPUT_SPECFILE")
	if !exists {
		if _, err := os.Stat("openapi.yaml"); errors.Is(err, os.ErrNotExist) {
			if _, err := os.Stat("openapi.yml"); errors.Is(err, os.ErrNotExist) {
				if _, err := os.Stat("openapi.json"); errors.Is(err, os.ErrNotExist) {
					log.Fatal("no openapi file found.")
				} else {
					openapiFilePath = "openapi.json"
				}
			} else {
				openapiFilePath = "openapi.yml"
			}
		} else {
			openapiFilePath = "openapi.yaml"
		}
	}

	var version string
	if openapiFilePath[0:4] == "http" {
		parsed, err := url.Parse(openapiFilePath)
		if err != nil {
			log.Fatal(err)
		}

		doc, err := openapi3.NewLoader().LoadFromURI(parsed)
		if err != nil {
			log.Fatal(err)
		}
		version = doc.Info.Version
	} else {
		doc, err := openapi3.NewLoader().LoadFromFile(openapiFilePath)
		if err != nil {
			log.Fatal(err)
		}
		version = doc.Info.Version
	}

	fmt.Println(fmt.Sprintf(`::set-output name=apiVersion::%s`, version))
}
