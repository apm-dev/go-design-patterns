package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Fatal("GetInstance() result should not be nil")
	}
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("When call AddOne() method for first time, counter value should be 1, but is %d", currentCount)
	}
	counter2 := GetInstance()
	if counter2 != counter1 {
		t.Fatal("Expect to get same instance from GetInstance() but got different one.")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("When call AddOne() method for second time, counter value should be 2, but is %d", currentCount)
	}
}
