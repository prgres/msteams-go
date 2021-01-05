package msteams


const TextBlockBlockType = "TextBlock"

type TextBlock struct {
	Type   string              `json:"type"`
	Size   TextBlockTextSize   `json:"size"`
	Weight TextBlockTextWeight `json:"weight"`
	Wrap   bool                `json:"wrap"`
	Text   string              `json:"text"`
}

func (b TextBlock) GetType() string {
	return b.Type
}

func (b TextBlock) WrapDisable() TextBlock {
	b.Wrap = false
	return b
}

func (b TextBlock) WrapEnable() TextBlock {
	b.Wrap = true
	return b
}

type TextBlockTextSize string

var (
	TextBlockTextSizeSmall      = TextBlockTextSize("small")
	TextBlockTextSizeMedium     = TextBlockTextSize("medium")
	TextBlockTextSizeLarge      = TextBlockTextSize("large")
	TextBlockTextSizeExtraLarge = TextBlockTextSize("extraLarge")
)

type TextBlockTextWeight string

var TextBlockTextWeightBolder = TextBlockTextWeight("bolder")

func NewTextBlock(size TextBlockTextSize, weight TextBlockTextWeight, text string) TextBlock {
	return TextBlock{
		Type:   TextBlockBlockType,
		Size:   size,
		Weight: weight,
		Wrap:   true,
		Text:   text,
	}
}
