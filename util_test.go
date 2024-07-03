package main

import (	
	"reflect"
	"testing"
	"strings"
)

type T struct {
	*testing.T
}

func (t *T) assertEqual (x,y any ) {
	if reflect.DeepEqual(x,y) == false {
		t.Fatalf("%v != %v", x,y)
	}

}


func TestPermutations1(t *testing.T) {

	w := T{t}

	w.assertEqual(Permutations([]int{}), [][]int{})
	w.assertEqual(Permutations([]int{1}), [][]int{{1}})

	w.assertEqual(Permutations([]int{1,2}), [][]int{
		{1, 2, },
		{2, 1, },
	})

	w.assertEqual(Permutations([]int{1,2,3}), [][]int{
		{1,2,3,},
		{1,3,2,},
		{2,1,3,},
		{2,3,1,},
		{3,1,2,},
		{3,2,1,},
	})

	w.assertEqual(Permutations(strings.Split("Apr 27 00:25"," "),


	),[][]string{


	})

}