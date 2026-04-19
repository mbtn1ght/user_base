package dto

import (
	"gitlab.golang-school.ru/potok-2/lessons/lesson-22/internal/domain"
)

type GetProfileOutput struct {
	domain.Profile
}

type GetProfileInput struct {
	ID string
}
