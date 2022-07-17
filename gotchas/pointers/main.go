package main

import "fmt"

type MyStruct struct {
	Name string
}

func main() {
	printResult(got())
	printResult(gotMod())
	printResult(modifyWithGot(got()))
	printStructResult(gotStruct())
	printStructResult(gotStructFix())
	printStructResult(gotStructMod())
}

func printResult(result []*int) {
	for _, r := range result {
		fmt.Println(*r)
	}
	fmt.Println(result)
}

func printStructResult(result []*MyStruct) {
	for _, r := range result {
		fmt.Println(*r)
	}
	fmt.Println(result)
}

// modifying all values using dereference
func modifyWithGot(in []*int) []*int {
	*in[0] = 5
	return in
}

// problem
// variable v is always the same in loop
// it has exactly the one memory address
func got() []*int {
	a := []int{1, 2, 3, 4}
	result := make([]*int, 0)
	for _, v := range a {
		result = append(result, &v)
	}
	return result
}

func gotMod() []*int {
	a := []int{1, 2, 3, 4}
	result := make([]*int, 0)
	for _, v := range a {
		vv := v
		result = append(result, &vv)
	}
	return result
}

// in result all the same struct
// that's why linter bans such cases
func gotStruct() []*MyStruct {
	a := []MyStruct{{Name: "1"}, {Name: "2"}, {Name: "3"}}
	result := make([]*MyStruct, 0)
	for _, v := range a {
		result = append(result, &v)
	}
	return result
}

func gotStructFix() []*MyStruct {
	a := []MyStruct{{Name: "1"}, {Name: "2"}, {Name: "3"}}
	result := make([]*MyStruct, 0)
	for _, v := range a {
		vv := v
		result = append(result, &vv)
	}
	return result
}

// but when we insert pointers without & - it's all right
// because we pointing not to inner variable but input pointers
func gotStructMod() []*MyStruct {
	a := []*MyStruct{{Name: "1"}, {Name: "2"}, {Name: "3"}}
	result := make([]*MyStruct, 0)
	for _, v := range a {
		result = append(result, v)
	}
	return result
}
