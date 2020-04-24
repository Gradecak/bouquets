package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

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
	var (
		outchan           = make(chan *types.Bouquet, 5000)
		inchanL           = make(chan *types.Flower, 5000)
		inchanS           = make(chan *types.Flower, 5000)
		endS              = make(chan bool)
		endL              = make(chan bool)
		largeAssembler, _ = assembler.New(a.largeDesigns)
		smallAssembler, _ = assembler.New(a.smallDesigns)
	)

	go printResults(outchan)
	go largeAssembler.Run(inchanL, outchan, endL)
	go smallAssembler.Run(inchanS, outchan, endS)

	// begin parsing and sorting input stream of flowers
	for flowerStream.Scan() {
		flower, err := a.parser.ParseFlower(flowerStream.Text())
		if err != nil {
			log.Printf("malformed flower found in stream... skipping (reason: %v)", err.Error())
			continue
		}

		// after parsing flower sort it into the appropriate assembler
		if flower.Size == 'L' {
			inchanL <- flower
		} else {
			inchanS <- flower
		}
	}

	// input stream has finished. wait for all processing to finish and
	// then clean up
	for len(outchan) > 0 || len(inchanL) > 0 || len(inchanS) > 0 {
		// sleep the thread to free up available resources
		time.Sleep(time.Second)
	}
	endL <- true
	endS <- true
}

func printResults(in chan *types.Bouquet) {
	for {
		bouquet := <-in
		fmt.Printf("%s\n", bouquet)
	}
}
