package Fuzzy

const START int16 = 1
const END int16 = 2

type FuzzyMembershipFunction struct {
	Label  string
	Weight float32
}

type StartOrEndDomain struct {
	FirstPoint  float32
	SecondPoint float32
	Label       string
	Weight      float32
}

type FuzzyLogic struct {
	Membership         []FuzzyMembership
	MembershipFunction FuzzyMembershipFunction
	MinDomain          StartOrEndDomain
	MaxDomain          StartOrEndDomain
}

func (soed *StartOrEndDomain) CalculateWeight(input float32, DomainType int16) {
	if DomainType == START {
		if input <= soed.FirstPoint {
			soed.Weight = 1
		} else if input > soed.FirstPoint && input <= soed.SecondPoint {
			soed.Weight = (soed.SecondPoint - input) / (soed.SecondPoint - soed.FirstPoint)
		} else {
			soed.Weight = 0
		}
	} else if DomainType == END {
		if input <= soed.FirstPoint {
			soed.Weight = 0
		} else if input > soed.FirstPoint && input <= soed.SecondPoint {
			soed.Weight = (input - soed.FirstPoint) / (soed.SecondPoint - soed.FirstPoint)
		} else {
			soed.Weight = 1
		}
	} else {
		panic("Error")
	}
}

func (fl *FuzzyLogic) SetMinDomain(label string, firstPoint float32, secondPoint float32) {
	fl.MinDomain.Label = label
	fl.MinDomain.FirstPoint = firstPoint
	fl.MinDomain.SecondPoint = secondPoint
}

func (fl *FuzzyLogic) SetMaxDomain(label string, firstPoint float32, secondPoint float32) {
	fl.MaxDomain.Label = label
	fl.MaxDomain.FirstPoint = firstPoint
	fl.MaxDomain.SecondPoint = secondPoint
}

func (fl *FuzzyLogic) AddTriangleFunction(firstPoint float32, secondPoint float32, thirdPoint float32, label string) {
	Triangle := NewTriangleFunction(firstPoint, secondPoint, thirdPoint, label)
	fl.Membership = append(fl.Membership, *Triangle)
}

func (fl *FuzzyLogic) CalculateWeight(input float32) {
	// Mencari Bobot dari Domain Terendah
	fl.MinDomain.CalculateWeight(input, START)

	// Mencari Bobot dar Domain Tertinggi
	fl.MaxDomain.CalculateWeight(input, END)

	// Mendapatkan Bobot dari Fungsi Segitiga
	for i, _ := range fl.Membership {
		fl.Membership[i].CalculateWeight(input)
	}
}

func (fl *FuzzyLogic) GetResult() []FuzzyMembershipFunction {
	var Membership []FuzzyMembershipFunction

	Membership = append(Membership, FuzzyMembershipFunction{Label: fl.MinDomain.Label, Weight: fl.MinDomain.Weight})

	for _, Member := range fl.Membership {
		Membership = append(Membership, FuzzyMembershipFunction{Label: Member.Label, Weight: Member.Weight})
	}

	Membership = append(Membership, FuzzyMembershipFunction{Label: fl.MaxDomain.Label, Weight: fl.MaxDomain.Weight})

	return Membership
}

func (fl *FuzzyLogic) GetMaxWeight() FuzzyMembershipFunction {
	var Membership []FuzzyMembershipFunction
	var MaxWeightMember FuzzyMembershipFunction
	var MaxWeight float32 = 0

	Membership = append(Membership, FuzzyMembershipFunction{Label: fl.MinDomain.Label, Weight: fl.MinDomain.Weight})

	for _, Member := range fl.Membership {
		Membership = append(Membership, FuzzyMembershipFunction{Label: Member.Label, Weight: Member.Weight})
	}

	for _, Member := range Membership {
		if Member.Weight >= MaxWeight {
			MaxWeight = Member.Weight
			MaxWeightMember = Member
		}
	}

	return MaxWeightMember
}

func NewFuzzyLogic() *FuzzyLogic {
	return &FuzzyLogic{}
}
