package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	collections, run, verbose := cliParser()

	fmt.Println(*collections)

	if *run {
		fmt.Println("m")
		return
	}

	if *verbose {
		fmt.Println("m")
		return
	}
}

func cliParser() (*string, *bool, *bool) {
	collections := flag.String("collections", "", "list collections by NAME")
	run := flag.Bool("run", false, "run configs")
	verbose := flag.Bool("verbose", false, "display more information")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\naba - Just like Insomnia but cli.\n")
		fmt.Fprintln(flag.CommandLine.Output(), "\nUsage information:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !*run && *collections == "" {
		flag.Usage()
		os.Exit(0)
	}

	return collections, run, verbose
}
