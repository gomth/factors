package factors

import (
	"testing"
)

func cmp(left []int, right []int) bool {

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func TestFactorsInvalid(t *testing.T) {
	actual, err := GetUsingTrialDivisionForInt(-1)
	if err == nil {
		t.Errorf("GenerateInvalid failed.\nError expected")
	}

	if actual != nil {
		t.Errorf("GenerateInvalid failed.\nExpected: %v\nActual: %v", []int{}, actual)
	}
}

func TestFactors30(t *testing.T) {
	actual, err := GetUsingTrialDivisionForInt(30)
	if err != nil {
		t.Errorf("TestFactors failed.\nError: %s", err.Error())
	}

	if len(actual) == 0 {
		t.Errorf("TestFactors failed.")
	}

	expected := []int{2, 3, 5}
	if !cmp(actual, expected) {
		t.Errorf("TestFactors30 failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestFactorsOfPrime(t *testing.T) {
	actual, err := GetUsingTrialDivisionForInt(31)
	if err == nil {
		t.Errorf("TestFactorsOfPrime failed.\nError expected")
	}

	if len(actual) == 0 {
		t.Errorf("TestFactorsOfPrime failed.")
	}

	expected := []int{31}
	if !cmp(actual, expected) {
		t.Errorf("TestFactorsOfPrime failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}
