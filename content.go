package msteams

const (
	ContentSchemaAdaptiveCard = "http://adaptivecards.io/schemas/adaptive-card.json"
	ContentTypeCardAdaptive   = "AdaptiveCard"
	ContentVersion12          = "1.2"
)

type Content struct {
	Schema  string  `json:"$schema"`
	Type    string  `json:"type"`
	Version string  `json:"version"`
	Body    []Block `json:"body"`
}

func NewContentCardAdaptive(blocks []Block) Content {
	return Content{
		Schema:  ContentSchemaAdaptiveCard,
		Type:    ContentTypeCardAdaptive,
		Version: ContentVersion12,
		Body:    blocks,
	}
}
