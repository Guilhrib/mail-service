package dtos

import "roboteasy.com/main-service/internal/models/enums"

type BodyDTO struct {
	ContentType enums.BodyFormatEnum `json:"contentType"`
	Content     string               `json:"content"`
}
