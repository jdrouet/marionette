package parser

import "testing"

func TestParseEmpty(t *testing.T) {
	repo, err := Parse("../../sample/empty.json")
	if err != nil {
		t.Error(err)
	}
	if len(repo.Projects) != 0 {
		t.Error("Should not have any project")
	}
}

func TestParseSimple(t *testing.T) {
	repo, err := Parse("../../sample/simple.json")
	if err != nil {
		t.Error(err)
	}
	if len(repo.Projects) == 0 {
		t.Error("Should have some project")
	}
}

func TestParseUndefined(t *testing.T) {
	_, err := Parse("../../sample/undefined.json")
	if err == nil {
		t.Error("Should throw an error")
	}
}

func TestParseNotJson(t *testing.T) {
	_, err := Parse("../../sample/notjson.txt")
	if err == nil {
		t.Error("Should throw an error")
	}
}
