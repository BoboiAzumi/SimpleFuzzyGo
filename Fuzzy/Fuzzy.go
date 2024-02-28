package Fuzzy

const START int16 = 1
const END int16 = 2

type FuzzyMembershipFunction struct {
	Label  string
	Weight float32
}

type StartOrEndDomain struct {
	FirstEdge  float32
	SecondEdge float32
	Label      string
	Weight     float32
}

type FuzzyLogic struct {
	Membership         []FuzzyMembership
	MembershipFunction FuzzyMembershipFunction
	MinDomain          StartOrEndDomain
	MaxDomain          StartOrEndDomain
}

func (soed *StartOrEndDomain) CalculateWeight(input float32, DomainType int16) {
	if DomainType == START {
		if input <= soed.FirstEdge {
			soed.Weight = 1
		} else if input > soed.FirstEdge && input <= soed.SecondEdge {
			soed.Weight = (soed.SecondEdge - input) / (soed.SecondEdge - soed.FirstEdge)
		} else {
			soed.Weight = 0
		}
	} else if DomainType == END {
		if input <= soed.FirstEdge {
			soed.Weight = 0
		} else if input > soed.FirstEdge && input <= soed.SecondEdge {
			soed.Weight = (input - soed.FirstEdge) / (soed.SecondEdge - soed.FirstEdge)
		} else {
			soed.Weight = 1
		}
	} else {
		panic("Error")
	}
}

func (fl *FuzzyLogic) SetMinDomain(label string, firstEdge float32, secondEdge float32) {
	fl.MinDomain.Label = label
	fl.MinDomain.FirstEdge = firstEdge
	fl.MinDomain.SecondEdge = secondEdge
}

func (fl *FuzzyLogic) SetMaxDomain(label string, firstEdge float32, secondEdge float32) {
	fl.MaxDomain.Label = label
	fl.MaxDomain.FirstEdge = firstEdge
	fl.MaxDomain.SecondEdge = secondEdge
}

func (fl *FuzzyLogic) AddTriangleFunction(firstEdge float32, secondEdge float32, thirdEdge float32, label string) {
	Triangle := NewTriangleFunction(firstEdge, secondEdge, thirdEdge, label)
	fl.Membership = append(fl.Membership, *Triangle)
}

func (fl *FuzzyLogic) CalculateWeight(input float32) {
	// Get Min Domain Weight
	fl.MinDomain.CalculateWeight(input, START)

	// Get Max Domain Weight
	fl.MaxDomain.CalculateWeight(input, END)

	// Get Triangle Domain Weight
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
