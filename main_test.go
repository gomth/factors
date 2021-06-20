package factors

import "testing"

func TestGetFactors30(t *testing.T) {
	actual, err := Get(30)
	if err != nil {
		t.Errorf("TestGetFactors30 failed.\nError: %s", err.Error())
	}

	if len(actual) == 0 {
		t.Errorf("TestGetFactors30 failed.")
	}

	expected := []int{2, 3, 5}
	if !cmp(actual, expected) {
		t.Errorf("TestGetFactors30 failed.\nExpected: %v\nActual: %v", expected, actual)
	}
}
