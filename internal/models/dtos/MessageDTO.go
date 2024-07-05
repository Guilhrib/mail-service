package dtos

type MessageDTO struct {
	Subject      string          `json:"subject"`
	Body         BodyDTO         `json:"body"`
	ToRecipients []RecipientDTO  `json:"toRecipients"`
	Attachments  []AttachmentDTO `json:"attachments"`
}
