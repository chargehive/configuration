package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/chargehive/configuration/utils"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	generateCmdList := generateCmd.Bool("list", false, "list configs to generate")
	generateCmdConfig := generateCmd.String("config", "", "name of config to generate")
	generateCmdPretty := generateCmd.Bool("pretty", false, "pretty print the output (optional)")
	generateCmdVersion := generateCmd.String("version", "v1", "version of config to generate")
	generateCmdOutput := generateCmd.String("output", "", "specify a filename to write the output")

	cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanCmdVersion := cleanCmd.String("version", "v1", "version of config to clean")
	cleanCmdPretty := cleanCmd.Bool("pretty", false, "pretty print the output (optional)")
	cleanCmdJson := cleanCmd.String("json", "", "specify a json string to clean")
	cleanCmdFile := cleanCmd.String("file", "", "specify a config file to clean")
	cleanCmdOutput := cleanCmd.String("output", "", "specify a filename to write the output")

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	validateCmdVersion := validateCmd.String("version", "v1", "version of config to validate")
	validateCmdJson := validateCmd.String("json", "", "specify a json string")
	validateCmdFile := validateCmd.String("file", "", "specify a config file")

	if len(os.Args) < 2 {
		fmt.Println("usage: generate, clean or validate")
		os.Exit(1)
	}

	var currentVersion = "v1"

	switch os.Args[1] {
	case "generate":
		_ = generateCmd.Parse(os.Args[2:])
	case "clean":
		_ = cleanCmd.Parse(os.Args[2:])
	case "validate":
		_ = validateCmd.Parse(os.Args[2:])
	default:
		fmt.Println("usage: generate, clean or validate")
		os.Exit(1)
	}

	// Generate configs
	if generateCmd.Parsed() {
		if *generateCmdVersion != currentVersion {
			fmt.Printf("%v is not a valid config version\n", *generateCmdVersion)
			os.Exit(1)
		}

		if *generateCmdList == true {
			// list templates available
			fmt.Printf("%-30v%v\n-----------------------------------------\n", "Template Name", "Description")
			keys := make([]string, 0)
			for k, _ := range utils.Templates {
				keys = append(keys, string(k))
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("%-30v%v\n", k, utils.Templates[utils.Template(k)])
			}
		} else if _, validChoice := utils.Templates[utils.Template(*generateCmdConfig)]; validChoice {
			// generate template output
			out, err := utils.Generate(utils.Template(*generateCmdConfig), *generateCmdVersion, *generateCmdPretty)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// write template output
			if *generateCmdOutput != "" {
				err := ioutil.WriteFile(*generateCmdOutput, out, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("\nWritten data to '%v'\n", *generateCmdOutput)
				}
			}
			fmt.Println(string(out))
		} else {
			// invalid inputs
			if *generateCmdConfig == "" {
				fmt.Println("you must specify a config to generate")
			} else {
				fmt.Printf("%v is not a valid config\n", *generateCmdConfig)
			}
			generateCmd.PrintDefaults()
			os.Exit(1)
		}
	}

	// Clean Configs
	if cleanCmd.Parsed() {
		if *cleanCmdVersion != currentVersion {
			fmt.Printf("%v is not a valid config version\n", *cleanCmdVersion)
			os.Exit(1)
		}

		// load json
		json, err := getJson(cleanCmdJson, cleanCmdFile)
		if err != nil {
			fmt.Println(err.Error())
			generateCmd.PrintDefaults()
			os.Exit(1)
		}

		// perform clean
		cleaned, result, err := utils.Clean(json, *cleanCmdVersion, *cleanCmdPretty)
		if err != nil {
			fmt.Println("Error:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// confirm clean status
		if cleaned {
			fmt.Println("changes have been made to the structure of this config")
		} else {
			fmt.Println("config is up to date")
		}

		// write output
		if *generateCmdOutput != "" {
			err := ioutil.WriteFile(*generateCmdOutput, result, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nWritten data to '%v'\n", *cleanCmdOutput)
			}
		}
		fmt.Println(string(result))
	}

	// Validate Configs
	if validateCmd.Parsed() {
		if *validateCmdVersion != currentVersion {
			fmt.Printf("%v is not a valid config version\n", *validateCmdVersion)
			os.Exit(1)
		}

		json, err := getJson(validateCmdJson, validateCmdFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if errs := utils.Validate(json, *validateCmdVersion); len(errs) > 0 {
			fmt.Println("errors found:")
			for k, v := range errs {
				fmt.Printf("%v : %v", k, v)
			}
			os.Exit(1)
		}
		fmt.Println("no errors found")
	}
}

func getJson(jsonCmd *string, fileCmd *string) ([]byte, error) {
	if *jsonCmd != "" {
		return []byte(*jsonCmd), nil
	}
	if *fileCmd != "" {
		json, err := ioutil.ReadFile(*fileCmd)
		if err != nil {
			return nil, err
		}
		return json, nil
	}
	return nil, errors.New("must specify either a file or json string")
}
