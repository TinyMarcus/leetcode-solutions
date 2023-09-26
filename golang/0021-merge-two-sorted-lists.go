/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    node := new(ListNode)
    l1_ptr := list1
    l2_ptr := list2
    tail := node

    for l1_ptr != nil && l2_ptr != nil {
        if l1_ptr.Val < l2_ptr.Val {
            tail.Next = l1_ptr
            l1_ptr = l1_ptr.Next
        } else {
            tail.Next = l2_ptr
            l2_ptr = l2_ptr.Next
        }

        tail = tail.Next
    }

    if l1_ptr != nil {
        tail.Next = l1_ptr
    } else if l2_ptr != nil {
        tail.Next = l2_ptr
    }

    return node.Next
}