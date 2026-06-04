package utils

import "testing"

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"1234567890", true},
		{"123456789", false},
		{"12345678901", false},
		{"abcdefghij", false},
	}

	for _, test := range tests {
		result := IsValidPhone(test.phone)
		if result != test.expected {
			t.Errorf("IsValidPhone(%s) = %v; expected %v", test.phone, result, test.expected)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"invalid.email", false},
		{"another@domain.org", true},
	}

	for _, test := range tests {
		result := IsValidEmail(test.email)
		if result != test.expected {
			t.Errorf("IsValidEmail(%s) = %v; expected %v", test.email, result, test.expected)
		}
	}
}
