package dtos

type AttachmentDTO struct {
	OdataType    string `json:"odata.type"`
	Name         string `json:"name"`
	ContentType  string `json:"contentType"`
	ContentBytes string `json:"contentBytes"`
}
