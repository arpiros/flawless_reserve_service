package middleware_test

import (
	"reserve_service/middleware"
	"reserve_service/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		validate(
			t,
			models.Member{
				ID:    1,
				Name:  "aaa",
				Email: "aaaa@aaaa.aaa",
			},
			"",
		)
	})
	t.Run("Failure: Required", func(t *testing.T) {
		validate(
			t,
			models.Member{
				ID:   1,
				Name: "aaa",
			},
			"Key: 'Member.Email' Error:Field validation for 'Email' failed on the 'required' tag",
		)
	})
	t.Run("Failure: Email Format", func(t *testing.T) {
		validate(
			t,
			models.Member{
				ID:    1,
				Name:  "aaa",
				Email: "aaaa",
			},
			"Key: 'Member.Email' Error:Field validation for 'Email' failed on the 'email' tag",
		)
	})
}

func validate(t *testing.T, member models.Member, expected string) {
	middleware_validator := middleware.NewValidator()
	err := middleware_validator.Validate(member)
	errorString := ""
	if err != nil {
		errorString = err.Error()
	}
	assert.Equal(t, errorString, expected)
}
