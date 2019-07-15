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
	flagExample := flag.Bool("example", false, "Prints an example configuration file that you can use to create your own.")
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

	if *flagExample {
		example, err := exampleConfig()
		if err != nil {
			fmt.Printf("Something went really wrong. Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(example)
		return
	}

	config, err := newConfig(*flagConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	recipes, err := pick(config.Location, config.RecipeCount)
	if err != nil {
		fmt.Printf("Failed to collect recipes. Error: %s\n", err)
		os.Exit(1)
	}

	body := makeBody(config.Prefix, recipes)
	if config.SendEmail {
		err = sendEmail(config, body)
		if err != nil {
			fmt.Printf("Failed to send email. Error: %s\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Email is turned off in config. Outputting body of email here.")
		fmt.Println(body)
	}
}
