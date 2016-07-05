package reflect_example

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	// 同じ型で同じデータなので正しくマッチする
	if !reflect.DeepEqual(m1, m2) {
		t.Errorf("want %#v got %#v", m1, m2)
	}

	// 型が違うとマッチしない
	m3 := map[string]string{"a": "1", "b": "2", "c": "3"}
	if !reflect.DeepEqual(m1, m3) {
		t.Errorf("want %#v got %#v", m1, m3)
	}

	// これはマッチする
	m4 := map[string]int{"c": 3, "b": 2, "a": 1}
	if !reflect.DeepEqual(m1, m4) {
		t.Errorf("want %#v got %#v", m1, m4)
	}
}

type T struct {
	x  int
	ss []string
	m  map[string]int
}

func TestStruct(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	t1 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}
	t2 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("want %#v got %#v", t1, t2)
	}
}

func TestSlice(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("want %#v got %#v", s1, s2)
	}
}

func TestInterfaceMap(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	var m2 interface{}
	m2 = map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(m1, m2) {
		t.Errorf("want %#v got %#v", m1, m2)
	}
}

type mapTest struct {
	a, b map[string]int
	eq   bool
}

var mapTests = []mapTest{
	{map[string]int{"a": 1}, map[string]int{"b": 1}, false},
	{map[string]int{"a": 1}, map[string]int{"a": 1}, true},
}

func TestMapTable(t *testing.T) {
	for _, test := range mapTests {
		if r := reflect.DeepEqual(test.a, test.b); r != test.eq {
			t.Errorf("when a = %#v and b = %#v, want %t, got %t",
				test.a, test.b, r, test.eq)
		}
	}
}
