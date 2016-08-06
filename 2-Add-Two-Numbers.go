/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// this solution change l1 's elements
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry,l1_backup,temp:=0,l1,l1
	for l1!=nil && l2!=nil {
		l1.Val+=l2.Val+carry
		if l1.Val>9 {
			l1.Val-=10
			carry=1
		} else {
			carry=0
		}
		temp=l1
		l1=l1.Next
		l2=l2.Next
	}
	if l1==nil {
		temp.Next=l2
	}
	l1=temp.Next
	for l1!=nil {
		l1.Val+=carry
		if l1.Val>9 {
			l1.Val-=10
			carry=1
		} else {
			carry=0
			break
		}
		temp=l1
		l1=l1.Next
	}
	if carry==1 {
		last:=ListNode{1,nil}
		temp.Next=&last
	}
	return l1_backup
}