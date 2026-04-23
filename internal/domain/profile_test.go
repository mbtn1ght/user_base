package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"user_base/internal/domain"
)

func TestNewProfile_Valid(t *testing.T) {
	p, err := domain.NewProfile("John Doe", 30, "john@example.com", "+15551234567")
	require.NoError(t, err)

	require.NotEqual(t, uuid.Nil, p.ID)
	require.Equal(t, "John Doe", string(p.Name))
	require.Equal(t, 30, int(p.Age))
	require.Equal(t, "john@example.com", p.Contacts.Email)
	require.Equal(t, "+15551234567", p.Contacts.Phone)
	require.Equal(t, domain.Pending, p.Status)
	require.False(t, p.Verified)
	require.False(t, p.IsDeleted())
}

func TestNewProfile_ValidationErrors(t *testing.T) {
	cases := []struct {
		name  string
		age   int
		email string
		phone string
	}{
		{"", 25, "john@example.com", "+15551234567"},      // required name
		{"Al", 25, "john@example.com", "+15551234567"},    // min len 3
		{"John", 17, "john@example.com", "+15551234567"},  // age < 18
		{"John", 121, "john@example.com", "+15551234567"}, // age > 120
		{"John", 25, "invalid-email", "+15551234567"},     // invalid email
		{"John", 25, "john@example.com", "not-e164"},      // invalid phone
	}

	for _, c := range cases {
		_, err := domain.NewProfile(c.name, c.age, c.email, c.phone)
		require.Error(t, err)
	}
}

func TestProfile_Validate_Direct(t *testing.T) {
	p := domain.Profile{
		Name:     domain.Name("Alice"),
		Age:      domain.Age(22),
		Status:   domain.Pending,
		Verified: false,
		Contacts: domain.Contacts{
			Email: "alice@example.com",
			Phone: "+447911123456",
		},
	}

	require.NoError(t, p.Validate())

	p.Age = 0
	require.Error(t, p.Validate())

	p.Age = 22
	p.Name = ""
	require.Error(t, p.Validate())
}

func TestProfile_IsDeleted(t *testing.T) {
	var p domain.Profile
	require.False(t, p.IsDeleted())

	p.DeletedAt = time.Now()
	require.True(t, p.IsDeleted())
}
