package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPhoneNumber(t *testing.T) {
	testCases := []struct {
		name     string
		phone    string
		expected bool
	}{
		{"Valid Phone", "01012345678", true},
		{"Invalid Phone", "invalid_phone", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidPhoneNumber(tc.phone)
			assert.Equal(t, tc.expected, result, "Unexpected result for IsValidPhoneNumber")
		})
	}
}
