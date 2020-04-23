package main

import (
	"bufio"
	"log"

	"github.com/gradecak/bouquets/internal/assembler"
	"github.com/gradecak/bouquets/internal/parser"
	"github.com/gradecak/bouquets/internal/types"
)

type App struct {
	parser       *parser.Parser
	largeDesigns []*types.BouquetDesign
	smallDesigns []*types.BouquetDesign
}

func NewApp(designs []string) (*App, error) {
	parser := parser.New()
	largeDesigns := []*types.BouquetDesign{}
	smallDesigns := []*types.BouquetDesign{}

	for _, designStr := range designs {
		design, err := parser.ParseDesign(designStr)
		if err != nil {
			return nil, err
		}

		if design.Size == 'L' {
			largeDesigns = append(largeDesigns, design)
		} else {
			smallDesigns = append(smallDesigns, design)
		}
	}

	return &App{
		parser:       parser,
		largeDesigns: largeDesigns,
		smallDesigns: smallDesigns,
	}, nil
}

func (a *App) Run(flowerStream *bufio.Scanner) {
	outchan := make(chan *types.Bouquet, 500)
	inchan := make(chan *types.Flower, 5000)
	end := make(chan bool)

	go printResults(outchan)
	largeAssembler, _ := assembler.New()
	go largeAssembler.Run(inchan, outchan, end)

	for flowerStream.Scan() {
		flowerStream.Text()
		inchan <- nil
	}

	for len(outchan) > 0 || len(inchan) > 0 {
		// wait for printing to finish and for the processing of all
		// flowers to be complete
	}
}

func printResults(in chan *types.Bouquet) {
	log.Print("got to here")
	for {
		x := <-in
		log.Print("got to here")
		log.Printf("%+v", x)
	}
}
