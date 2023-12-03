package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func toJSON(inputPath, outputPath string) {
	yamlData, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var jsonData interface{}

	err = yaml.Unmarshal(yamlData, &jsonData)
	if err != nil {
		log.Fatalf("Error converting YAML to JSON: %v", err)
	}

	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	writeOutput(jsonBytes, outputPath)
}

func toYAML(inputPath, outputPath string) {
	jsonData, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var yamlData interface{}
	err = json.Unmarshal(jsonData, &yamlData)
	if err != nil {
		log.Fatalf("Error converting JSON to YAML: %v", err)
	}

	yamlBytes, err := yaml.Marshal(yamlData)
	if err != nil {
		log.Fatalf("Error marshalling YAML: %v", err)
	}

	writeOutput(yamlBytes, outputPath)
}

func writeOutput(data []byte, outputPath string) {
	if outputPath != "" {
		err := os.WriteFile(outputPath, data, 0644)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	} else {
		fmt.Println(string(data))
	}
}

func main() {
	toJSONCommand := flag.NewFlagSet("tojson", flag.ExitOnError)
	toJSONFile := toJSONCommand.String("file", "", "Input YAML file")
	toJSONOutput := toJSONCommand.String("output", "", "Output JSON file")

	toYAMLCommand := flag.NewFlagSet("toyaml", flag.ExitOnError)
	toYAMLFile := toYAMLCommand.String("file", "", "Input JSON file")
	toYAMLOutput := toYAMLCommand.String("output", "", "Output YAML file")

	if len(os.Args) < 2 {
		fmt.Println("Please specify a command: tojson or toyaml")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "tojson":
		toJSONCommand.Parse(os.Args[2:])
		if *toJSONFile == "" {
			fmt.Println("Please provide the path to the input YAML file")
			os.Exit(1)
		}
		toJSON(*toJSONFile, *toJSONOutput)

	case "toyaml":
		toYAMLCommand.Parse(os.Args[2:])
		if *toYAMLFile == "" {
			fmt.Println("Please provide the path to the input JSON file")
			os.Exit(1)
		}
		toYAML(*toYAMLFile, *toYAMLOutput)

	default:
		fmt.Println("Invalid command. Please use tojson or toyaml.")
		os.Exit(1)
	}
}
