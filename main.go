package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	version = "0.0.1"

	seeded = false
)

func seed() {
	if !seeded {
		rand.Seed(time.Now().UnixNano())
		seeded = true
	}
}

func main() {
	flagConfig := flag.String("c", "./config.json", "Location of the configuration file")
	flagHelp := flag.Bool("h", false, "Shows the help menu")
	flagVersion := flag.Bool("v", false, "Shows the version")

	flag.Parse()
	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	if *flagVersion {
		fmt.Println(version)
		return
	}

	config, err := newConfig(*flagConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	recipes, err := pick(config.Location, config.RecipeCount)
	if err != nil {
		fmt.Println("Failed to collect recipes. Error:", err)
		os.Exit(1)
	}

	body := makeBody(config, recipes)
	err = sendEmail(config, body)
	if err != nil {
		fmt.Printf("Failed to send email. Error: %s", err)
		os.Exit(1)
	}
}
