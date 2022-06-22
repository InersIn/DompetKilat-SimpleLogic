package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	const str string = "abcba"
	fmt.Println(IsPalindrome(str))
}

func TestFindPrimeByRange(t *testing.T) {
	fmt.Println(FindPrimeByRange(11, 40))
}

func TestArrayGroup(t *testing.T) {
	arr := []string{"a", "a", "a", "b", "c", "c", "b", "b", "b", "d", "d", "e", "e", "e"}
	fmt.Println(ArrayGroup(arr))
}

func TestCountSameElement(t *testing.T) {
	arr := []string{"a", "a", "a", "b", "c", "c", "b", "b", "b", "d", "d", "e", "e", "e"}
	fmt.Println(CountSameElements(arr))
}
