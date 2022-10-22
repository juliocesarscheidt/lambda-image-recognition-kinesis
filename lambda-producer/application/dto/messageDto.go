package dto

type MessageDto struct {
	Path         string            `json:"path"`
	MessageTexts []MessageTextsDto `json:"message_texts"`
}
