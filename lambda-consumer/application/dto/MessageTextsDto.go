package dto

type MessageTextsDto struct {
	TextType     string  `json:"text_type"`
	Confidence   float64 `json:"confidence"`
	DetectedText string  `json:"detected_text"`
}
