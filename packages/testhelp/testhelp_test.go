package testhelp

import (
	"testing"
)

func TestAreInterfaceSlicesEqual(t *testing.T) {

	result := AreInterfaceSlicesEqual([]interface{}{1, 2, 3, 4}, []interface{}{1, 2, 3, 4})
	expected := true
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreInterfaceSlicesEqual([]interface{}{"1", "2", "3", "4"}, []interface{}{"1", "2", "3", "4"})
	expected = true
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreInterfaceSlicesEqual([]interface{}{"1", "3", "4"}, []interface{}{"1", "2", "3", "4"})
	expected = false
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreInterfaceSlicesEqual([]interface{}{}, []interface{}{})
	expected = true
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreInterfaceSlicesEqual([]interface{}{"A", 2, true, 3.14}, []interface{}{"A", 2, true, 3.14})
	expected = true
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreInterfaceSlicesEqual([]interface{}{"A", 2, 3.14, "true"}, []interface{}{"A", 2, 3.14, 3.14})
	expected = false
	if result != expected {
		t.Errorf("AreInterfaceSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

}

func TestAreIntSlicesEqual(t *testing.T) {

	result := AreIntSlicesEqual([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	expected := true
	if result != expected {
		t.Errorf("AreIntSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreIntSlicesEqual([]int{}, []int{})
	expected = true
	if result != expected {
		t.Errorf("AreIntSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreIntSlicesEqual([]int{2, 3}, []int{1, 2, 3, 4})
	expected = false
	if result != expected {
		t.Errorf("AreIntSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}
}

func TestAreStringSlicesEqual(t *testing.T) {
	result := AreStringSlicesEqual([]string{"1", "2", "3", "4"}, []string{"1", "2", "3", "4"})
	expected := true
	if result != expected {
		t.Errorf("AreStringSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreStringSlicesEqual([]string{}, []string{})
	expected = true
	if result != expected {
		t.Errorf("AreStringSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}

	result = AreStringSlicesEqual([]string{"A", "B", "C"}, []string{"C"})
	expected = false
	if result != expected {
		t.Errorf("AreStringSlicesEqual function test failed. Expected %t, Got %t", expected, result)
	}
}
