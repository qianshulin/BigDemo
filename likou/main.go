package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list1 = new(ListNode)
	var list2 = new(ListNode)
	list1.Val = 1
	list1.Val = 2

	addTwoNumbers(list1, list2)
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一个变量carry
	var carry int
	//将结构体对象进行初始化(地址值引用传递)
	resultList := &ListNode{}
	//将地址值赋值给current
	current := resultList
	//死循环
	for {
		//判断参数是否为空,为空则停止循环
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		//判断l1是否为空，
		if l1 != nil {
			carry += l1.Val

			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node := ListNode{}
		if carry <= 9 {
			node = ListNode{
				Val: carry,
			}
			carry = 0
		} else {
			node = ListNode{
				Val: carry - 10,
			}
			carry = 1
		}
		current.Next = &node
		current = &node
	}
	return resultList.Next
}
