package types

type BouquetDesign struct {
	Name   rune
	Size   rune
	Design map[rune]int // map[species]amount
	Total  int
}

type Bouquet struct {
	Name    rune
	Size    rune
	Flowers map[rune]int
}
