package lib

import (
	"fmt"
	"reflect"
)

func (code *Tree) EvalAny() interface{} { //評価
	switch code.Kind {
	case "INT", "FLOAT":
		return code.EvalNumeric() //数字なら数字を返す
	case "+":
		return code.Node[0].EvalAny().(int) + code.Node[1].EvalAny().(int)
	default:
		return 0
	}
}

func (code *Tree) EvalNumeric() interface{} { //数字の評価
	i := code.Value
	return i
}

func InterAdd(left interface{}, right ...interface{}) interface{} { //加算を評価(interfaceを使用)
	fmt.Println("START:", left, "+", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}

	switch leftVal := left.(type) {
	case int:
		//fmt.Println("INT")
		switch rightVal := right[0].(type) {
		case int:
			//fmt.Println("int")
			left = leftVal + rightVal
			return InterAdd(left, right[1:]...)
		case float64:
			//fmt.Println("float64")
			left = float64(leftVal) + rightVal
			return InterAdd(left, right[1:]...)
		case END:
			//fmt.Println("END")
			return left
		default:
			//fmt.Println(reflect.TypeOf(left))
			//fmt.Println("$$$$$$")
		}
	case float64:
		//fmt.Println("FLOAT")
		switch rightVal := right[0].(type) {
		case int:
			//fmt.Println("int")
			left = leftVal + float64(rightVal)
			return InterAdd(left, right[1:]...)
		case float64:
			fmt.Println("float64")
			left = leftVal + rightVal
			fmt.Println(left)
			return InterAdd(left, right[1:]...)
		case END:
			//fmt.Println("END")
			return left
		default:
			//fmt.Println("type = ", reflect.TypeOf(left))
			//fmt.Println("###")
		}
	default:
		return 0
	}
	fmt.Println("ENDED")
	return 0
}
