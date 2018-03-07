package lib

import "reflect"

type Tree struct { //構文木
	Kind  string
	Value interface{}
	Node  []*Tree
}

type END struct { //スライスの末尾を示す
	isEnd bool
}

func CreateAddTree(expr1, expr2 int) Tree { //加算用の構文木を返してくれる
	resultTree := Tree{Kind: "+", Value: 0, Node: nil}
	resultTree.Node = append(resultTree.Node, &Tree{Kind: "INT", Value: expr1, Node: nil})
	resultTree.Node = append(resultTree.Node, &Tree{Kind: "INT", Value: expr2, Node: nil})
	return resultTree
}

func CreateOperation(operator string, args ...Tree) Tree { //引数にとった演算子の構文木を返してくれる
	ResultTree := Tree{Kind: operator, Value: 0, Node: nil}
	for _, v := range args {
		ResultTree.AddTree(v)
	}
	return ResultTree
}

func (tree *Tree) AddTree(inputTree Tree) { //構文木に構文木を追加
	tree.Node = append(tree.Node, new(Tree))
	tree.Node[len(tree.Node)-1] = &inputTree
}

func TypeInference(expr interface{}) string { //型情報を返す
	typeInformation := reflect.ValueOf(expr)
	switch typeInformation.Type().Kind() {
	case reflect.Int:
		return "INT"
	case reflect.Float64, reflect.Float32:
		return "FLOAT"
	case reflect.String:
		return "STRING"
	default:
		return "UNKNOWN"
	}
}

func (tree *Tree) getOperator() string { //
	return tree.Value.(string)
}
