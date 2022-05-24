package dto

type MessageDto struct {
	ID       uint   `json:"id"`
	Sender   uint   `json:"sender"`
	Receiver uint   `json:"receiver"`
	Text     string `json:"text"`
}
