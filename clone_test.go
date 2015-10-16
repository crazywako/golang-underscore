package underscore

import (
	"testing"
)

func TestClone(t *testing.T) {
	arr := []int{ 1, 2, 3 }
	v, _ := Clone(arr)
	dstArr, _ := v.([]int)
	dstArr[0] = 11
	if arr[0] == dstArr[0] {
		t.Error("wrong")
	}
}

func TestChain_Clone(t *testing.T) {
	dict := map[string]int{ "a": 1, "b": 2 }
	v, _ := Chain(dict).Clone().Value()
	dstMap, _ := v.(map[string]int)
	dstMap["a"] = 11
	if dict["a"] == dstMap["a"] {
		t.Error("wrong")
	}
}