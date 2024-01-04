package model

type Message struct {
	ID     string `gorm:"primaryKey"`
	Data   string
	UserID string
}

func (Message) TableName() string {
	return "messages"
}

type PostMessageRequest struct {
	UserID  string
	Message string
}

type PostMessageResponse struct {
	IsSuccess bool   `json:"is_success"`
	Err       string `json:"err"`
}

type GetAllMessagesResponse struct {
	Messages  []Message
	IsSuccess bool   `json:"is_success"`
	Err       string `json:"err"`
}
