package msteams

const FactSetBlockType = "FactSet"

type FactSet struct {
	Type  string `json:"type"`
	Facts []Fact `json:"facts"`
}

func (b FactSet) GetType() string {
	return b.Type
}

func NewFactSet(facts ...Fact) FactSet {
	return FactSet{
		Type:  FactSetBlockType,
		Facts: facts,
	}
}
