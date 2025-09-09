package dto

type Captcha struct {
	CaptchaID string `json:"captcha_id"` // Captcha ID
}

type UpdateLoginPassword struct {
	OldPassword string `json:"old_password" binding:"required"` // Old password (md5 hash)
	NewPassword string `json:"new_password" binding:"required"` // New password (md5 hash)
}

type LoginToken struct {
	AccessToken string `json:"access_token"` // Access token (JWT)
	TokenType   string `json:"token_type"`   // Token type (Usage: Authorization=${token_type} ${access_token})
	ExpiresAt   int64  `json:"expires_at"`   // Expired time (Unit: second)
}

type UpdateCurrentUser struct {
	Name   string `json:"name" binding:"required,max=64"` // Name of user
	Phone  string `json:"phone" binding:"max=32"`         // Phone number of user
	Email  string `json:"email" binding:"max=128"`        // Email of user
	Remark string `json:"remark" binding:"max=1024"`      // Remark of user
}
