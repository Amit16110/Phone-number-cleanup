package phone

import (
	"testing"
)

func TestNumber(t *testing.T) {
	testCases := []struct {
		phoneNumber      string
		expectedNumber   string
		expectedFormat   string
		expectedAreaCode string
	}{
		{"+1 (613)-995-0253", "6139950253", "(613) 995-0253", "613"},
		{"613-995-0253", "6139950253", "(613) 995-0253", "613"},
		{"1 613 995 0253", "6139950253", "(613) 995-0253", "613"},
		{"613.995.0253", "6139950253", "(613) 995-0253", "613"},
		{"+44 123 456789", "", "", ""},
		{"+81 123456789", "", "", ""},
		{"1234", "", "", ""},
		{"ABC1234567", "", "", ""},
	}

	for _, testCase := range testCases {
		t.Run(testCase.phoneNumber, func(t *testing.T) {
			number := Number(testCase.phoneNumber)

			if number != testCase.expectedNumber {
				t.Errorf("Number() returned %s, expected %s", number, testCase.expectedNumber)
			}

			format := Format(testCase.phoneNumber)

			if format != testCase.expectedFormat {
				t.Errorf("Format() returned %s, expected %s", format, testCase.expectedFormat)
			}

			areaCode := AreaCode(testCase.phoneNumber)
			if areaCode != testCase.expectedAreaCode {
				t.Errorf("AreaCode() returned %s, expected %s", areaCode, testCase.expectedAreaCode)
			}
		})
	}
}

// Benchmark
func BenchmarkNumber(b *testing.B) {
	phoneNumber := "+1 (613)-995-0253"
	for i := 0; i < b.N; i++ {
		Number(phoneNumber)
	}
}

func BenchmarkFormat(b *testing.B) {
	phoneNumber := "+1 (613)-995-0253"
	for i := 0; i < b.N; i++ {
		Format(phoneNumber)
	}
}

func BenchmarkAreaCode(b *testing.B) {
	phoneNumber := "+1 (613)-995-0253"
	for i := 0; i < b.N; i++ {
		AreaCode(phoneNumber)
	}
}
