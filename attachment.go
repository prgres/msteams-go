package msteams

const AttachmentContentTypeCardAdaptive = "application/vnd.microsoft.card.adaptive"

type Attachment struct {
	ContentType string  `json:"contentType"`
	Content     Content `json:"content"`
}

func NewAttachmentCardAdaptive(content Content) Attachment {
	return Attachment{
		ContentType: AttachmentContentTypeCardAdaptive,
		Content:     content,
	}
}
