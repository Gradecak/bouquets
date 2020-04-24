package types

import (
	"bytes"
	"fmt"
	"sort"
)

type BouquetDesign struct {
	Name rune
	Size rune
	// we keep a list of species as well as the design map due to the
	// fact that golang doesnt provide a way to order maps. In order to
	// exploit the fact that design species appear in alphabetical order
	// we have to double store the species required
	Design map[int]int // map[species]amount
	// as an optimisation when parsing record the amount of "filler"
	// flowers that are required to complete the design
	FillAmount int
	Total      int
}

func (b BouquetDesign) NewBouquet() *Bouquet {
	return &Bouquet{
		Name:    b.Name,
		Size:    b.Size,
		Flowers: make(map[int]int),
	}
}

type Bouquet struct {
	Name    rune
	Size    rune
	Flowers map[int]int // map[species]amount
}

func (b Bouquet) String() string {
	var (
		buf     = new(bytes.Buffer)
		flowers = []int{}
	)
	// because golang doesnt provide ordering on maps, we have to first
	// extract the key and values sort and then print
	for flower, _ := range b.Flowers {
		flowers = append(flowers, flower)
	}
	sort.Ints(flowers)
	for _, flower := range flowers {
		fmt.Fprintf(buf, "%d%s", b.Flowers[flower], string(flower))
	}
	return fmt.Sprintf("%s%s%s", string(b.Name), string(b.Size), buf.String())
}
