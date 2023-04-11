package main

import "testing"

func TestCustomTestOne(t *testing.T) {
	t.Log("Hello World test")
}

// Just another test
func TestMyAdd(t *testing.T) {
	type testStruct struct {
		data []int
		res  int
	}

	dt := []testStruct{{
		data: []int{1, 2},
		res:  3,
	}, {
		data: []int{10, 20},
		res:  30,
	}, {
		data: []int{100, 200},
		res:  300,
	}, {
		data: []int{7, 9},
		res:  16,
	},
	}

	for _, v := range dt {
		act := adder(v.data...)
		if act != v.res {
			t.Errorf("[Test Failed] Expected %d Got %d", v.res, act)
		}
	}
}
