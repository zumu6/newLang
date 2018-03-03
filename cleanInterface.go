package main

import (
	"fmt"

	"./lib"
)

func testAdd(left interface{}, right interface{}) interface{} {
	left = left.(float64) + right.(float64)
	return left
}

func testAdd2(left float64, right float64) float64 {
	left = left + right
	return left
}

func main() {

	var tree lib.Tree
	tree.Value = 0

	tk1 := lib.Tree{Kind: "INT", Value: 1, Node: nil}
	tk2 := lib.Tree{Kind: "INT", Value: 5, Node: nil}

	result := lib.CreateOperation("+", tk1, tk2)
	fmt.Println(result)

	fmt.Println(result.EvalAny())

	fmt.Println(lib.TypeInference(1.1))

	fmt.Println(lib.InterAdd(1.2, 1, 3.9))
	//tree.showAST()

	left := 2.3
	right := 3.9
	as := left + right
	fmt.Println(as)
}
