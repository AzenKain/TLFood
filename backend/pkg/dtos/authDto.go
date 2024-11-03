package dtos

type SignUpDto struct {
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"PhoneNumber" validate:"required"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateOtpDto struct {
	Email   string `json:"email" validate:"required"`
	OtpType string `json:"otp_type" validate:"required"`
}

type VerifyOtpDto struct {
	OtpId   string `json:"otp_id" validate:"required"`
	Email   string `json:"email" validate:"required"`
	OtpType string `json:"otp_type" validate:"required"`
	OtpCode string `json:"otp_code" validate:"required"`
}
