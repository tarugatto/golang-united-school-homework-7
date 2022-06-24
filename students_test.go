package coverage

import (
	"os"
	"testing"
	"time"
	//"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var ppl People

func TestLen(t *testing.T){
	for i:=0; i<3; i++{
		ppl= append(ppl, Person{})
	}
	
	expected := len(ppl)
	result := ppl.Len()

	if expected != result {
		t.Errorf("Expected: %d does not equal result: %d", expected, result);
	}
}

func TestLess(t *testing.T){
	ppl = People{
		Person{firstName:"Obi-Wan", lastName:"Kenobi", birthDay: time.Now()},
		Person{firstName:"Anakin", lastName:"Skywalker", birthDay: time.Now()},
	}
	result := ppl.Less(0,1)
	expected := false
	if result != expected{
		t.Errorf("Expected: %t does not equal result: %t", expected, result)
	}
}

func TestSwap(t *testing.T){
	
	sw := People{
		Person{"Anakin", "Skywalker", time.Now()},
		Person{"Luke", "Skywalker", time.Now()},
	}
	sw.Swap(0, 1)
	if sw[0].firstName != "Luke" {
		t.Errorf("Swap isn't possible")
	}
}