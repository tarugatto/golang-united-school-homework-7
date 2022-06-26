package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
	"strconv"
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

var jedi Person = Person{
	firstName: "Obi-Wan",
	lastName:  "Kenobi",
	birthDay:  time.Now().AddDate(1977, 05, 04)}
var sith Person = Person{
	firstName: "Anakin",
	lastName:  "Skywalker",
	birthDay:  time.Now().AddDate(1999, 05, 04)}
var ppl People = People{jedi, sith}

func TestLen(t *testing.T){
	//t.Parallel()
		
	expected := len(ppl)
	result := ppl.Len()

	if expected != result {
		t.Errorf("Expected: %d does not equal result: %d", expected, result);
	}
}

func TestLess(t *testing.T){
	//t.Parallel()
	/*ppl = People{
		Person{firstName:"Obi-Wan", lastName:"Kenobi", birthDay: time.Now()},
		Person{firstName:"Anakin", lastName:"Skywalker", birthDay: time.Now()},
	}*/
	result := ppl.Less(0,1)
	expected := false
	if result != expected{
		t.Errorf("Expected: %t does not equal result: %t", expected, result)
	}
}

func TestSwap(t *testing.T){
	//t.Parallel()
	/*sw := People{
		Person{"Anakin", "Skywalker", time.Now()},
		Person{"Luke", "Skywalker", time.Now()},
	}*/
	ppl.Swap(0, 1)
	if ppl[0].firstName != "Anakin" {
		t.Errorf("Swap isn't possible")
	}
}

func TestNew(t *testing.T) {
	//t.Parallel()

	result1, err1 := New("cover")
	if result1 != nil && err1 != strconv.ErrSyntax {
		t.Errorf("Wrong parsing matrix")
	}

	result2, err2 := New("1 2 \n 3")
	if result2 != nil && err2.Error() != "Rows need to be the same length" {
		t.Errorf("Wrong Matrix")
	}

	result3, err3 := New("1 2 \n 3 4")
	expect3 := &Matrix{2, 2, []int{1,2,3,4}}
	
	if result3.cols != expect3.cols || result3.rows != expect3.rows || !reflect.DeepEqual(result3.data, expect3.data) || err3 != nil {
		t.Errorf("Wrong new Matrix")
	}
}

func TestRows(t *testing.T) {
	//t.Parallel()

	m :=&Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := [][]int{{1, 2}, {3, 4}}

	result := m.Rows()

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Wrong Rows")
	}
}

func TestCols(t *testing.T) {
	//t.Parallel()

	m :=&Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := [][]int{{1, 3}, {2, 4}}

	result := m.Cols()

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Wrong Cols")
	}
}

func TestSet(t *testing.T) {
	//t.Parallel()

	m :=&Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := []int{1, 2, 3, 10}

	result1 := m.Set(1, 1, 10)

	if !reflect.DeepEqual(m.data, expect) || !result1{
		t.Errorf("Wrong Matrix after Set value")
	}

	result2 := m.Set(2, 2, 10)
	if result2 {
		t.Errorf("wring Set value for Matrix")
	}
}