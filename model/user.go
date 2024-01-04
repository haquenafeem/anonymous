package model

type User struct {
	ID             string `gorm:"primaryKey"`
	Email          string
	Password       string
	Profile_Pic_ID string
}

func (User) TableName() string {
	return "users"
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Err       string `json:"err"`
	IsSuccess bool   `json:"is_success"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Err       string `json:"err"`
	IsSuccess bool   `json:"is_success"`
	Token     string `json:"jwt_token"`
}

type UploadResponse struct {
	Err       string `json:"err"`
	IsSuccess bool   `json:"is_success"`
}

type ProfilePicResponse struct {
	Err          string `json:"err"`
	IsSuccess    bool   `json:"is_success"`
	ProfilePicID string `json:"profile_pic_id"`
}
