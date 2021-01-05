package msteams

const TypeMessage = "message"

type Message struct {
	Type        string       `json:"type"`
	Attachments []Attachment `json:"attachments"`
}

func NewMessage(blocks []Block) Message {
	return Message{
		Type: TypeMessage,
		Attachments: []Attachment{
			NewAttachmentCardAdaptive(
				NewContentCardAdaptive(blocks)),
		},
	}
}
