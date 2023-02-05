package main

import "fmt"

type MyStack struct {
	data []string
}

func (myStack *MyStack) push(ch string) {
	myStack.data = append(myStack.data, ch)
}

func (myStack *MyStack) pop() string {
	length := len(myStack.data)
	if length == 0 {
		return ""
	}
	topData := myStack.data[length-1]
	myStack.data = myStack.data[:len(myStack.data)-1]
	return topData
}

func main() {
	myEmptyStack := &MyStack{
		data: []string{},
	}

	expression := "[({})]"

	var flag int = 1

	for i := 0; i < len(expression); i++ {
		if string(expression[i]) == "(" || string(expression[i]) == "{" || string(expression[i]) == "[" {
			myEmptyStack.push(string(expression[i]))
		} else if string(expression[i]) == ")" || string(expression[i]) == "}" || string(expression[i]) == "]" {
			fromStack := myEmptyStack.pop()
			if string(expression[i]) == ")" && fromStack != "(" {
				flag = -1
				break
			}
			if string(expression[i]) == "}" && fromStack != "{" {
				flag = -1
				break
			}
			if string(expression[i]) == "]" && fromStack != "[" {
				flag = -1
				break
			}
		} else {
			flag = -1
			break
		}
	}
	if flag == -1 || len(myEmptyStack.data) != 0 {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid")
	}
}
