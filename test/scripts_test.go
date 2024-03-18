package test

import (
	"testing"
	"time"

	"github.com/dkshi/vktest/scripts"
)

func TestStringToDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "Valid date",
			input:    "2024-03-18",
			expected: time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "Invalid date format",
			input:    "18-03-2024",
			expected: time.Time{},
			wantErr:  true,
		},
		{
			name:     "Invalid date value",
			input:    "2024-02-30",
			expected: time.Time{},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := scripts.StringToDate(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("StringToDate() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !result.Equal(test.expected) {
				t.Errorf("StringToDate() got = %v, want %v", result, test.expected)
			}
		})
	}
}

func TestСonvertFloat64ToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid conversion",
			input:    123.45,
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Conversion of large float64 value",
			input:    1e20,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := scripts.СonvertFloat64ToInt(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("СonvertFloat64ToInt() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if result != test.expected {
				t.Errorf("СonvertFloat64ToInt() got = %v, want %v", result, test.expected)
			}
		})
	}
}
