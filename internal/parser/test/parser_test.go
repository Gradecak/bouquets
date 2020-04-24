package test

import (
	"testing"

	"github.com/gradecak/bouquets/internal/parser"
	"github.com/gradecak/bouquets/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestParseDesignWithFillerOK(t *testing.T) {
	parser := parser.New()
	expectedDesign := &types.BouquetDesign{
		Name: 'A',
		Size: 'L',
		Design: map[int]int{
			int('a'): 4,
			int('b'): 3,
		},
		FillAmount: 1,
		Total:      8,
	}

	designStr := "AL4a3b8"
	parsedDesign, err := parser.ParseDesign(designStr)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectedDesign, parsedDesign)
}

func TestParseDesignOK(t *testing.T) {
	parser := parser.New()
	expectedDesign := &types.BouquetDesign{
		Name: 'A',
		Size: 'L',
		Design: map[int]int{
			int('a'): 4,
			int('b'): 3,
		},
		FillAmount: 0,
		Total:      7,
	}

	designStr := "AL4a3b7"
	parsedDesign, err := parser.ParseDesign(designStr)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expectedDesign, parsedDesign)
}

func TestParseDesignMalformed(t *testing.T) {
	parser := parser.New()
	designStr := "ABCD"
	_, err := parser.ParseDesign(designStr)
	if err == nil {
		t.Errorf("Error not raised")
	}
}

func TestParseEmptyDesign(t *testing.T) {
	parser := parser.New()
	if _, err := parser.ParseDesign(""); err == nil {
		t.Errorf("Error not raised")
	}
}

func TestParseFlowerLargeOK(t *testing.T) {
	parser := parser.New()
	flowerStr := "aL"
	expected := &types.Flower{Species: int('a'), Size: rune('L')}
	flower, err := parser.ParseFlower(flowerStr)
	if err != nil {
		t.Errorf("Failed to parse valid flower string")
	}
	assert.Equal(t, expected, flower)
}

func TestParseFlowerSmallOK(t *testing.T) {
	parser := parser.New()
	flowerStr := "aS"
	expected := &types.Flower{Species: int('a'), Size: rune('S')}
	flower, err := parser.ParseFlower(flowerStr)
	if err != nil {
		t.Errorf("Failed to parse valid flower string")
	}
	assert.Equal(t, expected, flower)
}

func TestParseFlowerEmpty(t *testing.T) {
	parser := parser.New()
	if _, err := parser.ParseFlower(""); err == nil {
		t.Error("Failed to raise error on invalid parse string")
	}
}
