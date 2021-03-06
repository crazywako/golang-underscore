package underscore

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	var item TestModel
	Find(arr, func(r TestModel, _ int) bool {
		return r.ID == 1
	}, &item)
	if item != arr[0] {
		t.Error("wrong")
	}
}

func TestFindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	var item TestModel
	FindBy(arr, map[string]interface{}{
		"id": 2,
	}, &item)
	if item != arr[1] {
		t.Error("wrong")
	}
}

func TestChain_Find(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := TestModel{}
	Chain(arr).Find(func(r TestModel, _ int) bool {
		return r.ID == 1
	}).Value(&res)
	if res != arr[0] {
		t.Error("wrong")
	}
}

func TestChain_FindBy(t *testing.T) {
	arr := []TestModel{
		TestModel{ID: 1, Name: "one"},
		TestModel{ID: 2, Name: "two"},
		TestModel{ID: 3, Name: "three"},
	}
	res := TestModel{}
	Chain(arr).FindBy(map[string]interface{}{
		"id": 2,
	}).Value(&res)
	if res != arr[1] {
		t.Error("wrong")
	}
}
