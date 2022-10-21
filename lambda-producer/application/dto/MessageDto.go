package dto

type MessageDto struct {
	Path         string         `json:"path"`
	MessageTexts []MessageTexts `json:"message_texts"`
}

type MessageTexts struct {
	TextType     string  `json:"text_type"`
	Confidence   float64 `json:"confidence"`
	DetectedText string  `json:"detected_text"`
}
