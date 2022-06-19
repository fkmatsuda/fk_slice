package fkslice

import "testing"

func TestFindInSlice_int(t *testing.T) {
	slice := []int{1, 5, 7, 8}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(slice, 1)
	if !found {
		t.Error("\"1\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"1\" must be in index 0")
	}
	idx, found = FindInSlice(slice, 7)
	if !found {
		t.Error("\"7\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"7\" must be in index 2")
	}
	_, found = FindInSlice(slice, 9)
	if found {
		t.Error("\"9\" not exists in slice")
	}
}

func TestFindInSlice_string(t *testing.T) {
	slice := []string{"alpha", "beta", "gamma", "theta"}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(slice, "alpha")
	if !found {
		t.Error("\"alpha\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = FindInSlice(slice, "gamma")
	if !found {
		t.Error("\"gamma\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	_, found = FindInSlice(slice, "delta")
	if found {
		t.Error("\"delta\" not exists in slice")
	}
}

type testStruct struct {
	str   string
	value int
}

func TestFindInSlice_struct(t *testing.T) {
	slice := []testStruct{{"alpha", 1}, {"beta", 5}, {"gamma", 7}, {"theta", 8}}
	var (
		idx   int
		found bool
	)
	idx, found = FindInSlice(slice, testStruct{"alpha", 1})
	if !found {
		t.Error("\"alpha\" exists in slice")
	}
	if idx != 0 {
		t.Error("\"alpha\" must be in index 0")
	}
	idx, found = FindInSlice(slice, testStruct{"gamma", 7})
	if !found {
		t.Error("\"gamma\" exists in slice")
	}
	if idx != 2 {
		t.Error("\"gamma\" must be in index 2")
	}
	_, found = FindInSlice(slice, testStruct{"delta", 9})
	if found {
		t.Error("\"delta\" not exists in slice")
	}
}

func TestInterfaceSlice_with_three_elements(t *testing.T) {
	slice := []int{1, 2, 3}
	iSlice := InterfaceSlice(slice)
	if iSlice == nil {
		t.Error("InterfaceSlice() must return a non-nil slice")
	}
	if len(iSlice) != 3 {
		t.Error("InterfaceSlice() must return a slice with 3 elements")
	}
	if iSlice[0] != 1 {
		t.Error("InterfaceSlice() must return a slice with 1st element as 1")
	}
	if iSlice[1] != 2 {
		t.Error("InterfaceSlice() must return a slice with 2nd element as 2")
	}
	if iSlice[2] != 3 {
		t.Error("InterfaceSlice() must return a slice with 3rd element as 3")
	}
}

func TestInterfaceSlice_panic(t *testing.T) {
	var slice interface{}
	defer func() {
		if r := recover(); r == nil {
			t.Error("InterfaceSlice() should panic")
		}
	}()
	InterfaceSlice(slice)
}

func TestInterfaceSlice_nil(t *testing.T) {
	var slice []int
	if InterfaceSlice(slice) != nil {
		t.Error("InterfaceSlice() should return nil")
	}
}
