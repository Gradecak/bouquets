package parser

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gradecak/bouquets/internal/types"
)

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

var designExp = regexp.MustCompile(`^(?P<name>[A-Z])(?P<size>L|S)(?P<flowers>([\d]+[a-z]{1})+)(?P<total>[\d]+)$`)
var flowerExp = regexp.MustCompile(`([\d]+[a-z]{1})`)

func (_ Parser) ParseDesign(line string) (*types.BouquetDesign, error) {
	design := &types.BouquetDesign{}
	match := designExp.FindStringSubmatch(line)

	// go over all of the captured named groups to extract parts of
	// design
	for i, group := range designExp.SubexpNames() {
		if i != 0 && group != "" {
			switch group {
			case "name":
				// since we are only matching a single character in regex
				// we're relatively safe to just take index 0
				design.Name = rune(match[i][0])
			case "size":
				design.Size = rune(match[i][0])
			case "total":
				total, err := strconv.Atoi(match[i])
				if err != nil {
					return nil, errors.New("Could not parse total number of flowers")
				}
				design.Total = total
			case "flowers":
				design.Design = make(map[rune]int)
				// match all of the individual flowers in flowers capture
				// group
				flowers := flowerExp.FindAllString(match[i], -1)
				for _, flower := range flowers {
					number, err := strconv.Atoi(flower[:len(flower)-1])
					if err != nil {
						return nil, errors.New("could not extract flower composition of design")
					}
					// last character of each match is the species
					species := rune(flower[len(flower)-1])
					design.Design[species] = number
				}
			}
		}
	}
	return design, nil
}

func (_ Parser) ParseFlower(line string) (*types.Flower, error) {
	return nil, errors.New("not implemented")
}
