package model

import (
	"testing"
)

func TestFilterSuccess(t *testing.T) {
	p := Project{
		Exclude: []string{
			"*.txt",
		},
	}
	input := []string{
		"toto.txt",
		"somewhere/tata.txt",
		"elsewhere/trilili.json",
	}
	result, err := p.filterExcluded(input)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("should only return one file")
	}
	if result[0] != input[2] {
		t.Error("wrong file filtered")
	}
}
