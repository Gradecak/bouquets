package test

import (
	"testing"
	"time"

	"github.com/gradecak/bouquets/internal/assembler"
	"github.com/gradecak/bouquets/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestAssemblerSingleDesign(t *testing.T) {
	var (
		designs = []*types.BouquetDesign{
			&types.BouquetDesign{
				Name: 'A',
				Size: 'L',
				Design: map[int]int{
					int('a'): 1,
				},
				FillAmount: 0,
				Total:      1,
			},
		}
		expected    = &types.Bouquet{Name: 'A', Size: 'L', Flowers: map[int]int{int('a'): 1}}
		assemble, _ = assembler.New(designs)
		in          = make(chan *types.Flower, 5)
		out         = make(chan *types.Bouquet)
		end         = make(chan bool)
	)
	go assemble.Run(in, out, end)
	in <- &types.Flower{Species: int('a')}
	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			t.Error("Assembler did not complete within timeout")
			break
		}
		if len(in) == 0 {
			result := <-out
			end <- true
			assert.Equal(t, result, expected)
			break
		}
	}
}

func TestAssemblerMultiDesign(t *testing.T) {
	var (
		designs = []*types.BouquetDesign{
			&types.BouquetDesign{
				Name: 'A',
				Size: 'L',
				Design: map[int]int{
					int('a'): 1,
				},
				FillAmount: 0,
				Total:      1,
			},
			&types.BouquetDesign{
				Name: 'B',
				Size: 'L',
				Design: map[int]int{
					int('b'): 2,
				},
				FillAmount: 0,
				Total:      2,
			},
		}
		assemble, _ = assembler.New(designs)
		in          = make(chan *types.Flower, 5)
		out         = make(chan *types.Bouquet, 2)
		end         = make(chan bool)
	)

	go assemble.Run(in, out, end)

	flowerStream := []*types.Flower{
		&types.Flower{Species: int('a')},
		&types.Flower{Species: int('b')},
		&types.Flower{Species: int('b')},
	}
	for _, flower := range flowerStream {
		in <- flower
	}

	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			t.Error("Assembler did not complete within timeout")
			break
		}
		// VERY BAD HACK wait for len(out) == 2 here because the assembler
		// does not have the capacity to signal to the supervisor that it
		// is in the middle of processing an incoming flower
		if len(in) == 0 && len(out) == 2 {
			assert.Equal(t, 2, len(out))
			break
		}
	}
}
