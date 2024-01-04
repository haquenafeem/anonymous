package model

type Message struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Data   string `json:"data"`
	UserID string `json:"user_id"`
}

func (Message) TableName() string {
	return "messages"
}

type PostMessageRequest struct {
	UserID  string `json:"user_id"`
	Message string
}

type PostMessageResponse struct {
	IsSuccess bool   `json:"is_success"`
	Err       string `json:"err"`
}

type GetAllMessagesResponse struct {
	Messages  []Message `json:"messages"`
	IsSuccess bool      `json:"is_success"`
	Err       string    `json:"err"`
}
