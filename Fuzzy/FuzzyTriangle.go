package Fuzzy

type FuzzyMembership struct {
	firstPoint  float32
	secondPoint float32
	thirdPoint  float32
	Label       string
	Weight      float32
}

func (fr *FuzzyMembership) CalculateWeight(value float32) {
	if value < fr.firstPoint {
		fr.Weight = 0
	} else if (value >= fr.firstPoint) && (value < fr.secondPoint) {
		fr.Weight = (value - fr.firstPoint) / (fr.secondPoint - fr.firstPoint)
	} else if (value >= fr.secondPoint) && (value < fr.thirdPoint) {
		fr.Weight = (fr.thirdPoint - value) / (fr.thirdPoint - fr.secondPoint)
	} else {
		fr.Weight = 0
	}
}

func NewTriangleFunction(firstPoint float32, secondPoint float32, thirdPoint float32, label string) *FuzzyMembership {
	return &FuzzyMembership{firstPoint: firstPoint, secondPoint: secondPoint, thirdPoint: thirdPoint, Label: label}
}
