package dto

type UpdateProfileInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Age   *int    `json:"age,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

func (i UpdateProfileInput) IsEmpty() bool {
	return i.Name == nil &&
		i.Age == nil &&
		i.Email == nil &&
		i.Phone == nil
}
