package coverage

import (
	//"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
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
	t.Parallel()
	expected := len(ppl)
	result := ppl.Len()

	if expected != result {
		t.Errorf("Expected: %d does not equal result: %d", expected, result);
	}
}

func Test_Less(t *testing.T) {
	var birthDate time.Time = time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	testData := map[string]struct {
		p       People
		boolLess bool
	}{
		"all_force_users": {[]Person{{"Luke", "Skywalker", birthDate}, {"Luke", "Skywalker", birthDate}}, false},
		"son_father":  {[]Person{{"Luke", "Skywalker", birthDate}, {"Anakin", "Skywalker", birthDate}}, false},
		"father_son":  {[]Person{{"Anakin", "Skywalker", birthDate}, {"Luke", "Skywalker", birthDate}}, true},
		"master_apprentice": {[]Person{{"Qui-Gon", "Jinn", birthDate}, {"Obi-Wan", "Kenobi", birthDate}}, false},
		"apprentice_master": {[]Person{{"Obi-Wan", "Jinn", birthDate}, {"Qui-Gon", "Kenobi", birthDate}}, true},
		"master_apprentice1": {[]Person{{"Obi-Wan", "Kenobi", birthDate}, {"Anakin", "Skywalker", birthDate.AddDate(1, 0, 0)}}, false},
		"master_apprentice2": {[]Person{{"Obi-Wan", "Kenobi", birthDate}, {"Anakin", "Skywalker", birthDate.AddDate(-1, 0, 0)}}, true}}

	for name, tcase := range testData {
		got := tcase.p.Less(0, 1)
		if got != tcase.boolLess {
			t.Errorf("[%s] expected: %t, got %t", name, tcase.boolLess, got)
		}
	}
}

func TestSwap(t *testing.T){
	t.Parallel()
	ppl.Swap(0, 1)
	if ppl[0] != sith {
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
		t.Errorf("Wrong Rows count")
	}
}

func TestCols(t *testing.T) {
	//t.Parallel()

	m :=&Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := [][]int{{1, 3}, {2, 4}}

	result := m.Cols()

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Wrong Cols count")
	}
}

func TestSet(t *testing.T) {
	//t.Parallel()

	m :=&Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := []int{1, 2, 3, 10}

	result1 := m.Set(1, 1, 10)

	if !reflect.DeepEqual(m.data, expect) || !result1{
		t.Errorf("Error: Wrong Matrix after Set value")
	}

	result2 := m.Set(2, 2, 10)
	if result2 {
		t.Errorf("Error: Set value for Matrix")
	}
}