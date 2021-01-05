package msteams

const ContainerBlockType = "Container"

type Container struct {
	Type      string  `json:"type"`
	ID        string  `json:"id,omitempty"`
	Separator bool    `json:"separator"`
	Padding   Padding `json:"padding,omitempty"`
	Spacing   Spacing `json:"spacing,omitempty"`
	Items     []Block `json:"items"`
}

func (b Container) GetType() string {
	return b.Type
}

func (b Container) SetID(id string) Container {
	b.ID = id
	return b
}

func (b Container) SetPadding(padding Padding) Container {
	b.Padding = padding
	return b
}

func (b Container) SeparatorEnable() Container {
	b.Separator = true
	return b
}

func (b Container) SeparatorDisable() Container {
	b.Separator = false
	return b
}

func (b Container) SetSpacing(spacing Spacing) Container {
	b.Spacing = spacing
	return b
}
func (b Container) SetItems(items ...Block) Container {
	b.Items = items
	return b
}

func NewContainer(items ...Block) Container {
	return Container{
		Type:      ContainerBlockType,
		Padding:   PaddingDefault,
		Spacing:   SpacingDefault,
		Separator: false,
		Items:     items,
	}
}
