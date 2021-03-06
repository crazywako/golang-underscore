package underscore

import (
	"testing"
)

func TestUniq(t *testing.T) {
	res := make([]int, 0)
	Uniq([]int{1, 2, 1, 2, 1, 3}, nil, &res)
	if len(res) != 3 {
		t.Error("wrong")
	}
}

func TestUniqBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "a"},
		TestModel{ID: 2, Name: "a"},
		TestModel{ID: 3, Name: "a"},
	}
	res := make([]TestModel, 0)
	UniqBy(arr, "Name", &res)
	if len(res) != 1 {
		t.Error("wrong")
	}
}

func TestChain_Uniq(t *testing.T) {
	res := make([]int, 0)
	Chain([]int{1, 2, 1, 4, 1, 3}).Uniq(func(n, _ int) (int, error) {
		return n % 2, nil
	}).Value(&res)
	if len(res) != 2 {
		t.Error(res)
	}
}

func TestChain_UniqBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "a"},
		TestModel{ID: 2, Name: "a"},
		TestModel{ID: 3, Name: "a"},
	}
	res := make([]TestModel, 0)
	Chain(arr).UniqBy("Name").Value(&res)
	if len(res) != 1 {
		t.Error("wrong")
	}
}
