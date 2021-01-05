package msteams

const ColumnBlockType = "Column"

type ColumnWidth string

var ColumnWidthStretch = ColumnWidth("stretch")

type Column struct {
	Type  string      `json:"type"`
	Width ColumnWidth `json:"width"`
	Items []Block     `json:"items"`
}

func (b Column) GetType() string {
	return b.Type
}

func NewColumn(width ColumnWidth, items ...Block) Column {
	return Column{
		Type:  ColumnBlockType,
		Width: width,
		Items: items,
	}
}
