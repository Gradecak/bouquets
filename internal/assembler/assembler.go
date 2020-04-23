package assembler

import (
	"log"

	"github.com/gradecak/bouquets/internal/types"
)

type Assembler struct {
	flowers map[rune]int // map[species]amount
}

func New() (*Assembler, error) {
	return &Assembler{make(map[rune]int)}, nil
}

func (a *Assembler) Run(in chan *types.Flower, out chan *types.Bouquet, end chan bool) {
	for {
		select {
		case <-end:
			log.Print("terminating")
			return
		case flower := <-in:
			bouquet, found := a.processFlower(flower)

			if found {
				out <- bouquet
			}
		}
	}
}

func (a *Assembler) processFlower(flower *types.Flower) (*types.Bouquet, bool) {
	return &types.Bouquet{Name: 'A', Size: 'L'}, true
}
