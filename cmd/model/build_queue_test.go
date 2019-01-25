package model

import (
	"testing"
)

func areSimilar(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for index, _ := range first {
		if first[index] != second[index] {
			return false
		}
	}
	return true
}

func TestOmitExisting(t *testing.T) {
	origin := []string{"a", "b", "c"}
	expected := []string{"a", "c"}
	result := Omit(origin, "b")
	if !areSimilar(result, expected) {
		t.Error("should be similar")
	}
}

func TestOmitMissing(t *testing.T) {
	origin := []string{"a", "b", "c"}
	result := Omit(origin, "d")
	if !areSimilar(origin, result) {
		t.Error("should be similar")
	}
}

func TestOmitNil(t *testing.T) {
	result := Omit(nil, "d")
	if !areSimilar([]string{}, result) {
		t.Error("should be empty")
	}
}
