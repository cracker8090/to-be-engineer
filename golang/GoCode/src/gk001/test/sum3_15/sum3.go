package sum3_test

import "sort"

//15-sum3
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.IntSlice(nums).Sort()
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	zeros := make([][]int,0)
	for m := 1; m <= len(nums)-2; m++ {
		l,r := 0, len(nums)-1
		if m-1 >= 0 && nums[m] == nums[m-1] {
			if m-2 >= 0 && nums[m-2] == nums[m-1] {
				continue
			}
			l,r = m-1,len(nums)-1
		}

		for l < m && r > m {
			sum := nums[l] + nums[m] + nums[r]
			if sum == 0 {
				zero := []int{nums[l],nums[m],nums[r]}
				zeros = append(zeros,zero)
			}
			if sum > 0 {
				for cur := r;r > m && nums[cur] == nums[r]; cur,r = r,r-1 {
				}
			} else {
				for cur := l; l < m && nums[cur] == nums[l]; cur,l = l,l+1 {
				}
			}
		}
	}
	return zeros
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//206-反转链表
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	var neadHead,next *ListNode

	for head != nil {
		next = head.Next
		head.Next = neadHead
		neadHead = head
		head = next
	}
	return neadHead
	//
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//
	//newHead := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//
	//return newHead




}
//21-merge list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// if l1 == nil {
	//     return l2
	// }
	// if l2 == nil {
	//     return l1
	// }

	// if l1.Val <l2.Val {
	//     l1.Next = mergeTwoLists(l1.Next,l2)
	//     return l1
	// } else {
	//     l2.Next = mergeTwoLists(l1,l2.Next)
	//     return l2
	// }

	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 == nil {
		preHead.Next = l2
	}
	if l2 == nil {
		preHead.Next = l1
	}
	return result.Next

}
