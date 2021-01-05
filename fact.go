package msteams


type Fact struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func NewFact(title, value string) Fact {
	return Fact{
		Title: title,
		Value: value,
	}
}
