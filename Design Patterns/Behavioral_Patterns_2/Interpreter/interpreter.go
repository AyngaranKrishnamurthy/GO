package interpreter

import (
	"strconv"
	"strings"
)

func Calculate(o string) (int, error) { //Gets a string and returns an integer or an error
	stack := PolishNotationStack{}
	operators := strings.Split(o, " ")

	for _, operatorString := range operators {
		if isOperator(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperationFunc(operatorString)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}
	return int(stack.Pop()), nil
}

func isOperator(o string) bool { //Checking if the returned element is an operator
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}
	return false
}

const ( //Defining possible operators in constants
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

type PolishNotationStack []int

func (p *PolishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *PolishNotationStack) Pop() int {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1] //returns last element in stack as temp
		*p = (*p)[:length-1]   //removes last stack element
		return temp
	}
	return 0
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}

	}
	return nil
}
