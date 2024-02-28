package Fuzzy

type FuzzyMembership struct {
	firstEdge  float32
	secondEdge float32
	thirdEdge  float32
	Label      string
	Weight     float32
}

func (fr *FuzzyMembership) CalculateWeight(value float32) {
	if value < fr.firstEdge {
		fr.Weight = 0
	} else if (value >= fr.firstEdge) && (value < fr.secondEdge) {
		fr.Weight = (value - fr.firstEdge) / (fr.secondEdge - fr.firstEdge)
	} else if (value >= fr.secondEdge) && (value < fr.thirdEdge) {
		fr.Weight = (fr.thirdEdge - value) / (fr.thirdEdge - fr.secondEdge)
	} else {
		fr.Weight = 0
	}
}

func NewTriangleFunction(firstEdge float32, secondEdge float32, thirdEdge float32, label string) *FuzzyMembership {
	return &FuzzyMembership{firstEdge: firstEdge, secondEdge: secondEdge, thirdEdge: thirdEdge, Label: label}
}
