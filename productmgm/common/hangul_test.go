package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConsonants(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"ㅅㅋㄹ", true},
	}

	for _, tc := range testCases {
		result := IsConsonants(tc.input)
		assert.Equal(t, tc.expected, result, "Unexpected result for IsConsonants")
	}
}

func TestGetWhereClause(t *testing.T) {
	testCases := []struct {
		input    string
		column   string
		expected string
	}{
		{"ㅅㅋㄹ", "name", "SUBSTRING(name, 1, 1) >= '사' AND SUBSTRING(name, 1, 1) < '싸' AND SUBSTRING(name, 2, 1) >= '카' AND SUBSTRING(name, 2, 1) < '타' AND SUBSTRING(name, 3, 1) >= '라' AND SUBSTRING(name, 3, 1) < '마'"},
	}

	for _, tc := range testCases {
		result := GetWhereClause(tc.input, tc.column)
		assert.Equal(t, tc.expected, result, "Unexpected result for GetWhereClause")
	}
}
