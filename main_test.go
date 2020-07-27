package main

import (
	"reflect"
	"testing"
)

func TestParseGuess(t *testing.T) {
	got, _ := ParseGuess('l', "hello")
	want := []int{2,3}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("Got %v, want %v", got, want)
	}
}
