package lib

import (
	"errors"
	"fmt"
	"reflect"
)

func (tree *Tree) EvalAny() interface{} { //評価
	switch tree.Kind {
	case "INT", "FLOAT":
		return tree.EvalNumeric() //数字なら数字を返す
	case "+":
		return tree.Node[0].EvalAny().(int) + tree.Node[1].EvalAny().(int)
	default:
		return 0
	}
}

func (tree *Tree) EvalNumeric() interface{} { //数字の評価
	i := tree.Value
	return i
}

func (tree *Tree) Eval() interface{} {
	switch tree.getOperator {
	case "+":

	}
}

func EvalAddition(left interface{}, right ...interface{}) (interface{}, error) { //加算を評価(interfaceを使用)
	fmt.Println("START:", left, "+", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}

	switch leftVal := left.(type) {
	case int:
		switch rightVal := right[0].(type) {
		case int:
			left = leftVal + rightVal
			return EvalAddition(left, right[1:]...)
		case float64:
			left = float64(leftVal) + rightVal
			return EvalAddition(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	case float64:

		switch rightVal := right[0].(type) {
		case int:

			left = leftVal + float64(rightVal)
			return EvalAddition(left, right[1:]...)
		case float64:
			fmt.Println("float64")
			left = leftVal + rightVal
			fmt.Println(left)
			return EvalAddition(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	default:
		return 0, errors.New("found undefined type in left value")
	}
}

func EvalSubtraction(left interface{}, right ...interface{}) (interface{}, error) { //減算を評価
	fmt.Println("START:", left, "-", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}
	switch leftVal := left.(type) {
	case int:
		switch rightVal := right[0].(type) {
		case int:
			left = leftVal - rightVal
			return EvalSubtraction(left, right[1:]...)
		case float64:
			left = float64(leftVal) - rightVal
			return EvalSubtraction(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	case float64:

		switch rightVal := right[0].(type) {
		case int:

			left = leftVal - float64(rightVal)
			return EvalSubtraction(left, right[1:]...)
		case float64:
			left = leftVal - rightVal
			fmt.Println(left)
			return EvalSubtraction(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	default:
		return 0, errors.New("found undefined type in left value")
	}
}

func EvalMultiple(left interface{}, right ...interface{}) (interface{}, error) { //乗算を評価
	fmt.Println("START:", left, "*", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}
	switch leftVal := left.(type) {
	case int:
		switch rightVal := right[0].(type) {
		case int:
			left = leftVal * rightVal
			return EvalMultiple(left, right[1:]...)
		case float64:
			left = float64(leftVal) * rightVal
			return EvalMultiple(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	case float64:

		switch rightVal := right[0].(type) {
		case int:

			left = leftVal * float64(rightVal)
			return EvalMultiple(left, right[1:]...)
		case float64:
			left = leftVal * rightVal
			fmt.Println(left)
			return EvalMultiple(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	default:
		return 0, errors.New("found undefined type in left value")
	}
}

func EvalDivision(left interface{}, right ...interface{}) (interface{}, error) { //除算を評価
	fmt.Println("START:", left, "/", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}
	switch leftVal := left.(type) {
	case int:
		switch rightVal := right[0].(type) {
		case int:
			left = leftVal / rightVal
			return EvalDivision(left, right[1:]...)
		case float64:
			left = float64(leftVal) / rightVal
			return EvalDivision(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	case float64:

		switch rightVal := right[0].(type) {
		case int:

			left = leftVal / float64(rightVal)
			return EvalDivision(left, right[1:]...)
		case float64:
			left = leftVal / rightVal
			fmt.Println(left)
			return EvalDivision(left, right[1:]...)
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	default:
		return 0, errors.New("found undefined type in left value")
	}
}

func EvalRemainder(left interface{}, right ...interface{}) (interface{}, error) { //剰余を評価
	fmt.Println("START:", left, "%", right)

	if (reflect.TypeOf(right[len(right)-1]) != reflect.TypeOf(END{})) {
		right = append(right, END{true})
	}
	switch leftVal := left.(type) {
	case int:
		switch rightVal := right[0].(type) {
		case int:
			left = leftVal % rightVal
			return EvalRemainder(left, right[1:]...)
		case float64:
			return 0, nil
		case END:
			return left, nil
		default:
			return 0, errors.New("found undefined type in right value")
		}
	case float64:
		return 0, nil
	default:
		return 0, errors.New("found undefined type in left value")
	}
}
