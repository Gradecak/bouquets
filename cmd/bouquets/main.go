package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage " + os.Args[0] + " <input file>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open file for streaming")
	}
	// use a buffered input reader as the specifications mentions a
	// stream of flowers.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// parse all of the designs
	designs := []string{}
	scanner.Scan()
	for designString := scanner.Text(); designString != ""; designString = scanner.Text() {
		scanner.Scan()
		designs = append(designs, designString)
	}

	// setup App
	app, err := NewApp(designs)
	if err != nil {
		log.Fatalf("Could not initialise the application (reason: %v)", err.Error())
	}

	app.Run(scanner)
}
