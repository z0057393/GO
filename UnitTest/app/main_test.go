package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "Fizz"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		var result string = check(test.input)

		if result != test.expected {
			t.Errorf("Pour l'entrée %d, résultat attendu : %s, résultat obtenu : %s", test.input, test.expected, result)
		}
	}
}

func TestFizz(t *testing.T) {
	expected := "Fizz"
	result := Fizz()
	if result != expected {
		t.Errorf("Résultat attendu : %s, résultat obtenu : %s", expected, result)
	}
}

func TestBuzz(t *testing.T) {
	expected := "Buzz"
	result := Buzz()
	if result != expected {
		t.Errorf("Résultat attendu : %s, résultat obtenu : %s", expected, result)
	}
}

func TestFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	result := FizzBuzz()
	if result != expected {
		t.Errorf("Résultat attendu : %s, résultat obtenu : %s", expected, result)
	}
}
