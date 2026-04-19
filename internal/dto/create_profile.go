package dto

type CreateProfileInput struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required,min=18,max=120"`
	Email string `json:"email" validate:"email"`
	Phone string `json:"phone" validate:"e164"`
}

type CreateProfileOutput struct {
	ID string `json:"id"` // Здесь это поле, чтобы возвращать ID нового профиля
}
