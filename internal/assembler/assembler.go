package assembler

import (
	"fmt"
	"log"

	"github.com/gradecak/bouquets/internal/types"
)

type Assembler struct {
	flowers map[int]int // map[species]amount
	designs []*types.BouquetDesign
}

func New(designs []*types.BouquetDesign) (*Assembler, error) {
	flowers := make(map[int]int)
	for species := 97; species < 123; species++ {
		flowers[species] = 0
	}

	return &Assembler{flowers, designs}, nil
}

func (a *Assembler) Run(in chan *types.Flower, out chan *types.Bouquet, end chan bool) {
	for {
		select {
		case <-end:
			log.Print("terminating")
			return
		case flower := <-in:
			// log.Println("processing flower")
			bouquet, found := a.processFlower(flower)
			if found {
				// consume the flowers
				for species, amount := range bouquet.Flowers {
					a.flowers[species] = a.flowers[species] - amount
				}
				fmt.Printf("found bouquet %+v", bouquet)
				out <- bouquet
			}
		}
	}
}

func (a *Assembler) processFlower(flower *types.Flower) (*types.Bouquet, bool) {
	a.flowers[flower.Species]++
	for _, design := range a.designs {
		// use a closure because golang break statements terminate ALL
		// loops not just the one in current scope
		bouquet, ok := func(design *types.BouquetDesign) (*types.Bouquet, bool) {

			toFill := design.FillAmount
			bouquet := design.NewBouquet()
			for species, available := range a.flowers {

				// if we need this species in our design, check if there is
				// enough of it
				if required, ok := design.Design[species]; ok {
					if available < required {
						return nil, false
					}
					bouquet.Flowers[species] = required
				} else if toFill > 0 && available > 0 {
					// if we dont need this species, check if we can use it for
					// filler flowers
					if available >= toFill {
						bouquet.Flowers[species] = toFill
						toFill = 0
					} else {
						toFill = toFill - available
						bouquet.Flowers[species] = available
					}
				}
			}
			return bouquet, toFill == 0
		}(design)

		if ok {
			return bouquet, ok
		}
	}
	return nil, false
}
