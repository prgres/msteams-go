package msteams

const ColumnSetBlockType = "ColumnSet"

type ColumnSet struct {
	Type    string   `json:"type"`
	Columns []Column `json:"columns"`
}

func (b ColumnSet) GetType() string {
	return b.Type
}

func NewColumnSet(columns ...Column) ColumnSet {
	return ColumnSet{
		Type:    ColumnSetBlockType,
		Columns: columns,
	}
}
