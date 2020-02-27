package main

import (
	"flag"
	"os"
)

func main() {

	actionPtr := flag.String("action", "", "Metric {generate|update|validate};. (Required)")
	flag.Parse()

	if *actionPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	// flag.PrintDefaults()
	//
	// outputLocation := ""
	// if outputDir != nil {
	// 	outputLocation = *outputDir
	// }
	//
	// if outputFile != nil && *outputFile != "" {
	// 	outputLocation = path.Join(outputLocation, *outputFile)
	// } else if outputLocation != "" {
	// 	outputLocation = path.Join(outputLocation, def.GetID()+".json")
	// }
	//
	// if outputLocation == "" {
	// 	log.Print(string(data))
	// } else {
	// 	err := ioutil.WriteFile(outputLocation, data, os.ModePerm)
	// 	if err != nil {
	// 		log.Print(err)
	// 	} else {
	// 		log.Print("Written data to ", outputLocation)
	// 	}
	// }
}
