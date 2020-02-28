package main

import (
	"flag"
	"fmt"
	"github.com/chargehive/configuration/utils"
	"os"
	"sort"
)

func main() {
	// sub commands
	generateCommand := flag.NewFlagSet("generate", flag.ExitOnError)
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	validateCommand := flag.NewFlagSet("validate", flag.ExitOnError)

	generateCommandList := generateCommand.Bool("list", false, "List configs to generate")
	generateCommandConfig := generateCommand.String("config", "", "Name of config to generate")
	generateCommandPretty := generateCommand.Bool("pretty", false, "Pretty print the output")

	// os.Arg[0] is the main command
	// os.Arg[1] will be the sub commands
	if len(os.Args) < 2 {
		fmt.Println("generate, update or validate required")
		os.Exit(1)
	}

	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "generate":
		_ = generateCommand.Parse(os.Args[2:])
	case "update":
		_ = updateCommand.Parse(os.Args[2:])
	case "validate":
		_ = validateCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if generateCommand.Parsed() {
		if *generateCommandList == true {
			fmt.Printf("%-30v%v\n-----------------------------------------\n", "Template Name", "Description")
			keys := make([]string, 0)
			for k, _ := range utils.Templates {
				keys = append(keys, string(k))
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("%-30v%v\n", k, utils.Templates[utils.Template(k)])
			}
		} else if _, validChoice := utils.Templates[utils.Template(*generateCommandConfig)]; validChoice {
			out, err := utils.Generate(utils.Template(*generateCommandConfig), *generateCommandPretty)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println(string(out))
			}
		} else {
			fmt.Printf("%v is not a valid config\n", *generateCommandConfig)
			generateCommand.PrintDefaults()
			os.Exit(1)
		}
	}
	if updateCommand.Parsed() {

	}
	if validateCommand.Parsed() {

	}

}
