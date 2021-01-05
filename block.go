package msteams

type Block interface {
	GetType() string
}

type Padding string

var PaddingDefault Padding = "default"

type Spacing string

var (
	SpacingDefault Spacing = "default"
	SpacingSmall   Spacing = "small"
)
